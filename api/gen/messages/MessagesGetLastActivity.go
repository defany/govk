// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesGetLastActivity Returns a user's current status and date of last activity.
type MessagesGetLastActivityRequest api.Params

func NewMessagesGetLastActivityRequest() MessagesGetLastActivityRequest {
	params := make(MessagesGetLastActivityRequest, 2)
	return params
}

func (m MessagesGetLastActivityRequest) WithUserId(m_user_id int) MessagesGetLastActivityRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MessagesGetLastActivityRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/messages.getLastActivity
func (m *Messages) MessagesGetLastActivity(params ...api.MethodParams) (resp models.MessagesGetLastActivityResponse, err error) {
	req := api.NewRequest[models.MessagesGetLastActivityResponse](m.api)

	res, err := req.Execute("messages.getLastActivity", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

