// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsDeleteCampaigns Archives advertising campaigns.
type AdsDeleteCampaignsRequest api.Params

func NewAdsDeleteCampaignsRequest() AdsDeleteCampaignsRequest {
	params := make(AdsDeleteCampaignsRequest, 3)
	return params
}

func (a AdsDeleteCampaignsRequest) WithAccountId(a_account_id int) AdsDeleteCampaignsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsDeleteCampaignsRequest) WithIds(a_ids string) AdsDeleteCampaignsRequest{
	a["ids"] = a_ids
	return a
}

func (a AdsDeleteCampaignsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AdsObjectDeleted, Error_AdsPartialSuccess, Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.deleteCampaigns
func (a *Ads) AdsDeleteCampaigns(params ...api.MethodParams) (resp models.AdsDeleteCampaignsResponse, err error) {
	req := api.NewRequest[models.AdsDeleteCampaignsResponse](a.api)

	res, err := req.Execute("ads.deleteCampaigns", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

