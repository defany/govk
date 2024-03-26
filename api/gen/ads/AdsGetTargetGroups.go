// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsGetTargetGroups Returns a list of target groups.
type AdsGetTargetGroupsRequest api.Params

func NewAdsGetTargetGroupsRequest() AdsGetTargetGroupsRequest {
	params := make(AdsGetTargetGroupsRequest, 4)
	return params
}

func (a AdsGetTargetGroupsRequest) WithAccountId(a_account_id int) AdsGetTargetGroupsRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsGetTargetGroupsRequest) WithClientId(a_client_id int) AdsGetTargetGroupsRequest{
	a["client_id"] = a_client_id
	return a
}

func (a AdsGetTargetGroupsRequest) WithExtended(a_extended bool) AdsGetTargetGroupsRequest{
	a["extended"] = a_extended
	return a
}

func (a AdsGetTargetGroupsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.getTargetGroups
func (a *Ads) AdsGetTargetGroups(params ...api.MethodParams) (resp models.AdsGetTargetGroupsResponse, err error) {
	req := api.NewRequest[models.AdsGetTargetGroupsResponse](a.api)

	res, err := req.Execute("ads.getTargetGroups", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

