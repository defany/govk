package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/groups/model"
	usermodel "github.com/defany/govk/api/users/model"
)

type GetByIDRes struct {
	Groups   []groupmodel.Group `json:"groups"`
	Profiles []usermodel.User   `json:"profiles"`
}

// TODO: add a builder

func (g *Groups) GetByID(params ...api.MethodParams) (GetByIDRes, error) {
	req := api.NewRequest[GetByIDRes](g.api)

	res, err := req.Execute(g.methodsGroup+"getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
