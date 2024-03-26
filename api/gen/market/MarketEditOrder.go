// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketEditOrder Edit order
type MarketEditOrderRequest api.Params

func NewMarketEditOrderRequest() MarketEditOrderRequest {
	params := make(MarketEditOrderRequest, 12)
	return params
}

func (m MarketEditOrderRequest) WithUserId(m_user_id int) MarketEditOrderRequest{
	m["user_id"] = m_user_id
	return m
}

func (m MarketEditOrderRequest) WithOrderId(m_order_id int) MarketEditOrderRequest{
	m["order_id"] = m_order_id
	return m
}

func (m MarketEditOrderRequest) WithMerchantComment(m_merchant_comment string) MarketEditOrderRequest{
	m["merchant_comment"] = m_merchant_comment
	return m
}

func (m MarketEditOrderRequest) WithStatus(m_status int) MarketEditOrderRequest{
	m["status"] = m_status
	return m
}

func (m MarketEditOrderRequest) WithTrackNumber(m_track_number string) MarketEditOrderRequest{
	m["track_number"] = m_track_number
	return m
}

func (m MarketEditOrderRequest) WithPaymentStatus(m_payment_status string) MarketEditOrderRequest{
	m["payment_status"] = m_payment_status
	return m
}

func (m MarketEditOrderRequest) WithDeliveryPrice(m_delivery_price int) MarketEditOrderRequest{
	m["delivery_price"] = m_delivery_price
	return m
}

func (m MarketEditOrderRequest) WithWidth(m_width int) MarketEditOrderRequest{
	m["width"] = m_width
	return m
}

func (m MarketEditOrderRequest) WithLength(m_length int) MarketEditOrderRequest{
	m["length"] = m_length
	return m
}

func (m MarketEditOrderRequest) WithHeight(m_height int) MarketEditOrderRequest{
	m["height"] = m_height
	return m
}

func (m MarketEditOrderRequest) WithWeight(m_weight int) MarketEditOrderRequest{
	m["weight"] = m_weight
	return m
}

func (m MarketEditOrderRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MarketOrdersNoCartItems, Error_MarketInvalidDimensions, Error_MarketCantChangeVkpayStatus ]
//
// https://dev.vk.com/method/market.editOrder
func (m *Market) MarketEditOrder(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](m.api)

	res, err := req.Execute("market.editOrder", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

