// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PollsGetBackgrounds ...
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/polls.getBackgrounds
func (p *Polls) PollsGetBackgrounds(params ...api.MethodParams) (resp models.PollsGetBackgroundsResponse, err error) {
	req := api.NewRequest[models.PollsGetBackgroundsResponse](p.api)

	res, err := req.Execute("polls.getBackgrounds", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
