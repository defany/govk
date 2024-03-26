// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsUpdateClients Edits clients of an advertising agency.
type AdsUpdateClientsRequest api.Params

func NewAdsUpdateClientsRequest() AdsUpdateClientsRequest {
	params := make(AdsUpdateClientsRequest, 3)
	return params
}

func (a AdsUpdateClientsRequest) WithAccountId(a_account_id int) AdsUpdateClientsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsUpdateClientsRequest) WithData(a_data string) AdsUpdateClientsRequest{
	a["data"] = a_data
	return a
}

func (a AdsUpdateClientsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.updateClients
func (a *Ads) AdsUpdateClients(params ...api.MethodParams) (resp models.AdsUpdateClientsResponse, err error) {
	req := api.NewRequest[models.AdsUpdateClientsResponse](a.api)

	res, err := req.Execute("ads.updateClients", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
