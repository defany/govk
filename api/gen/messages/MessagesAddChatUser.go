// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesAddChatUser Adds a new user to a chat.
type MessagesAddChatUserRequest api.Params

func NewMessagesAddChatUserRequest() MessagesAddChatUserRequest {
	params := make(MessagesAddChatUserRequest, 4)
	return params
}

func (m MessagesAddChatUserRequest) WithChatId(m_chat_id int) MessagesAddChatUserRequest{
	m["chat_id"] = m_chat_id
	return m
}

func (m MessagesAddChatUserRequest) WithUserId(m_user_id int) MessagesAddChatUserRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MessagesAddChatUserRequest) WithVisibleMessagesCount(m_visible_messages_count int) MessagesAddChatUserRequest{
	m["visible_messages_count"] = m_visible_messages_count
	return m
}

func (m MessagesAddChatUserRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_Limits, Error_MessagesChatNotAdmin, Error_MessagesMessageRequestAlreadySent, Error_MessagesContactNotFound, Error_MessagesChatDisabled, Error_MessagesMemberAccessToGroupDenied, Error_MessagesChatUnsupported, Error_MessagesGroupPeerAccess ]
//
// https://dev.vk.com/method/messages.addChatUser
func (m *Messages) MessagesAddChatUser(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("messages.addChatUser", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
