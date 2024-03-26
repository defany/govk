// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StoreGetProducts ...
type StoreGetProductsRequest api.Params

func NewStoreGetProductsRequest() StoreGetProductsRequest {
	params := make(StoreGetProductsRequest, 7)
	return params
}

func (s StoreGetProductsRequest) WithType(s_type string) StoreGetProductsRequest{
	s["type"] = s_type
	return s
}

func (s StoreGetProductsRequest) WithMerchant(s_merchant string) StoreGetProductsRequest{
	s["merchant"] = s_merchant
	return s
}

func (s StoreGetProductsRequest) WithSection(s_section string) StoreGetProductsRequest{
	s["section"] = s_section
	return s
}

func (s StoreGetProductsRequest) WithProductIds(s_product_ids []int) StoreGetProductsRequest{
	s["product_ids"] = s_product_ids
	return s
}

func (s StoreGetProductsRequest) WithFilters(s_filters []string) StoreGetProductsRequest{
	s["filters"] = s_filters
	return s
}

func (s StoreGetProductsRequest) WithExtended(s_extended bool) StoreGetProductsRequest{
	s["extended"] = s_extended
	return s
}

func (s StoreGetProductsRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/store.getProducts
func (s *Store) StoreGetProducts(params ...api.MethodParams) (resp models.StoreGetProductsResponse, err error) {
	req := api.NewRequest[models.StoreGetProductsResponse](s.api)

	res, err := req.Execute("store.getProducts", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
