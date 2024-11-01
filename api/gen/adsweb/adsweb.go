// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Adsweb struct {
	api *api.API
}

func NewAdsweb(api *api.API) *Adsweb {
	return &Adsweb{
		api: api,
	}
}

// AdswebGetAdCategories ...
type AdswebGetAdCategoriesRequest api.Params

func NewAdswebGetAdCategoriesRequest() AdswebGetAdCategoriesRequest {
	params := make(AdswebGetAdCategoriesRequest, 2)
	return params
}

func (a AdswebGetAdCategoriesRequest) WithOfficeId(a_office_id int) AdswebGetAdCategoriesRequest {
	a["office_id"] = a_office_id
	return a
}

func (a AdswebGetAdCategoriesRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getAdCategories
func (a *Adsweb) AdswebGetAdCategories(params ...api.MethodParams) (resp models.AdswebGetAdCategoriesResponse, err error) {
	req := api.NewRequest[models.AdswebGetAdCategoriesResponse](a.api)

	res, err := req.Execute("adsweb.getAdCategories", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AdswebGetAdUnitCode ...
// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getAdUnitCode
func (a *Adsweb) AdswebGetAdUnitCode(params ...api.MethodParams) (resp models.AdswebGetAdUnitCodeResponse, err error) {
	req := api.NewRequest[models.AdswebGetAdUnitCodeResponse](a.api)

	res, err := req.Execute("adsweb.getAdUnitCode", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AdswebGetAdUnits ...
type AdswebGetAdUnitsRequest api.Params

func NewAdswebGetAdUnitsRequest() AdswebGetAdUnitsRequest {
	params := make(AdswebGetAdUnitsRequest, 7)
	return params
}

func (a AdswebGetAdUnitsRequest) WithOfficeId(a_office_id int) AdswebGetAdUnitsRequest {
	a["office_id"] = a_office_id
	return a
}

func (a AdswebGetAdUnitsRequest) WithSitesIds(a_sites_ids string) AdswebGetAdUnitsRequest {
	a["sites_ids"] = a_sites_ids
	return a
}

func (a AdswebGetAdUnitsRequest) WithAdUnitsIds(a_ad_units_ids string) AdswebGetAdUnitsRequest {
	a["ad_units_ids"] = a_ad_units_ids
	return a
}

func (a AdswebGetAdUnitsRequest) WithFields(a_fields string) AdswebGetAdUnitsRequest {
	a["fields"] = a_fields
	return a
}

func (a AdswebGetAdUnitsRequest) WithLimit(a_limit int) AdswebGetAdUnitsRequest {
	a["limit"] = a_limit
	return a
}

func (a AdswebGetAdUnitsRequest) WithOffset(a_offset int) AdswebGetAdUnitsRequest {
	a["offset"] = a_offset
	return a
}

func (a AdswebGetAdUnitsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getAdUnits
func (a *Adsweb) AdswebGetAdUnits(params ...api.MethodParams) (resp models.AdswebGetAdUnitsResponse, err error) {
	req := api.NewRequest[models.AdswebGetAdUnitsResponse](a.api)

	res, err := req.Execute("adsweb.getAdUnits", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AdswebGetFraudHistory ...
type AdswebGetFraudHistoryRequest api.Params

func NewAdswebGetFraudHistoryRequest() AdswebGetFraudHistoryRequest {
	params := make(AdswebGetFraudHistoryRequest, 5)
	return params
}

func (a AdswebGetFraudHistoryRequest) WithOfficeId(a_office_id int) AdswebGetFraudHistoryRequest {
	a["office_id"] = a_office_id
	return a
}

func (a AdswebGetFraudHistoryRequest) WithSitesIds(a_sites_ids string) AdswebGetFraudHistoryRequest {
	a["sites_ids"] = a_sites_ids
	return a
}

func (a AdswebGetFraudHistoryRequest) WithLimit(a_limit int) AdswebGetFraudHistoryRequest {
	a["limit"] = a_limit
	return a
}

func (a AdswebGetFraudHistoryRequest) WithOffset(a_offset int) AdswebGetFraudHistoryRequest {
	a["offset"] = a_offset
	return a
}

func (a AdswebGetFraudHistoryRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getFraudHistory
func (a *Adsweb) AdswebGetFraudHistory(params ...api.MethodParams) (resp models.AdswebGetFraudHistoryResponse, err error) {
	req := api.NewRequest[models.AdswebGetFraudHistoryResponse](a.api)

	res, err := req.Execute("adsweb.getFraudHistory", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AdswebGetSites ...
type AdswebGetSitesRequest api.Params

func NewAdswebGetSitesRequest() AdswebGetSitesRequest {
	params := make(AdswebGetSitesRequest, 6)
	return params
}

func (a AdswebGetSitesRequest) WithOfficeId(a_office_id int) AdswebGetSitesRequest {
	a["office_id"] = a_office_id
	return a
}

func (a AdswebGetSitesRequest) WithSitesIds(a_sites_ids string) AdswebGetSitesRequest {
	a["sites_ids"] = a_sites_ids
	return a
}

func (a AdswebGetSitesRequest) WithFields(a_fields string) AdswebGetSitesRequest {
	a["fields"] = a_fields
	return a
}

func (a AdswebGetSitesRequest) WithLimit(a_limit int) AdswebGetSitesRequest {
	a["limit"] = a_limit
	return a
}

func (a AdswebGetSitesRequest) WithOffset(a_offset int) AdswebGetSitesRequest {
	a["offset"] = a_offset
	return a
}

func (a AdswebGetSitesRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getSites
func (a *Adsweb) AdswebGetSites(params ...api.MethodParams) (resp models.AdswebGetSitesResponse, err error) {
	req := api.NewRequest[models.AdswebGetSitesResponse](a.api)

	res, err := req.Execute("adsweb.getSites", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// AdswebGetStatistics ...
type AdswebGetStatisticsRequest api.Params

func NewAdswebGetStatisticsRequest() AdswebGetStatisticsRequest {
	params := make(AdswebGetStatisticsRequest, 10)
	return params
}

func (a AdswebGetStatisticsRequest) WithOfficeId(a_office_id int) AdswebGetStatisticsRequest {
	a["office_id"] = a_office_id
	return a
}

func (a AdswebGetStatisticsRequest) WithIdsType(a_ids_type string) AdswebGetStatisticsRequest {
	a["ids_type"] = a_ids_type
	return a
}

func (a AdswebGetStatisticsRequest) WithIds(a_ids string) AdswebGetStatisticsRequest {
	a["ids"] = a_ids
	return a
}

func (a AdswebGetStatisticsRequest) WithPeriod(a_period string) AdswebGetStatisticsRequest {
	a["period"] = a_period
	return a
}

func (a AdswebGetStatisticsRequest) WithDateFrom(a_date_from string) AdswebGetStatisticsRequest {
	a["date_from"] = a_date_from
	return a
}

func (a AdswebGetStatisticsRequest) WithDateTo(a_date_to string) AdswebGetStatisticsRequest {
	a["date_to"] = a_date_to
	return a
}

func (a AdswebGetStatisticsRequest) WithFields(a_fields string) AdswebGetStatisticsRequest {
	a["fields"] = a_fields
	return a
}

func (a AdswebGetStatisticsRequest) WithLimit(a_limit int) AdswebGetStatisticsRequest {
	a["limit"] = a_limit
	return a
}

func (a AdswebGetStatisticsRequest) WithPageId(a_page_id string) AdswebGetStatisticsRequest {
	a["page_id"] = a_page_id
	return a
}

func (a AdswebGetStatisticsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/adsweb.getStatistics
func (a *Adsweb) AdswebGetStatistics(params ...api.MethodParams) (resp models.AdswebGetStatisticsResponse, err error) {
	req := api.NewRequest[models.AdswebGetStatisticsResponse](a.api)

	res, err := req.Execute("adsweb.getStatistics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
