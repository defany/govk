package requests

import (
	"github.com/defany/govk/api"
)

// TODO: add a builder

func (m *Messages) Send(params api.MethodParams) (int, error) {
	req := api.NewRequest[int](m.api)

	res, err := req.Execute(m.methodsGroup+"send", params)
	if err != nil {
		return 0, err
	}

	return res, nil
}
