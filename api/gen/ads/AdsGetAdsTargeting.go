// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsGetAdsTargeting Returns ad targeting parameters.
type AdsGetAdsTargetingRequest api.Params

func NewAdsGetAdsTargetingRequest() AdsGetAdsTargetingRequest {
	params := make(AdsGetAdsTargetingRequest, 8)
	return params
}

func (a AdsGetAdsTargetingRequest) WithAccountId(a_account_id int) AdsGetAdsTargetingRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsGetAdsTargetingRequest) WithAdIds(a_ad_ids string) AdsGetAdsTargetingRequest{
	a["ad_ids"] = a_ad_ids
	return a
}

func (a AdsGetAdsTargetingRequest) WithCampaignIds(a_campaign_ids string) AdsGetAdsTargetingRequest{
	a["campaign_ids"] = a_campaign_ids
	return a
}

func (a AdsGetAdsTargetingRequest) WithClientId(a_client_id int) AdsGetAdsTargetingRequest{
	a["client_id"] = a_client_id
	return a
}

func (a AdsGetAdsTargetingRequest) WithIncludeDeleted(a_include_deleted bool) AdsGetAdsTargetingRequest{
	a["include_deleted"] = a_include_deleted
	return a
}

func (a AdsGetAdsTargetingRequest) WithLimit(a_limit int) AdsGetAdsTargetingRequest{
	a["limit"] = a_limit
	return a
}

func (a AdsGetAdsTargetingRequest) WithOffset(a_offset int) AdsGetAdsTargetingRequest{
	a["offset"] = a_offset
	return a
}

func (a AdsGetAdsTargetingRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.getAdsTargeting
func (a *Ads) AdsGetAdsTargeting(params ...api.MethodParams) (resp models.AdsGetAdsTargetingResponse, err error) {
	req := api.NewRequest[models.AdsGetAdsTargetingResponse](a.api)

	res, err := req.Execute("ads.getAdsTargeting", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

