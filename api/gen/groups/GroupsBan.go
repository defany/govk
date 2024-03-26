// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsBan ...
type GroupsBanRequest api.Params

func NewGroupsBanRequest() GroupsBanRequest {
	params := make(GroupsBanRequest, 7)
	return params
}

func (g GroupsBanRequest) WithGroupId(g_group_id int) GroupsBanRequest{
	g["group_id"] = g_group_id
	return g
}

func (g GroupsBanRequest) WithOwnerId(g_owner_id int) GroupsBanRequest{
	g["owner_id"] = g_owner_id
	return g
}

func (g GroupsBanRequest) WithEndDate(g_end_date int) GroupsBanRequest{
	g["end_date"] = g_end_date
	return g
}

func (g GroupsBanRequest) WithReason(g_reason int) GroupsBanRequest{
	g["reason"] = g_reason
	return g
}

func (g GroupsBanRequest) WithComment(g_comment string) GroupsBanRequest{
	g["comment"] = g_comment
	return g
}

func (g GroupsBanRequest) WithCommentVisible(g_comment_visible bool) GroupsBanRequest{
	g["comment_visible"] = g_comment_visible
	return g
}

func (g GroupsBanRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/groups.ban
func (g *Groups) GroupsBan(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](g.api)

	res, err := req.Execute("groups.ban", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
