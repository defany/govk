package model

import (
	"github.com/defany/govk/api"
	requests3 "github.com/defany/govk/api/groups/requests"
	"github.com/defany/govk/api/messages/requests"
	requests2 "github.com/defany/govk/api/users/requests"
)

type ApiProvider struct {
	Api *api.API

	Users    *requests2.Users
	Messages *requests.Messages
	Groups   *requests3.Groups
}
