package api

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
)

type Response[Res any] struct {
	Body  Res   `json:"response"`
	Error Error `json:"error"`
}

type Request[Res any] struct {
	response Response[Res]

	api *API

	retries uint

	retryNumber uint
}

func NewRequest[Res any](api *API) *Request[Res] {
	return &Request[Res]{
		api: api,

		retries: api.retries,
	}
}

func (r *Request[Res]) WithRetries(retries uint) *Request[Res] {
	r.retries = retries

	return r
}

func (r *Request[Res]) Execute(method string, params MethodParams) (Res, error) {
	r.api.mu.Lock()
	defer r.api.mu.Unlock()

	var token string
	if params == nil {
		params = Params{}

		params.Params()[versionParam] = r.api.apiVersion

		r.api.waitIfNeeded()
		token = r.api.token()
	} else {
		t, ok := params.Params()[tokenParam].(string)
		if !ok {
			r.api.waitIfNeeded()
			token = r.api.token()
		} else {
			token = t
		}

		if _, ok := params.Params()[versionParam]; !ok {
			params.Params()[versionParam] = Version
		}
	}
<<<<<<< HEAD
	reqBody, err := r.buildBody(params.Params())
	if err != nil {
		return r.response.Body, err
	}

	req, err := http.NewRequest(http.MethodPost, r.buildUrl(method), reqBody)
	if err != nil {
		return r.response.Body, err
	}

||||||| 06ea378
=======

>>>>>>> ec6cd9b5c35171d69d867df2d0a7cfd6f45a4d12
	r.setHeaders(req, token)

	res, err := r.api.http.Do(req)
	if err != nil {
		return r.response.Body, err
	}
	defer res.Body.Close()
	// TODO: parse by media type
	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		return r.response.Body, err
	}

	err = json.Unmarshal(respBody, &r.response.Body)
	if err != nil {
		return r.response.Body, err
	}

	switch {
	case r.response.Error.Is(ErrNoType):
		return r.response.Body, nil
	case r.response.Error.Is(ErrTooMany):
		return r.response.Body, &r.response.Error
	default:
		if r.retryNumber > r.retries {
			return r.Execute(method, params)
		}
		return r.response.Body, &r.response.Error
	}
}

func (r *Request[Res]) setHeaders(req *http.Request, token string) {
	req.Header.Set(fasthttp.HeaderAuthorization, "Bearer "+token)

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Accept-Encoding", gzip)
}

func (r *Request[Res]) buildUrl(method string) string {
	return r.api.apiURL + method
}

func (r *Request[Res]) buildBody(params Params) (io.Reader, error) {
	values, err := params.Build()
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBufferString(values.Encode())

	return reqBody, nil
}
