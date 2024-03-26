// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsAddList Creates a new friend list for the current user.
type FriendsAddListRequest api.Params

func NewFriendsAddListRequest() FriendsAddListRequest {
	params := make(FriendsAddListRequest, 3)
	return params
}

func (f FriendsAddListRequest) WithName(f_name string) FriendsAddListRequest{
	f["name"] = f_name
	return f
}

func (f FriendsAddListRequest) WithUserIds(f_user_ids []int) FriendsAddListRequest{
	f["user_ids"] = f_user_ids
	return f
}

func (f FriendsAddListRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_FriendsListLimit ]
//
// https://dev.vk.com/method/friends.addList
func (f *Friends) FriendsAddList(params ...api.MethodParams) (resp models.FriendsAddListResponse, err error) {
	req := api.NewRequest[models.FriendsAddListResponse](f.api)

	res, err := req.Execute("friends.addList", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
