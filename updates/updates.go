package updates

import (
	"context"
	"errors"
	"fmt"
	"github.com/defany/govk/api"
	apiModel "github.com/defany/govk/api/model"
	"github.com/defany/govk/updates/model/longpoll"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	"net/url"
	"strconv"
)

type longPoll struct {
	Key     string
	Server  string
	TS      string
	Wait    int
	cancel  context.CancelFunc
	isFixed bool
}

type callback[T any] func(ctx context.Context, data T)

type middleware[T any] func(ctx context.Context, data T) bool

type Updates struct {
	callbacks   map[string][]callback[json.RawMessage]
	middlewares map[string][]middleware[json.RawMessage]
	longPoll    *longPoll
	api         *apiModel.ApiProvider
	client      *fasthttp.Client
}

func NewUpdates(api *apiModel.ApiProvider) *Updates {
	return &Updates{
		api:         api,
		callbacks:   make(map[string][]callback[json.RawMessage]),
		middlewares: make(map[string][]middleware[json.RawMessage]),
		longPoll: &longPoll{
			Wait: 25,
		},
		client: &fasthttp.Client{},
	}
}

// Run handler.
func (u *Updates) Run() error {
	return u.RunWithContext(context.Background())
}

// RunWithContext handler.
func (u *Updates) RunWithContext(ctx context.Context) error {
	return u.run(ctx)
}

// WithFixedServerURL method should be used to fix the URL of the longpoll server.
// Keep in mind, that this URL will be used all time
func (u *Updates) WithFixedServerURL(serverUrl string) error {
	url, err := url.ParseRequestURI(serverUrl)
	if err != nil {
		return err
	}

	u.longPoll.Server = url.String()
	u.longPoll.isFixed = true
	return nil
}

func (u *Updates) run(ctx context.Context) error {
	if err := u.refreshLongPollParams(true); err != nil {
		return err
	}

	ctx, u.longPoll.cancel = context.WithCancel(ctx)

	for {
	NEXT:
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

			if len(resp.Updates) < 1 {
				continue
			}

			upd := resp.Updates[0]

			callbacks, ok := u.callbacks[upd.Type]
			if !ok {
				continue
			}

			middlewares, ok := u.middlewares[upd.Type]
			if ok {
				for _, midd := range middlewares {
					if !midd(ctx, upd.Object) {
						goto NEXT
					}
				}
			}

			for _, callback := range callbacks {
				callback(ctx, upd.Object)
			}
		}
	}
}

func (u *Updates) check() (response model.Response, err error) {
	uri, err := url.Parse(u.longPoll.Server)
	if err != nil {
		return response, err
	}
	params := url.Values{}

	params.Add("act", "a_check")
	params.Add("key", u.longPoll.Key)
	params.Add("ts", u.longPoll.TS)
	params.Add("wait", strconv.Itoa(u.longPoll.Wait))

	uri.RawQuery = params.Encode()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri.String())
	req.Header.SetMethod(fasthttp.MethodGet)

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

func (u *Updates) refreshLongPollParams(isUpdateTs bool) error {
	groupID, err := u.api.Groups.GetByID()
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

	if !u.longPoll.isFixed {
		u.longPoll.Server = lpServer.Server
	}

	if isUpdateTs {
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
				return response, fmt.Errorf("%w float64", ErrFailedRawCast)
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
	case LongpollOkStatus:
		u.longPoll.TS = response.Ts
	case LongpollHistoryMissedStatus:
		u.longPoll.TS = response.Ts
	case LongpollKeyExpiredStatus:
		err = u.refreshLongPollParams(false)
	case LongpollInfoMissedStatus:
		err = u.refreshLongPollParams(true)
	default:
		err = fmt.Errorf("%w Code: %d", ErrLongpollFailed, response.Failed)
	}

	return nil
}
