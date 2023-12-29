package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/groups/model"
)

type GetByIDRes model.GetByIDResponse

// TODO: add a builder

func (g *Groups) GetByID(params ...api.MethodParams) (GetByIDRes, error) {
	req := api.NewRequest[GetByIDRes](g.api)

	res, err := req.Execute(g.methodsGroup+"getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
