// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsEditList Edits a friend list of the current user.
type FriendsEditListRequest api.Params

func NewFriendsEditListRequest() FriendsEditListRequest {
	params := make(FriendsEditListRequest, 6)
	return params
}

func (f FriendsEditListRequest) WithName(f_name string) FriendsEditListRequest{
	f["name"] = f_name
	return f
}

func (f FriendsEditListRequest) WithListId(f_list_id int) FriendsEditListRequest{
	f["list_id"] = f_list_id
	return f
}

func (f FriendsEditListRequest) WithUserIds(f_user_ids []int) FriendsEditListRequest{
	f["user_ids"] = f_user_ids
	return f
}

func (f FriendsEditListRequest) WithAddUserIds(f_add_user_ids []int) FriendsEditListRequest{
	f["add_user_ids"] = f_add_user_ids
	return f
}

func (f FriendsEditListRequest) WithDeleteUserIds(f_delete_user_ids []int) FriendsEditListRequest{
	f["delete_user_ids"] = f_delete_user_ids
	return f
}

func (f FriendsEditListRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_FriendsListId ]
//
// https://dev.vk.com/method/friends.editList
func (f *Friends) FriendsEditList(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](f.api)

	res, err := req.Execute("friends.editList", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
