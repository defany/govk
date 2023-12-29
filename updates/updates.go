package updates

import (
	"context"
	"errors"
	"fmt"
	"github.com/defany/govk/api"
	apiModel "github.com/defany/govk/api/model"
	"github.com/defany/govk/updates/model"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	url2 "net/url"
	"reflect"
	"strconv"
)

type longPoll struct {
	Key    string
	Server string
	TS     string
	Wait   int
	cancel context.CancelFunc
}

type callback[T any] func(data T)

type Updates struct {
	callbacks map[string][]func(data any)
	longPoll  *longPoll
	api       *apiModel.ApiProvider
	client    *fasthttp.Client
}

func NewUpdates(api *apiModel.ApiProvider) *Updates {
	return &Updates{
		api:       api,
		callbacks: make(map[string][]func(data any)),
		longPoll: &longPoll{
			Wait: 25,
		},
		client: &fasthttp.Client{},
	}
}

// Check handler.
func (u *Updates) Check() error {
	return u.CheckWithContext(context.Background())
}

// CheckWithContext handler.
func (u *Updates) CheckWithContext(ctx context.Context) error {
	return u.run(ctx)
}

// TODO: Add check for new updates

func (u *Updates) run(ctx context.Context) error {
	if err := u.getLongPollParams(true); err != nil {
		return err
	}

	ctx, u.longPoll.cancel = context.WithCancel(ctx)

	for {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				return nil
			}
		default:
			resp, err := u.check()
			if err != nil {
				return err
			}
			if len(resp.Updates) > 0 {
				if callbackFuncs, ok := u.callbacks[resp.Updates[0].Type]; ok {
					for _, callback := range callbackFuncs {
						callback(resp.Updates[0].Object)
					}
				}
			}
		}
	}
	return nil
}

func (u *Updates) check() (response model.Response, err error) {
	uri, err := url2.Parse(u.longPoll.Server)
	if err != nil {
		return response, err
	}
	params := url2.Values{}

	params.Add("act", "a_check")
	params.Add("key", u.longPoll.Key)
	params.Add("ts", u.longPoll.TS)
	params.Add("wait", strconv.Itoa(u.longPoll.Wait))

	uri.RawQuery = params.Encode()

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri.String())
	req.Header.SetMethodBytes([]byte(fasthttp.MethodGet))

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	res.StreamBody = true

	err = u.client.Do(req, res)
	if err != nil {
		return response, err
	}

	reader := res.BodyStream()

	response, err = parseResponse(reader)
	if err != nil {
		return response, err
	}

	err = u.checkResponse(response)
	if err != nil {
		return response, err
	}

	return response, err
}

// On - function to helps you handle updates from VK
//
// With generics it can be helpful to work without interface{}, just pass it into function and see a magic.
//
// Try to use types only from <any>/model package, but if you trust yourself you can provide your own types. For instance - types for user bots.
// Remember, that you can pass only struct as generic. No pointers, no something else.
//
// All event objects named as in vk, for example: messages_new -> MessagesNew
//
// Keep in mind, that this function will panic if something went wrong
func On[T any](updates *Updates, event string, callback callback[T]) {
	var zero T

	typ := reflect.TypeOf(zero)
	if typ.Kind() == reflect.Pointer {
		panic("pointer has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	if typ.Kind() != reflect.Struct {
		panic("something not a struct has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	updates.callbacks[event] = append(updates.callbacks[event],
		func(data any) {
			pd, ok := data.(T)
			if !ok {
				panic(fmt.Sprintf("failed to cast incoming type %s to %s", reflect.TypeOf(data), typ))
			}

			callback(pd)
		})
}

func (u *Updates) getLongPollParams(updateTS bool) error {
	groupID, err := u.api.Groups.GetByID(nil)
	if err != nil {
		return err
	}

	params := api.Params{
		"group_id": groupID.Groups[0].ID,
	}

	lpServer, err := u.api.Groups.GetLongPollServer(params)
	if err != nil {
		return err
	}

	u.longPoll.Key = lpServer.Key
	u.longPoll.Server = lpServer.Server

	if updateTS {
		u.longPoll.TS = lpServer.TS
	}

	return nil
}

func parseResponse(reader io.Reader) (response model.Response, err error) {
	decoder := json.NewDecoder(reader)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return response, err
		}

		t, ok := token.(string)
		if !ok {
			continue
		}

		switch t {
		case LongpollResponseFailed:
			raw, err := decoder.Token()
			if err != nil {
				return response, err
			}

			parsedRaw, ok := raw.(float64)
			if !ok {
				return response, ErrFailedToCastRawToFloat64
			}
			response.Failed = int(parsedRaw)
		case LongpollResponseUpdates:
			var updates []model.Update

			err = decoder.Decode(&updates)
			if err != nil {
				return response, err
			}

			response.Updates = updates
		case LongpollResponseTS:
			// can be a number in the response with "failed" field: {"ts":8,"failed":1}
			// or string, e.g. {"ts":"8","updates":[]}
			rawTs, err := decoder.Token()
			if err != nil {
				return response, err
			}

			if ts, isNumber := rawTs.(float64); isNumber {
				response.Ts = strconv.Itoa(int(ts))
			} else {
				response.Ts = rawTs.(string)
			}
		}
	}

	return response, err
}

func (u *Updates) checkResponse(response model.Response) (err error) {
	switch response.Failed {
	case LongpollUnknownStatus:
		u.longPoll.TS = response.Ts
	case LongpollHistoryMissedStatus:
		u.longPoll.TS = response.Ts
	case LongpollKeyExpiredStatus:
		err = u.getLongPollParams(false)
	case LongpollInfoMissedStatus:
		err = u.getLongPollParams(true)
	default:
		err = fmt.Errorf("%w Code: %d", ErrLongpollFailed, response.Failed)
	}

	return
}
