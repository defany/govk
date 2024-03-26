// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsCreateCampaigns Creates advertising campaigns.
type AdsCreateCampaignsRequest api.Params

func NewAdsCreateCampaignsRequest() AdsCreateCampaignsRequest {
	params := make(AdsCreateCampaignsRequest, 3)
	return params
}

func (a AdsCreateCampaignsRequest) WithAccountId(a_account_id int) AdsCreateCampaignsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsCreateCampaignsRequest) WithData(a_data string) AdsCreateCampaignsRequest{
	a["data"] = a_data
	return a
}

func (a AdsCreateCampaignsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AdsPartialSuccess, Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.createCampaigns
func (a *Ads) AdsCreateCampaigns(params ...api.MethodParams) (resp models.AdsCreateCampaignsResponse, err error) {
	req := api.NewRequest[models.AdsCreateCampaignsResponse](a.api)

	res, err := req.Execute("ads.createCampaigns", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

