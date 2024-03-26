// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesJoinChatByInviteLink ...
type MessagesJoinChatByInviteLinkRequest api.Params

func NewMessagesJoinChatByInviteLinkRequest() MessagesJoinChatByInviteLinkRequest {
	params := make(MessagesJoinChatByInviteLinkRequest, 2)
	return params
}

func (m MessagesJoinChatByInviteLinkRequest) WithLink(m_link string) MessagesJoinChatByInviteLinkRequest{
	m["link"] = m_link
	return m
}

func (m MessagesJoinChatByInviteLinkRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesChatUserNoAccess, Error_Limits ]
//
// https://dev.vk.com/method/messages.joinChatByInviteLink
func (m *Messages) MessagesJoinChatByInviteLink(params ...api.MethodParams) (resp models.MessagesJoinChatByInviteLinkResponse, err error) {
	req := api.NewRequest[models.MessagesJoinChatByInviteLinkResponse](m.api)

	res, err := req.Execute("messages.joinChatByInviteLink", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

