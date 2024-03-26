// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// OrdersGetUserSubscriptionById ...
type OrdersGetUserSubscriptionByIdRequest api.Params

func NewOrdersGetUserSubscriptionByIdRequest() OrdersGetUserSubscriptionByIdRequest {
	params := make(OrdersGetUserSubscriptionByIdRequest, 3)
	return params
}

func (o OrdersGetUserSubscriptionByIdRequest) WithUserId(o_user_id int) OrdersGetUserSubscriptionByIdRequest{
	o["user_id"] = o_user_id
	return o
}

func (o OrdersGetUserSubscriptionByIdRequest) WithSubscriptionId(o_subscription_id int) OrdersGetUserSubscriptionByIdRequest{
	o["subscription_id"] = o_subscription_id
	return o
}

func (o OrdersGetUserSubscriptionByIdRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AppsSubscriptionNotFound ]
//
// https://dev.vk.com/method/orders.getUserSubscriptionById
func (o *Orders) OrdersGetUserSubscriptionById(params ...api.MethodParams) (resp models.OrdersGetUserSubscriptionByIdResponse, err error) {
	req := api.NewRequest[models.OrdersGetUserSubscriptionByIdResponse](o.api)

	res, err := req.Execute("orders.getUserSubscriptionById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
