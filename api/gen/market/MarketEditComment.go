// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketEditComment Chages item comment's text
type MarketEditCommentRequest api.Params

func NewMarketEditCommentRequest() MarketEditCommentRequest {
	params := make(MarketEditCommentRequest, 5)
	return params
}

func (m MarketEditCommentRequest) WithOwnerId(m_owner_id int) MarketEditCommentRequest{
	m["owner_id"] = m_owner_id
	return m
}

func (m MarketEditCommentRequest) WithCommentId(m_comment_id int) MarketEditCommentRequest{
	m["comment_id"] = m_comment_id
	return m
}

func (m MarketEditCommentRequest) WithMessage(m_message string) MarketEditCommentRequest{
	m["message"] = m_message
	return m
}

func (m MarketEditCommentRequest) WithAttachments(m_attachments []string) MarketEditCommentRequest{
	m["attachments"] = m_attachments
	return m
}

func (m MarketEditCommentRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.editComment
func (m *Market) MarketEditComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("market.editComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
