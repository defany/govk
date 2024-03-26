// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsDeleteAds Archives ads.
type AdsDeleteAdsRequest api.Params

func NewAdsDeleteAdsRequest() AdsDeleteAdsRequest {
	params := make(AdsDeleteAdsRequest, 3)
	return params
}

func (a AdsDeleteAdsRequest) WithAccountId(a_account_id int) AdsDeleteAdsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsDeleteAdsRequest) WithIds(a_ids string) AdsDeleteAdsRequest{
	a["ids"] = a_ids
	return a
}

func (a AdsDeleteAdsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AdsObjectDeleted, Error_AdsPartialSuccess, Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.deleteAds
func (a *Ads) AdsDeleteAds(params ...api.MethodParams) (resp models.AdsDeleteAdsResponse, err error) {
	req := api.NewRequest[models.AdsDeleteAdsResponse](a.api)

	res, err := req.Execute("ads.deleteAds", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
