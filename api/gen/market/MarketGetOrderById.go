// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketGetOrderById Get order
type MarketGetOrderByIdRequest api.Params

func NewMarketGetOrderByIdRequest() MarketGetOrderByIdRequest {
	params := make(MarketGetOrderByIdRequest, 4)
	return params
}

func (m MarketGetOrderByIdRequest) WithUserId(m_user_id int) MarketGetOrderByIdRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MarketGetOrderByIdRequest) WithOrderId(m_order_id int) MarketGetOrderByIdRequest{
	m["order_id"] = m_order_id
	return m
}

func (m MarketGetOrderByIdRequest) WithExtended(m_extended bool) MarketGetOrderByIdRequest{
	m["extended"] = m_extended
	return m
}

func (m MarketGetOrderByIdRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.getOrderById
func (m *Market) MarketGetOrderById(params ...api.MethodParams) (resp models.MarketGetOrderByIdResponse, err error) {
	req := api.NewRequest[models.MarketGetOrderByIdResponse](m.api)

	res, err := req.Execute("market.getOrderById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
