package requests

import (
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
)

type MessagesGetByIDResponse struct {
	Count int                `json:"count"`
	Items []msgmodel.Message `json:"items"`
	msgmodel.ExtendedResponse
}

type GetByIDParams api.Params

func NewGetByIDParams() GetByIDParams {
	params := make(GetByIDParams, 7)

	return params
}

func (g GetByIDParams) WithMessageIDs(messageIDs ...string) GetByIDParams {
	g[msgmodel.MessageIDsParam] = messageIDs
	return g
}

func (g GetByIDParams) WithPreviewLength(previewLength int) GetByIDParams {
	g[msgmodel.PreviewLengthParam] = previewLength

	return g
}

func (g GetByIDParams) WithExtended() GetByIDParams {
	g[msgmodel.ExtendedParam] = true

	return g
}

func (g GetByIDParams) WithFields(fields ...string) GetByIDParams {
	g[msgmodel.FieldsParam] = fields

	return g
}

func (g GetByIDParams) WithGroupID(groupID int) GetByIDParams {
	g[msgmodel.GroupIDParam] = groupID

	return g
}

func (g GetByIDParams) WithCmIDs(cmIDs string) GetByIDParams {
	g[msgmodel.CmIDsParam] = cmIDs

	return g
}

func (g GetByIDParams) WithPeerID(peerID int) GetByIDParams {
	g[msgmodel.PeerIDParam] = peerID

	return g
}

func (g GetByIDParams) Params() api.Params {
	return api.Params(g)
}

func (m *Messages) GetByID(params api.MethodParams) (MessagesGetByIDResponse, error) {
	req := api.NewRequest[MessagesGetByIDResponse](m.api)

	res, err := req.Execute(m.methodsGroup+"getById", params)
	if err != nil {
		return res, err
	}

	return res, nil
}
