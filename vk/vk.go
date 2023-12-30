package govk

import (
	"github.com/defany/govk/api"
	groupsReqs "github.com/defany/govk/api/groups/requests"
	messagesReqs "github.com/defany/govk/api/messages/requests"
	apiModel "github.com/defany/govk/api/model"
	usersReqs "github.com/defany/govk/api/users/requests"
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

	apiProvider := &apiModel.ApiProvider{
		Api:      api,
		Users:    usersReqs.NewUsers(api),
		Messages: messagesReqs.NewMessages(api),
		Groups:   groupsReqs.NewGroups(api),
	}

	updates := updates.NewUpdates(apiProvider)

	return &VK{
		Api:     apiProvider,
		Updates: updates,
	}, nil
}
