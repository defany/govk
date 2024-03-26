// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// OrdersUpdateSubscription ...
type OrdersUpdateSubscriptionRequest api.Params

func NewOrdersUpdateSubscriptionRequest() OrdersUpdateSubscriptionRequest {
	params := make(OrdersUpdateSubscriptionRequest, 4)
	return params
}

func (o OrdersUpdateSubscriptionRequest) WithUserId(o_user_id int) OrdersUpdateSubscriptionRequest{
	o["user_id"] = o_user_id
	return o
}

func (o OrdersUpdateSubscriptionRequest) WithSubscriptionId(o_subscription_id int) OrdersUpdateSubscriptionRequest{
	o["subscription_id"] = o_subscription_id
	return o
}

func (o OrdersUpdateSubscriptionRequest) WithPrice(o_price int) OrdersUpdateSubscriptionRequest{
	o["price"] = o_price
	return o
}

func (o OrdersUpdateSubscriptionRequest) Params() api.Params {
	return api.Params(o)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AppsSubscriptionNotFound, Error_AppsSubscriptionInvalidStatus ]
//
// https://dev.vk.com/method/orders.updateSubscription
func (o *Orders) OrdersUpdateSubscription(params ...api.MethodParams) (resp models.OrdersUpdateSubscriptionResponse, err error) {
	req := api.NewRequest[models.OrdersUpdateSubscriptionResponse](o.api)

	res, err := req.Execute("orders.updateSubscription", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

