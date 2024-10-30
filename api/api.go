package api

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type API struct {
	tokens []string

	currentToken uint32

	apiVersion string
	apiURL     string

	retries uint
	limit   uint

	mu          sync.Mutex
	tokenUsedAt time.Time
	rps         uint

	http *http.Client
}

func NewAPI(tokens ...string) (*API, error) {
	if len(tokens) == 0 {
		return nil, ErrTokensNotProvided
	}

	api := &API{
		limit:      maxGroupApiRequestsCalls,
		retries:    retries,
		apiVersion: Version,
		apiURL:     Url,
		tokens:     tokens,
		http:       &http.Client{},
	}

	return api, nil
}

// WithVersion - устанавливает версию api.
// Дефолтное значение: 5.199
func (a *API) WithVersion(version string) *API {
	a.apiVersion = version

	return a
}

// WithApiURL - устанавливает ссылку на api по которой отсылаются эти запросы.
func (a *API) WithApiURL(url string) *API {
	a.apiURL = url

	return a
}

// WithLimit - устанавливает максимальное количество запросов за секунду
func (a *API) WithLimit(limit uint) *API {
	a.limit = limit

	return a
}

// WithMaxRetries - максимальное количество попыток повтора запроса.
// Дефолтное значение: 3
// Если установить 0, то после первого же запроса будет возвращен результат.
func (a *API) WithMaxRetries(retries uint) *API {
	a.retries = retries

	return a
}

func (a *API) WithHTTP(client *http.Client) *API {
	a.http = client

	return a
}

func (a *API) token() string {
	i := atomic.AddUint32(&a.currentToken, 1)

	a.rps++

	return a.tokens[(int(i)-1)%len(a.tokens)]
}

func (a *API) waitIfNeeded() {
	sleepTime := time.Second - time.Since(a.tokenUsedAt)
	if sleepTime < 0 {
		a.tokenUsedAt = time.Now()
		a.rps = 0
	} else if a.rps == a.limit*uint(len(a.tokens)) {
		time.Sleep(sleepTime)

		a.tokenUsedAt = time.Now()
		a.rps = 0
	}
}
