package requests

import (
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
)

type SendMessageEventAnswerParams api.Params

func NewSendMessageEventAnswerParams() SendMessageEventAnswerParams {
	params := make(SendMessageEventAnswerParams, 4)

	return params
}

func (s SendMessageEventAnswerParams) WithEventID(eventID string) SendMessageEventAnswerParams {
	s[msgmodel.EventIDParam] = eventID

	return s
}

func (s SendMessageEventAnswerParams) WithUserID(userID int) SendMessageEventAnswerParams {
	s[msgmodel.UserIDParam] = userID

	return s
}

func (s SendMessageEventAnswerParams) WithPeerID(peerID int) SendMessageEventAnswerParams {
	s[msgmodel.PeerIDParam] = peerID

	return s
}

func (s SendMessageEventAnswerParams) WithEventData(eventData string) SendMessageEventAnswerParams {
	s[msgmodel.EventDataParam] = eventData

	return s
}

func (s SendMessageEventAnswerParams) Params() api.Params {
	return api.Params(s)
}

func (m *Messages) SendMessageEventAnswer(params api.MethodParams) (int, error) {
	req := api.NewRequest[int](m.api)

	res, err := req.Execute(m.methodsGroup+"sendMessageEventAnswer", params)
	if err != nil {
		return res, err
	}

	return res, nil
}
