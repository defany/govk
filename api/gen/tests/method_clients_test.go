package tests

import (
	"bytes"
	"github.com/defany/govk/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/url"
	"testing"
)

const (
	apiScheme = "https"
	apiHost   = "api.vk.com"
	apiPath   = "method"

	versionKey = "v"
)

type TestRoundTrip func(*http.Request) (*http.Response, error)

func (f TestRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func NewTestClient(t *testing.T, methodName string, params api.Params, responseBodyRaw []byte) *http.Client {
	transport := TestRoundTrip(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, apiScheme, req.URL.Scheme)
		assert.Equal(t, apiHost, req.URL.Host)
		assert.Equal(t, "/"+apiPath+"/"+methodName, req.URL.Path)

		if params != nil {
			requestValues, err := params.Build()
			require.NoError(t, err)
			reqBody, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			values, err := url.ParseQuery(string(reqBody))
			require.NoError(t, err)

			requestValues.Set(versionKey, api.Version)

			assert.EqualValues(t, requestValues, values)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBuffer(responseBodyRaw)),
		}, nil
	})

	return &http.Client{
		Transport: transport,
	}
}
