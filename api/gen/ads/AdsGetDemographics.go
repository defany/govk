// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsGetDemographics Returns demographics for ads or campaigns.
type AdsGetDemographicsRequest api.Params

func NewAdsGetDemographicsRequest() AdsGetDemographicsRequest {
	params := make(AdsGetDemographicsRequest, 7)
	return params
}

func (a AdsGetDemographicsRequest) WithAccountId(a_account_id int) AdsGetDemographicsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsGetDemographicsRequest) WithIdsType(a_ids_type string) AdsGetDemographicsRequest{
	a["ids_type"] = a_ids_type
	return a
}

func (a AdsGetDemographicsRequest) WithIds(a_ids string) AdsGetDemographicsRequest{
	a["ids"] = a_ids
	return a
}

func (a AdsGetDemographicsRequest) WithPeriod(a_period string) AdsGetDemographicsRequest{
	a["period"] = a_period
	return a
}

func (a AdsGetDemographicsRequest) WithDateFrom(a_date_from string) AdsGetDemographicsRequest{
	a["date_from"] = a_date_from
	return a
}

func (a AdsGetDemographicsRequest) WithDateTo(a_date_to string) AdsGetDemographicsRequest{
	a["date_to"] = a_date_to
	return a
}

func (a AdsGetDemographicsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.getDemographics
func (a *Ads) AdsGetDemographics(params ...api.MethodParams) (resp models.AdsGetDemographicsResponse, err error) {
	req := api.NewRequest[models.AdsGetDemographicsResponse](a.api)

	res, err := req.Execute("ads.getDemographics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
