package requests

import (
	"github.com/defany/govk/api"
)

const methodsGroup = "users."

type Users struct {
	methodsGroup string

	api *api.API
}

func NewUsers(api *api.API) *Users {
	return &Users{
		methodsGroup: methodsGroup,
		api:          api,
	}
}
