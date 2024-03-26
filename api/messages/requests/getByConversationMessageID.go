package requests

import (
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
)

type GetByConversationMessageIDResponse struct {
	Count int                `json:"count"`
	Items []msgmodel.Message `json:"items"`
	msgmodel.ExtendedResponse
}

type GetByConversationMessageIDParams api.Params

func NewGetByConversationMessageIDParams() GetByConversationMessageIDParams {
	params := make(GetByConversationMessageIDParams, 5)

	return params
}

func (g GetByConversationMessageIDParams) WithPeerID(peerID int) GetByConversationMessageIDParams {
	g[msgmodel.PeerIDParam] = peerID

	return g
}

func (g GetByConversationMessageIDParams) WithConversationMessageIDs(conversationMessageIDs ...string) GetByConversationMessageIDParams {
	g[msgmodel.ConversationMessageIDs] = conversationMessageIDs

	return g
}

func (g GetByConversationMessageIDParams) WithExtended() GetByConversationMessageIDParams {
	g[msgmodel.ExtendedParam] = true

	return g
}

func (g GetByConversationMessageIDParams) WithFields(fields ...string) GetByConversationMessageIDParams {
	g[msgmodel.FieldsParam] = fields

	return g
}

func (g GetByConversationMessageIDParams) WithGroupID(groupID int) GetByConversationMessageIDParams {
	g[msgmodel.GroupIDParam] = groupID

	return g
}

func (g GetByConversationMessageIDParams) Params() api.Params {
	return api.Params(g)
}

func (m *Messages) GetByConversationMessageID(params api.MethodParams) (GetByConversationMessageIDResponse, error) {
	req := api.NewRequest[GetByConversationMessageIDResponse](m.api)

	res, err := req.Execute(m.methodsGroup+"getByConversationMessageId", params)
	if err != nil {
		return res, err
	}

	return res, nil
}
