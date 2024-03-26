// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesRemoveChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
type MessagesRemoveChatUserRequest api.Params

func NewMessagesRemoveChatUserRequest() MessagesRemoveChatUserRequest {
	params := make(MessagesRemoveChatUserRequest, 4)
	return params
}

func (m MessagesRemoveChatUserRequest) WithChatId(m_chat_id int) MessagesRemoveChatUserRequest{
	m["chat_id"] = m_chat_id
	return m
}

func (m MessagesRemoveChatUserRequest) WithUserId(m_user_id int) MessagesRemoveChatUserRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MessagesRemoveChatUserRequest) WithMemberId(m_member_id int) MessagesRemoveChatUserRequest{
	m["member_id"] = m_member_id
	return m
}

func (m MessagesRemoveChatUserRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesChatNotAdmin, Error_MessagesChatUserNotInChat, Error_MessagesContactNotFound, Error_MessagesChatDisabled, Error_MessagesChatUnsupported ]
//
// https://dev.vk.com/method/messages.removeChatUser
func (m *Messages) MessagesRemoveChatUser(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("messages.removeChatUser", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
