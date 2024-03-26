// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsEditCallbackServer ...
type GroupsEditCallbackServerRequest api.Params

func NewGroupsEditCallbackServerRequest() GroupsEditCallbackServerRequest {
	params := make(GroupsEditCallbackServerRequest, 6)
	return params
}

func (g GroupsEditCallbackServerRequest) WithGroupId(g_group_id int) GroupsEditCallbackServerRequest{
	g["group_id"] = g_group_id
	return g
}

func (g GroupsEditCallbackServerRequest) WithServerId(g_server_id int) GroupsEditCallbackServerRequest{
	g["server_id"] = g_server_id
	return g
}

func (g GroupsEditCallbackServerRequest) WithUrl(g_url string) GroupsEditCallbackServerRequest{
	g["url"] = g_url
	return g
}

func (g GroupsEditCallbackServerRequest) WithTitle(g_title string) GroupsEditCallbackServerRequest{
	g["title"] = g_title
	return g
}

func (g GroupsEditCallbackServerRequest) WithSecretKey(g_secret_key string) GroupsEditCallbackServerRequest{
	g["secret_key"] = g_secret_key
	return g
}

func (g GroupsEditCallbackServerRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_NotFound ]
//
// https://dev.vk.com/method/groups.editCallbackServer
func (g *Groups) GroupsEditCallbackServer(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](g.api)

	res, err := req.Execute("groups.editCallbackServer", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

