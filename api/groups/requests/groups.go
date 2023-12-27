package requests

import "github.com/defany/govk/api"

const methodsGroup = "groups."

type Groups struct {
	methodsGroup string

	api *api.API
}

func NewGroups(api *api.API) *Groups {
	return &Groups{
		methodsGroup: methodsGroup,
		api:          api,
	}
}
