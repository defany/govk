// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// MarketGetCategories Returns a list of market categories.
type MarketGetCategoriesRequest api.Params

func NewMarketGetCategoriesRequest() MarketGetCategoriesRequest {
	params := make(MarketGetCategoriesRequest, 3)
	return params
}

func (m MarketGetCategoriesRequest) WithCount(m_count int) MarketGetCategoriesRequest{
	m["count"] = m_count
	return m
}

func (m MarketGetCategoriesRequest) WithOffset(m_offset int) MarketGetCategoriesRequest{
	m["offset"] = m_offset
	return m
}

func (m MarketGetCategoriesRequest) Params() api.Params {
	return api.Params(m)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/market.getCategories
func (m *Market) MarketGetCategories(params ...api.MethodParams) (resp models.MarketGetCategoriesResponse, err error) {
	req := api.NewRequest[models.MarketGetCategoriesResponse](m.api)

	res, err := req.Execute("market.getCategories", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

