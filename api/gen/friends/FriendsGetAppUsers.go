// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsGetAppUsers Returns a list of IDs of the current user's friends who installed the application.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/friends.getAppUsers
func (f *Friends) FriendsGetAppUsers(params ...api.MethodParams) (resp models.FriendsGetAppUsersResponse, err error) {
	req := api.NewRequest[models.FriendsGetAppUsersResponse](f.api)

	res, err := req.Execute("friends.getAppUsers", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

