// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// NewsfeedUnsubscribe Unsubscribes the current user from specified newsfeeds.
type NewsfeedUnsubscribeRequest api.Params

func NewNewsfeedUnsubscribeRequest() NewsfeedUnsubscribeRequest {
	params := make(NewsfeedUnsubscribeRequest, 4)
	return params
}

func (n NewsfeedUnsubscribeRequest) WithType(n_type string) NewsfeedUnsubscribeRequest{
	n["type"] = n_type
	return n
}

func (n NewsfeedUnsubscribeRequest) WithOwnerId(n_owner_id int) NewsfeedUnsubscribeRequest{
	n["owner_id"] = n_owner_id
	return n
}

func (n NewsfeedUnsubscribeRequest) WithItemId(n_item_id int) NewsfeedUnsubscribeRequest{
	n["item_id"] = n_item_id
	return n
}

func (n NewsfeedUnsubscribeRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/newsfeed.unsubscribe
func (n *Newsfeed) NewsfeedUnsubscribe(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](n.api)

	res, err := req.Execute("newsfeed.unsubscribe", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

