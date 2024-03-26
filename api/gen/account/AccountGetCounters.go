// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AccountGetCounters Returns non-null values of user counters.
type AccountGetCountersRequest api.Params

func NewAccountGetCountersRequest() AccountGetCountersRequest {
	params := make(AccountGetCountersRequest, 3)
	return params
}

func (a AccountGetCountersRequest) WithFilter(a_filter []string) AccountGetCountersRequest{
	a["filter"] = a_filter
	return a
}

func (a AccountGetCountersRequest) WithUserId(a_user_id int) AccountGetCountersRequest{
	a["user_id"] = a_user_id
	return a
}

func (a AccountGetCountersRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/account.getCounters
func (a *Account) AccountGetCounters(params ...api.MethodParams) (resp models.AccountGetCountersResponse, err error) {
	req := api.NewRequest[models.AccountGetCountersResponse](a.api)

	res, err := req.Execute("account.getCounters", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
