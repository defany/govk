// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesSendMessageEventAnswer ...
type MessagesSendMessageEventAnswerRequest api.Params

func NewMessagesSendMessageEventAnswerRequest() MessagesSendMessageEventAnswerRequest {
	params := make(MessagesSendMessageEventAnswerRequest, 5)
	return params
}

func (m MessagesSendMessageEventAnswerRequest) WithEventId(m_event_id string) MessagesSendMessageEventAnswerRequest{
	m["event_id"] = m_event_id
	return m
}

func (m MessagesSendMessageEventAnswerRequest) WithUserId(m_user_id int) MessagesSendMessageEventAnswerRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MessagesSendMessageEventAnswerRequest) WithPeerId(m_peer_id int) MessagesSendMessageEventAnswerRequest{
	m["peer_id"] = m_peer_id
	return m
}

func (m MessagesSendMessageEventAnswerRequest) WithEventData(m_event_data string) MessagesSendMessageEventAnswerRequest{
	m["event_data"] = m_event_data
	return m
}

func (m MessagesSendMessageEventAnswerRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ group ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/messages.sendMessageEventAnswer
func (m *Messages) MessagesSendMessageEventAnswer(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("messages.sendMessageEventAnswer", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

