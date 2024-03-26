// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsDeleteList Deletes a friend list of the current user.
type FriendsDeleteListRequest api.Params

func NewFriendsDeleteListRequest() FriendsDeleteListRequest {
	params := make(FriendsDeleteListRequest, 2)
	return params
}

func (f FriendsDeleteListRequest) WithListId(f_list_id int) FriendsDeleteListRequest{
	f["list_id"] = f_list_id
	return f
}

func (f FriendsDeleteListRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_FriendsListId ]
//
// https://dev.vk.com/method/friends.deleteList
func (f *Friends) FriendsDeleteList(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](f.api)

	res, err := req.Execute("friends.deleteList", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

