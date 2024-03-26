// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesSetActivity Changes the status of a user as typing in a conversation.
type MessagesSetActivityRequest api.Params

func NewMessagesSetActivityRequest() MessagesSetActivityRequest {
	params := make(MessagesSetActivityRequest, 5)
	return params
}

func (m MessagesSetActivityRequest) WithUserId(m_user_id int) MessagesSetActivityRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MessagesSetActivityRequest) WithType(m_type string) MessagesSetActivityRequest{
	m["type"] = m_type
	return m
}

func (m MessagesSetActivityRequest) WithPeerId(m_peer_id int) MessagesSetActivityRequest{
	m["peer_id"] = m_peer_id
	return m
}

func (m MessagesSetActivityRequest) WithGroupId(m_group_id int) MessagesSetActivityRequest{
	m["group_id"] = m_group_id
	return m
}

func (m MessagesSetActivityRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesGroupPeerAccess, Error_MessagesChatUserNoAccess, Error_MessagesContactNotFound ]
//
// https://dev.vk.com/method/messages.setActivity
func (m *Messages) MessagesSetActivity(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("messages.setActivity", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
