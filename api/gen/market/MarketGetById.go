// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketGetById Returns information about market items by their ids.
type MarketGetByIdRequest api.Params

func NewMarketGetByIdRequest() MarketGetByIdRequest {
	params := make(MarketGetByIdRequest, 4)
	return params
}

func (m MarketGetByIdRequest) WithItemIds(m_item_ids []string) MarketGetByIdRequest{
	m["item_ids"] = m_item_ids
	return m
}

func (m MarketGetByIdRequest) WithExtended(m_extended bool) MarketGetByIdRequest{
	m["extended"] = m_extended
	return m
}

func (m MarketGetByIdRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.getById
func (m *Market) MarketGetById(params ...api.MethodParams) (resp models.MarketGetByIdResponse, err error) {
	req := api.NewRequest[models.MarketGetByIdResponse](m.api)

	res, err := req.Execute("market.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// MarketGetByIdExtended Returns information about market items by their ids.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.getById
func (m *Market) MarketGetByIdExtended(params ...api.MethodParams) (resp models.MarketGetByIdExtendedResponse, err error) {
	req := api.NewRequest[models.MarketGetByIdExtendedResponse](m.api)

	res, err := req.Execute("market.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
