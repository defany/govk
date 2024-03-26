// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketRestoreComment Restores a recently deleted comment
type MarketRestoreCommentRequest api.Params

func NewMarketRestoreCommentRequest() MarketRestoreCommentRequest {
	params := make(MarketRestoreCommentRequest, 3)
	return params
}

func (m MarketRestoreCommentRequest) WithOwnerId(m_owner_id int) MarketRestoreCommentRequest{
	m["owner_id"] = m_owner_id
	return m
}

func (m MarketRestoreCommentRequest) WithCommentId(m_comment_id int) MarketRestoreCommentRequest{
	m["comment_id"] = m_comment_id
	return m
}

func (m MarketRestoreCommentRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.restoreComment
func (m *Market) MarketRestoreComment(params ...api.MethodParams) (resp models.MarketRestoreCommentResponse, err error) {
	req := api.NewRequest[models.MarketRestoreCommentResponse](m.api)

	res, err := req.Execute("market.restoreComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
