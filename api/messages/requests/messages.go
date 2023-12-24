package requests

import "github.com/defany/govk/api"

const methodsGroup = "messages."

type Messages struct {
	methodsGroup string

	api *api.API
}

func NewMessages(api *api.API) *Messages {
	return &Messages{
		methodsGroup: methodsGroup,
		api:          api,
	}
}
