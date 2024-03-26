// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketGet Returns items list for a community.
type MarketGetRequest api.Params

func NewMarketGetRequest() MarketGetRequest {
	params := make(MarketGetRequest, 11)
	return params
}

func (m MarketGetRequest) WithOwnerId(m_owner_id int) MarketGetRequest{
	m["owner_id"] = m_owner_id
	return m
}

func (m MarketGetRequest) WithAlbumId(m_album_id int) MarketGetRequest{
	m["album_id"] = m_album_id
	return m
}

func (m MarketGetRequest) WithCount(m_count int) MarketGetRequest{
	m["count"] = m_count
	return m
}

func (m MarketGetRequest) WithOffset(m_offset int) MarketGetRequest{
	m["offset"] = m_offset
	return m
}

func (m MarketGetRequest) WithExtended(m_extended bool) MarketGetRequest{
	m["extended"] = m_extended
	return m
}

func (m MarketGetRequest) WithDateFrom(m_date_from string) MarketGetRequest{
	m["date_from"] = m_date_from
	return m
}

func (m MarketGetRequest) WithDateTo(m_date_to string) MarketGetRequest{
	m["date_to"] = m_date_to
	return m
}

func (m MarketGetRequest) WithNeedVariants(m_need_variants bool) MarketGetRequest{
	m["need_variants"] = m_need_variants
	return m
}

func (m MarketGetRequest) WithWithDisabled(m_with_disabled bool) MarketGetRequest{
	m["with_disabled"] = m_with_disabled
	return m
}

func (m MarketGetRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.get
func (m *Market) MarketGet(params ...api.MethodParams) (resp models.MarketGetResponse, err error) {
	req := api.NewRequest[models.MarketGetResponse](m.api)

	res, err := req.Execute("market.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// MarketGetExtended Returns items list for a community.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.get
func (m *Market) MarketGetExtended(params ...api.MethodParams) (resp models.MarketGetExtendedResponse, err error) {
	req := api.NewRequest[models.MarketGetExtendedResponse](m.api)

	res, err := req.Execute("market.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
