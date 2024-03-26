// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// OrdersGetById Returns information about orders by their IDs.
type OrdersGetByIdRequest api.Params

func NewOrdersGetByIdRequest() OrdersGetByIdRequest {
	params := make(OrdersGetByIdRequest, 4)
	return params
}

func (o OrdersGetByIdRequest) WithOrderId(o_order_id int) OrdersGetByIdRequest{
	o["order_id"] = o_order_id
	return o
}

func (o OrdersGetByIdRequest) WithOrderIds(o_order_ids []int) OrdersGetByIdRequest{
	o["order_ids"] = o_order_ids
	return o
}

func (o OrdersGetByIdRequest) WithTestMode(o_test_mode bool) OrdersGetByIdRequest{
	o["test_mode"] = o_test_mode
	return o
}

func (o OrdersGetByIdRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/orders.getById
func (o *Orders) OrdersGetById(params ...api.MethodParams) (resp models.OrdersGetByIdResponse, err error) {
	req := api.NewRequest[models.OrdersGetByIdResponse](o.api)

	res, err := req.Execute("orders.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
