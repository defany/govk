package msgmodel

import (
	groupmodel "github.com/defany/govk/api/groups/model"
	usermodel "github.com/defany/govk/api/users/model"
)

type ExtendedResponse struct {
	Profiles []usermodel.User   `json:"profiles,omitempty"`
	Groups   []groupmodel.Group `json:"groups,omitempty"`
}
