// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsUpdateAds Edits ads.
type AdsUpdateAdsRequest api.Params

func NewAdsUpdateAdsRequest() AdsUpdateAdsRequest {
	params := make(AdsUpdateAdsRequest, 3)
	return params
}

func (a AdsUpdateAdsRequest) WithAccountId(a_account_id int) AdsUpdateAdsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsUpdateAdsRequest) WithData(a_data string) AdsUpdateAdsRequest{
	a["data"] = a_data
	return a
}

func (a AdsUpdateAdsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.updateAds
func (a *Ads) AdsUpdateAds(params ...api.MethodParams) (resp models.AdsUpdateAdsResponse, err error) {
	req := api.NewRequest[models.AdsUpdateAdsResponse](a.api)

	res, err := req.Execute("ads.updateAds", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

