// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
)

type Account struct {
	api *api.API
}

func NewAccount(api *api.API) *Account {
	return &Account{
		api: api,
	}
}
