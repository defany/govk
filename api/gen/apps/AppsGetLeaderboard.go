// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AppsGetLeaderboard Returns players rating in the game.
type AppsGetLeaderboardRequest api.Params

func NewAppsGetLeaderboardRequest() AppsGetLeaderboardRequest {
	params := make(AppsGetLeaderboardRequest, 5)
	return params
}

func (a AppsGetLeaderboardRequest) WithType(a_type string) AppsGetLeaderboardRequest{
	a["type"] = a_type
	return a
}

func (a AppsGetLeaderboardRequest) WithGlobal(a_global bool) AppsGetLeaderboardRequest{
	a["global"] = a_global
	return a
}

func (a AppsGetLeaderboardRequest) WithExtended(a_extended bool) AppsGetLeaderboardRequest{
	a["extended"] = a_extended
	return a
}

func (a AppsGetLeaderboardRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/apps.getLeaderboard
func (a *Apps) AppsGetLeaderboard(params ...api.MethodParams) (resp models.AppsGetLeaderboardResponse, err error) {
	req := api.NewRequest[models.AppsGetLeaderboardResponse](a.api)

	res, err := req.Execute("apps.getLeaderboard", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AppsGetLeaderboardExtended Returns players rating in the game.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/apps.getLeaderboard
func (a *Apps) AppsGetLeaderboardExtended(params ...api.MethodParams) (resp models.AppsGetLeaderboardExtendedResponse, err error) {
	req := api.NewRequest[models.AppsGetLeaderboardExtendedResponse](a.api)

	res, err := req.Execute("apps.getLeaderboard", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
