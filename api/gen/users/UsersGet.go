// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// UsersGet Returns detailed information on users.
type UsersGetRequest api.Params

func NewUsersGetRequest() UsersGetRequest {
	params := make(UsersGetRequest, 4)
	return params
}

func (u UsersGetRequest) WithUserIds(u_user_ids []string) UsersGetRequest{
	u["user_ids"] = u_user_ids
	return u
}

func (u UsersGetRequest) WithFields(u_fields []models.UsersFields) UsersGetRequest{
	u["fields"] = u_fields
	return u
}

func (u UsersGetRequest) WithNameCase(u_name_case string) UsersGetRequest{
	u["name_case"] = u_name_case
	return u
}

func (u UsersGetRequest) Params() api.Params {
	return api.Params(u)
}

// May execute with listed access token types:
//    [ user, group, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/users.get
func (u *Users) UsersGet(params ...api.MethodParams) (resp models.UsersGetResponse, err error) {
	req := api.NewRequest[models.UsersGetResponse](u.api)

	res, err := req.Execute("users.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

