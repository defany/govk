package govk

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/messages"
	"github.com/defany/govk/api/users"
)

type apiParams struct {
	api      *api.API
	Users    *users.Users
	Messages *messages.Messages
}

type VK struct {
	Api apiParams
}

func NewVK(tokens ...string) (*VK, error) {
	api, err := api.NewAPI(tokens...)
	if err != nil {
		return nil, err
	}

	return &VK{
		Api: apiParams{
			api:      api,
			Users:    users.NewUsers(api),
			Messages: messages.NewMessages(api),
		},
	}, nil
}

func (v *VK) WithApiVersion(version string) *VK {
	v.Api.api.WithVersion(version)

	return v
}

func (v *VK) WithApiURL(url string) *VK {
	v.Api.api.WithApiURL(url)

	return v
}

func (v *VK) WithApiLimit(limit uint) *VK {
	v.Api.api.WithLimit(limit)

	return v
}
