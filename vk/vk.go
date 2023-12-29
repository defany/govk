package govk

import (
	"github.com/defany/govk/api"
	requests3 "github.com/defany/govk/api/groups/requests"
	"github.com/defany/govk/api/messages/requests"
	apiModel "github.com/defany/govk/api/model"
	requests2 "github.com/defany/govk/api/users/requests"
	"github.com/defany/govk/updates"
)

type VK struct {
	Api     *apiModel.ApiProvider
	Updates *updates.Updates
}

func NewVK(tokens ...string) (*VK, error) {
	api, err := api.NewAPI(tokens...)
	if err != nil {
		return nil, err
	}

	var apiProvider *apiModel.ApiProvider

	apiProvider = &apiModel.ApiProvider{
		Api:      api,
		Users:    requests2.NewUsers(api),
		Messages: requests.NewMessages(api),
		Groups:   requests3.NewGroups(api),
	}

	updates := updates.NewUpdates(apiProvider)

	return &VK{
		Api:     apiProvider,
		Updates: updates,
	}, nil
}

func (v *VK) WithApiVersion(version string) *VK {
	v.Api.Api.WithVersion(version)

	return v
}

func (v *VK) WithApiURL(url string) *VK {
	v.Api.Api.WithApiURL(url)

	return v
}

func (v *VK) WithApiLimit(limit uint) *VK {
	v.Api.Api.WithLimit(limit)

	return v
}
