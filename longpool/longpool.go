package longpool

import (
	gripDecoder "compress/gzip"
	"context"
	"errors"
	"fmt"
	"github.com/defany/govk/api"
	"github.com/defany/govk/longpool/model"
	"github.com/defany/govk/vk"
	_ "github.com/defany/govk/vk"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	"strconv"
)

type LongPoll struct {
	Key                  string
	Server               string
	TS                   string
	Wait                 int
	GroupID              int
	VK                   *govk.VK
	Client               *fasthttp.Client
	cancel               context.CancelFunc
	funcFullResponseList []func(model.Response)
}

func NewLongPoll(vk *govk.VK, groupID int) (*LongPoll, error) {
	lp := &LongPoll{
		VK:      vk,
		GroupID: groupID,
		Wait:    25,
		Client:  &fasthttp.Client{},
	}

	err := lp.getLongPoolParams(true)

	return lp, err
}

func (lp *LongPoll) getLongPoolParams(updateTS bool) error {
	params := api.Params{
		"group_id": lp.GroupID,
	}

	lpServer, err := lp.VK.Api.Groups.GetLongPoolServer(params)
	if err != nil {
		return err
	}

	lp.Key = lpServer.Key
	lp.Server = lpServer.Server

	if updateTS {
		lp.TS = lpServer.TS
	}

	return nil
}

func (lp *LongPoll) check() (response model.Response, err error) {
	u := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.TS, lp.Wait)

	req := fasthttp.AcquireRequest()
	req.URI().Update(u)
	req.Header.SetMethodBytes([]byte("GET"))

	res := fasthttp.AcquireResponse()

	res.StreamBody = true

	body := res.BodyStream()

	reader, err := gripDecoder.NewReader(body)
	if err != nil {
		return response, err
	}
	defer reader.Close()

	err = lp.Client.Do(req, res)
	if err != nil {
		return response, err
	}
	defer res.CloseBodyStream()

	response, err = parseResponse(reader)
	if err != nil {
		return response, err
	}

	err = lp.checkResponse(response)

	return response, err
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
		case "failed":
			raw, err := decoder.Token()
			if err != nil {
				return response, err
			}

			response.Failed = int(raw.(float64))
		case "updates":
			var updates []model.Updates

			err = decoder.Decode(&updates)
			if err != nil {
				return response, err
			}

			response.Updates = updates
		case "ts":
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

func (lp *LongPoll) checkResponse(response model.Response) (err error) {
	switch response.Failed {
	case 0:
		lp.TS = response.Ts
	case 1:
		lp.TS = response.Ts
	case 2:
		err = lp.getLongPoolParams(false)
	case 3:
		err = lp.getLongPoolParams(true)
	default:
		err = fmt.Errorf("longpool get error! Code: %d", response.Failed)
	}

	return
}

// Run handler.
func (lp *LongPoll) Run() error {
	return lp.RunWithContext(context.Background())
}

// RunWithContext handler.
func (lp *LongPoll) RunWithContext(ctx context.Context) error {
	return lp.run(ctx)
}

func (lp *LongPoll) run(ctx context.Context) error {
	ctx, lp.cancel = context.WithCancel(ctx)

	for {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				return nil
			}
		default:
			resp, err := lp.check()
			if err != nil {
				return err
			}

			// TODO: create longpool handler

			for _, f := range lp.funcFullResponseList {
				f(resp)
			}
		}
	}
}

// FullResponse handler.
func (lp *LongPoll) FullResponse(f func(response model.Response)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
