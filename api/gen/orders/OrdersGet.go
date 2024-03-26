// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// OrdersGet Returns a list of orders.
type OrdersGetRequest api.Params

func NewOrdersGetRequest() OrdersGetRequest {
	params := make(OrdersGetRequest, 4)
	return params
}

func (o OrdersGetRequest) WithOffset(o_offset int) OrdersGetRequest{
	o["offset"] = o_offset
	return o
}

func (o OrdersGetRequest) WithCount(o_count int) OrdersGetRequest{
	o["count"] = o_count
	return o
}

func (o OrdersGetRequest) WithTestMode(o_test_mode bool) OrdersGetRequest{
	o["test_mode"] = o_test_mode
	return o
}

func (o OrdersGetRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.get
func (o *Orders) OrdersGet(params ...api.MethodParams) (resp models.OrdersGetResponse, err error) {
	req := api.NewRequest[models.OrdersGetResponse](o.api)

	res, err := req.Execute("orders.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

