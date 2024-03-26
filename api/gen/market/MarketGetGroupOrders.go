// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketGetGroupOrders Get market orders
type MarketGetGroupOrdersRequest api.Params

func NewMarketGetGroupOrdersRequest() MarketGetGroupOrdersRequest {
	params := make(MarketGetGroupOrdersRequest, 4)
	return params
}

func (m MarketGetGroupOrdersRequest) WithGroupId(m_group_id int) MarketGetGroupOrdersRequest{
	m["group_id"] = m_group_id
	return m
}

func (m MarketGetGroupOrdersRequest) WithOffset(m_offset int) MarketGetGroupOrdersRequest{
	m["offset"] = m_offset
	return m
}

func (m MarketGetGroupOrdersRequest) WithCount(m_count int) MarketGetGroupOrdersRequest{
	m["count"] = m_count
	return m
}

func (m MarketGetGroupOrdersRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MarketExtendedNotEnabled ]
//
// https://dev.vk.com/method/market.getGroupOrders
func (m *Market) MarketGetGroupOrders(params ...api.MethodParams) (resp models.MarketGetGroupOrdersResponse, err error) {
	req := api.NewRequest[models.MarketGetGroupOrdersResponse](m.api)

	res, err := req.Execute("market.getGroupOrders", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

