// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsGetRejectionReason Returns a reason of ad rejection for pre-moderation.
type AdsGetRejectionReasonRequest api.Params

func NewAdsGetRejectionReasonRequest() AdsGetRejectionReasonRequest {
	params := make(AdsGetRejectionReasonRequest, 3)
	return params
}

func (a AdsGetRejectionReasonRequest) WithAccountId(a_account_id int) AdsGetRejectionReasonRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsGetRejectionReasonRequest) WithAdId(a_ad_id int) AdsGetRejectionReasonRequest{
	a["ad_id"] = a_ad_id
	return a
}

func (a AdsGetRejectionReasonRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.getRejectionReason
func (a *Ads) AdsGetRejectionReason(params ...api.MethodParams) (resp models.AdsGetRejectionReasonResponse, err error) {
	req := api.NewRequest[models.AdsGetRejectionReasonResponse](a.api)

	res, err := req.Execute("ads.getRejectionReason", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

