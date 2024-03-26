// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdswebGetAdUnitCode ...
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getAdUnitCode
func (a *Adsweb) AdswebGetAdUnitCode(params ...api.MethodParams) (resp models.AdswebGetAdUnitCodeResponse, err error) {
	req := api.NewRequest[models.AdswebGetAdUnitCodeResponse](a.api)

	res, err := req.Execute("adsweb.getAdUnitCode", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
