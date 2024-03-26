// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsAddOfficeUsers Adds managers and/or supervisors to advertising account.
type AdsAddOfficeUsersRequest api.Params

func NewAdsAddOfficeUsersRequest() AdsAddOfficeUsersRequest {
	params := make(AdsAddOfficeUsersRequest, 3)
	return params
}

func (a AdsAddOfficeUsersRequest) WithAccountId(a_account_id int) AdsAddOfficeUsersRequest{
	a["account_id"] = a_account_id
	return a
}

func (a AdsAddOfficeUsersRequest) WithData(a_data []models.AdsUserSpecificationCutted) AdsAddOfficeUsersRequest{
	a["data"] = a_data
	return a
}

func (a AdsAddOfficeUsersRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WeightedFlood ]
//
// https://dev.vk.com/method/ads.addOfficeUsers
func (a *Ads) AdsAddOfficeUsers(params ...api.MethodParams) (resp models.AdsAddOfficeUsersResponse, err error) {
	req := api.NewRequest[models.AdsAddOfficeUsersResponse](a.api)

	res, err := req.Execute("ads.addOfficeUsers", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
