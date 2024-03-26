// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// UtilsGetShortLink Allows to receive a link shortened via vk.cc.
type UtilsGetShortLinkRequest api.Params

func NewUtilsGetShortLinkRequest() UtilsGetShortLinkRequest {
	params := make(UtilsGetShortLinkRequest, 3)
	return params
}

func (u UtilsGetShortLinkRequest) WithUrl(u_url string) UtilsGetShortLinkRequest{
	u["url"] = u_url
	return u
}

func (u UtilsGetShortLinkRequest) WithPrivate(u_private bool) UtilsGetShortLinkRequest{
	u["private"] = u_private
	return u
}

func (u UtilsGetShortLinkRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//    [ user, group, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/utils.getShortLink
func (u *Utils) UtilsGetShortLink(params ...api.MethodParams) (resp models.UtilsGetShortLinkResponse, err error) {
	req := api.NewRequest[models.UtilsGetShortLinkResponse](u.api)

	res, err := req.Execute("utils.getShortLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

