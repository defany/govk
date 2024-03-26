// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// NewsfeedSearch Returns search results by statuses.
type NewsfeedSearchRequest api.Params

func NewNewsfeedSearchRequest() NewsfeedSearchRequest {
	params := make(NewsfeedSearchRequest, 11)
	return params
}

func (n NewsfeedSearchRequest) WithQ(n_q string) NewsfeedSearchRequest{
	n["q"] = n_q
	return n
}

func (n NewsfeedSearchRequest) WithExtended(n_extended bool) NewsfeedSearchRequest{
	n["extended"] = n_extended
	return n
}

func (n NewsfeedSearchRequest) WithCount(n_count int) NewsfeedSearchRequest{
	n["count"] = n_count
	return n
}

func (n NewsfeedSearchRequest) WithLatitude(n_latitude float64) NewsfeedSearchRequest{
	n["latitude"] = n_latitude
	return n
}

func (n NewsfeedSearchRequest) WithLongitude(n_longitude float64) NewsfeedSearchRequest{
	n["longitude"] = n_longitude
	return n
}

func (n NewsfeedSearchRequest) WithStartTime(n_start_time int) NewsfeedSearchRequest{
	n["start_time"] = n_start_time
	return n
}

func (n NewsfeedSearchRequest) WithEndTime(n_end_time int) NewsfeedSearchRequest{
	n["end_time"] = n_end_time
	return n
}

func (n NewsfeedSearchRequest) WithStartFrom(n_start_from string) NewsfeedSearchRequest{
	n["start_from"] = n_start_from
	return n
}

func (n NewsfeedSearchRequest) WithFields(n_fields []models.BaseUserGroupFields) NewsfeedSearchRequest{
	n["fields"] = n_fields
	return n
}

func (n NewsfeedSearchRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/newsfeed.search
func (n *Newsfeed) NewsfeedSearch(params ...api.MethodParams) (resp models.NewsfeedSearchResponse, err error) {
	req := api.NewRequest[models.NewsfeedSearchResponse](n.api)

	res, err := req.Execute("newsfeed.search", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// NewsfeedSearchExtended Returns search results by statuses.
// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/newsfeed.search
func (n *Newsfeed) NewsfeedSearchExtended(params ...api.MethodParams) (resp models.NewsfeedSearchExtendedResponse, err error) {
	req := api.NewRequest[models.NewsfeedSearchExtendedResponse](n.api)

	res, err := req.Execute("newsfeed.search", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

