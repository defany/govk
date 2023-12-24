package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/users/model"
)

type GetRes []model.User

// TODO: add a builder

func (u *Users) Get(params ...api.MethodParams) (GetRes, error) {
	req := api.NewRequest[GetRes](u.api)

	res, err := req.Execute(u.methodsGroup+"get", api.ParamsOrNil(params))
	if err != nil {
		return nil, err
	}

	return res, nil
}
