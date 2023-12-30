package requests

import (
	"github.com/defany/govk/api"
)

type GetLongPollServerRes struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     string `json:"ts"`
}

// TODO: add a builder

func (g *Groups) GetLongPollServer(params api.MethodParams) (GetLongPollServerRes, error) {
	req := api.NewRequest[GetLongPollServerRes](g.api)

	res, err := req.Execute(g.methodsGroup+"getLongPollServer", params)
	if err != nil {
		return res, err
	}

	return res, nil
}
