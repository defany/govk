package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/groups/model"
)

type GetRes model.LongPollServer

// TODO: add a builder

func (g *Groups) GetLongPollServer(params ...api.MethodParams) (GetRes, error) {
	req := api.NewRequest[GetRes](g.api)

	res, err := req.Execute(g.methodsGroup+"getLongPollServer", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
