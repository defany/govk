package govk

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/messages/requests"
	requests2 "github.com/defany/govk/api/users/requests"
	"github.com/defany/govk/updates"
)

type apiParams struct {
	api *api.API

	Users    *requests2.Users
	Messages *requests.Messages
}

type VK struct {
	Api     apiParams
	Updates *updates.Updates
}

func NewVK(tokens ...string) (*VK, error) {
	api, err := api.NewAPI(tokens...)
	if err != nil {
		return nil, err
	}

	updates := updates.NewUpdates(api)

	return &VK{
		Api: apiParams{
			api:      api,
			Users:    requests2.NewUsers(api),
			Messages: requests.NewMessages(api),
		},
		Updates: updates,
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
