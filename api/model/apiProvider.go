package model

import (
	"github.com/defany/govk/api"
	groupsReqs "github.com/defany/govk/api/groups/requests"
	messagesReqs "github.com/defany/govk/api/messages/requests"
	userReqs "github.com/defany/govk/api/users/requests"
)

type ApiProvider struct {
	Api *api.API

	Users    *userReqs.Users
	Messages *messagesReqs.Messages
	Groups   *groupsReqs.Groups
}

func (a *ApiProvider) WithLimit(limit uint) *ApiProvider {
	a.Api.WithLimit(limit)

	return a
}

func (a *ApiProvider) WithUrl(url string) *ApiProvider {
	a.Api.WithApiURL(url)

	return a
}

func (a *ApiProvider) WithVersion(v string) *ApiProvider {
	a.Api.WithVersion(v)

	return a
}

func (a *ApiProvider) WithMaxRetries(retries uint) *ApiProvider {
	a.Api.WithMaxRetries(retries)

	return a
}
