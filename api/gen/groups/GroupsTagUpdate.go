// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsTagUpdate Update group's tag
type GroupsTagUpdateRequest api.Params

func NewGroupsTagUpdateRequest() GroupsTagUpdateRequest {
	params := make(GroupsTagUpdateRequest, 4)
	return params
}

func (g GroupsTagUpdateRequest) WithGroupId(g_group_id int) GroupsTagUpdateRequest{
	g["group_id"] = g_group_id
	return g
}

func (g GroupsTagUpdateRequest) WithTagId(g_tag_id int) GroupsTagUpdateRequest{
	g["tag_id"] = g_tag_id
	return g
}

func (g GroupsTagUpdateRequest) WithTagName(g_tag_name string) GroupsTagUpdateRequest{
	g["tag_name"] = g_tag_name
	return g
}

func (g GroupsTagUpdateRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/groups.tagUpdate
func (g *Groups) GroupsTagUpdate(params ...api.MethodParams) (resp models.BaseBoolResponse, err error) {
	req := api.NewRequest[models.BaseBoolResponse](g.api)

	res, err := req.Execute("groups.tagUpdate", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

