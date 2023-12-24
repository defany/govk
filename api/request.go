package api

import (
	gripDecoder "compress/gzip"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

type Response[Res any] struct {
	Body  Res   `json:"response"`
	Error Error `json:"error"`
}

type Request[Res any] struct {
	response Response[Res]

	api *API

	retries uint
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

// TODO: add retries
func (r *Request[Res]) Execute(method string, params MethodParams) (Res, error) {
	r.api.mu.Lock()
	defer r.api.mu.Unlock()

	r.api.waitIfNeeded()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	r.setHeaders(req)

	r.buildUrl(method, req)

	if _, ok := params.Params()[versionParam]; !ok {
		params.Params()[versionParam] = Version
	}

	if params != nil {
		if err := r.buildBody(req, params.Params()); err != nil {
			return r.response.Body, err
		}
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	res.StreamBody = true

	err := r.api.http.Do(req, res)
	if err != nil {
		return r.response.Body, err
	}

	// TODO: parse by media type
	body := res.BodyStream()

	reader, err := gripDecoder.NewReader(body)
	if err != nil {
		return r.response.Body, err
	}
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&r.response)
	if err != nil {
		return r.response.Body, err
	}

	switch {
	case r.response.Error.Is(ErrNoType):
		return r.response.Body, nil
	case r.response.Error.Is(ErrTooMany):
		return r.response.Body, &r.response.Error
	default:
		return r.response.Body, &r.response.Error
	}
}

func (r *Request[Res]) setHeaders(req *fasthttp.Request) {
	req.Header.SetMethod(fasthttp.MethodPost)

	req.Header.Set(fasthttp.HeaderAuthorization, "Bearer "+r.api.token())

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Accept-Encoding", gzip)
}

func (r *Request[Res]) buildUrl(method string, req *fasthttp.Request) {
	req.URI().Update(r.api.apiURL + method)
}

func (r *Request[Res]) buildBody(req *fasthttp.Request, params Params) error {
	values, err := params.Build()
	if err != nil {
		return err
	}

	req.SetBodyString(values.Encode())

	return nil
}
