// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FaveAddLink Adds a link to user faves.
type FaveAddLinkRequest api.Params

func NewFaveAddLinkRequest() FaveAddLinkRequest {
	params := make(FaveAddLinkRequest, 2)
	return params
}

func (f FaveAddLinkRequest) WithLink(f_link string) FaveAddLinkRequest{
	f["link"] = f_link
	return f
}

func (f FaveAddLinkRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/fave.addLink
func (f *Fave) FaveAddLink(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](f.api)

	res, err := req.Execute("fave.addLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
