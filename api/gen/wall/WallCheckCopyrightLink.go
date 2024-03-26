// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// WallCheckCopyrightLink ...
type WallCheckCopyrightLinkRequest api.Params

func NewWallCheckCopyrightLinkRequest() WallCheckCopyrightLinkRequest {
	params := make(WallCheckCopyrightLinkRequest, 2)
	return params
}

func (w WallCheckCopyrightLinkRequest) WithLink(w_link string) WallCheckCopyrightLinkRequest{
	w["link"] = w_link
	return w
}

func (w WallCheckCopyrightLinkRequest) Params() api.Params {
	return api.Params(w)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallCheckLinkCantDetermineSource ]
//
// https://dev.vk.com/method/wall.checkCopyrightLink
func (w *Wall) WallCheckCopyrightLink(params ...api.MethodParams) (resp models.BaseBoolResponse, err error) {
	req := api.NewRequest[models.BaseBoolResponse](w.api)

	res, err := req.Execute("wall.checkCopyrightLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

