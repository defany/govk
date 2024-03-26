// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AccountUnban ...
type AccountUnbanRequest api.Params

func NewAccountUnbanRequest() AccountUnbanRequest {
	params := make(AccountUnbanRequest, 2)
	return params
}

func (a AccountUnbanRequest) WithOwnerId(a_owner_id int) AccountUnbanRequest{
	a["owner_id"] = a_owner_id
	return a
}

func (a AccountUnbanRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/account.unban
func (a *Account) AccountUnban(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](a.api)

	res, err := req.Execute("account.unban", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
