// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesUnpin ...
type MessagesUnpinRequest api.Params

func NewMessagesUnpinRequest() MessagesUnpinRequest {
	params := make(MessagesUnpinRequest, 3)
	return params
}

func (m MessagesUnpinRequest) WithPeerId(m_peer_id int) MessagesUnpinRequest{
	m["peer_id"] = m_peer_id
	return m
}

func (m MessagesUnpinRequest) WithGroupId(m_group_id int) MessagesUnpinRequest{
	m["group_id"] = m_group_id
	return m
}

func (m MessagesUnpinRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesChatNotAdmin ]
//
// https://dev.vk.com/method/messages.unpin
func (m *Messages) MessagesUnpin(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("messages.unpin", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

