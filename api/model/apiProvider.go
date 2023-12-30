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
