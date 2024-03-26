// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsSearch Returns a list of friends matching the search criteria.
type FriendsSearchRequest api.Params

func NewFriendsSearchRequest() FriendsSearchRequest {
	params := make(FriendsSearchRequest, 7)
	return params
}

func (f FriendsSearchRequest) WithUserId(f_user_id int) FriendsSearchRequest{
	f["user_id"] = f_user_id
	return f
}

func (f FriendsSearchRequest) WithQ(f_q string) FriendsSearchRequest{
	f["q"] = f_q
	return f
}

func (f FriendsSearchRequest) WithFields(f_fields []models.UsersFields) FriendsSearchRequest{
	f["fields"] = f_fields
	return f
}

func (f FriendsSearchRequest) WithNameCase(f_name_case string) FriendsSearchRequest{
	f["name_case"] = f_name_case
	return f
}

func (f FriendsSearchRequest) WithOffset(f_offset int) FriendsSearchRequest{
	f["offset"] = f_offset
	return f
}

func (f FriendsSearchRequest) WithCount(f_count int) FriendsSearchRequest{
	f["count"] = f_count
	return f
}

func (f FriendsSearchRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/friends.search
func (f *Friends) FriendsSearch(params ...api.MethodParams) (resp models.FriendsSearchResponse, err error) {
	req := api.NewRequest[models.FriendsSearchResponse](f.api)

	res, err := req.Execute("friends.search", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

