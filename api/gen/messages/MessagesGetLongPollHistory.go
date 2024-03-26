// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MessagesGetLongPollHistory Returns updates in user's private messages.
type MessagesGetLongPollHistoryRequest api.Params

func NewMessagesGetLongPollHistoryRequest() MessagesGetLongPollHistoryRequest {
	params := make(MessagesGetLongPollHistoryRequest, 14)
	return params
}

func (m MessagesGetLongPollHistoryRequest) WithTs(m_ts int) MessagesGetLongPollHistoryRequest{
	m["ts"] = m_ts
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithPts(m_pts int) MessagesGetLongPollHistoryRequest{
	m["pts"] = m_pts
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithPreviewLength(m_preview_length int) MessagesGetLongPollHistoryRequest{
	m["preview_length"] = m_preview_length
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithOnlines(m_onlines bool) MessagesGetLongPollHistoryRequest{
	m["onlines"] = m_onlines
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithFields(m_fields []models.UsersFields) MessagesGetLongPollHistoryRequest{
	m["fields"] = m_fields
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithEventsLimit(m_events_limit int) MessagesGetLongPollHistoryRequest{
	m["events_limit"] = m_events_limit
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithMsgsLimit(m_msgs_limit int) MessagesGetLongPollHistoryRequest{
	m["msgs_limit"] = m_msgs_limit
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithMaxMsgId(m_max_msg_id int) MessagesGetLongPollHistoryRequest{
	m["max_msg_id"] = m_max_msg_id
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithGroupId(m_group_id int) MessagesGetLongPollHistoryRequest{
	m["group_id"] = m_group_id
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithLpVersion(m_lp_version int) MessagesGetLongPollHistoryRequest{
	m["lp_version"] = m_lp_version
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithLastN(m_last_n int) MessagesGetLongPollHistoryRequest{
	m["last_n"] = m_last_n
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithCredentials(m_credentials bool) MessagesGetLongPollHistoryRequest{
	m["credentials"] = m_credentials
	return m
}

func (m MessagesGetLongPollHistoryRequest) WithExtended(m_extended bool) MessagesGetLongPollHistoryRequest{
	m["extended"] = m_extended
	return m
}

func (m MessagesGetLongPollHistoryRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesTooOldPts, Error_MessagesTooNewPts, Error_Timeout ]
//
// https://dev.vk.com/method/messages.getLongPollHistory
func (m *Messages) MessagesGetLongPollHistory(params ...api.MethodParams) (resp models.MessagesGetLongPollHistoryResponse, err error) {
	req := api.NewRequest[models.MessagesGetLongPollHistoryResponse](m.api)

	res, err := req.Execute("messages.getLongPollHistory", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
