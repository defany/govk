// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StoriesSearch ...
type StoriesSearchRequest api.Params

func NewStoriesSearchRequest() StoriesSearchRequest {
	params := make(StoriesSearchRequest, 10)
	return params
}

func (s StoriesSearchRequest) WithQ(s_q string) StoriesSearchRequest{
	s["q"] = s_q
	return s
}

func (s StoriesSearchRequest) WithPlaceId(s_place_id int) StoriesSearchRequest{
	s["place_id"] = s_place_id
	return s
}

func (s StoriesSearchRequest) WithLatitude(s_latitude float64) StoriesSearchRequest{
	s["latitude"] = s_latitude
	return s
}

func (s StoriesSearchRequest) WithLongitude(s_longitude float64) StoriesSearchRequest{
	s["longitude"] = s_longitude
	return s
}

func (s StoriesSearchRequest) WithRadius(s_radius int) StoriesSearchRequest{
	s["radius"] = s_radius
	return s
}

func (s StoriesSearchRequest) WithMentionedId(s_mentioned_id int) StoriesSearchRequest{
	s["mentioned_id"] = s_mentioned_id
	return s
}

func (s StoriesSearchRequest) WithCount(s_count int) StoriesSearchRequest{
	s["count"] = s_count
	return s
}

func (s StoriesSearchRequest) WithExtended(s_extended bool) StoriesSearchRequest{
	s["extended"] = s_extended
	return s
}

func (s StoriesSearchRequest) WithFields(s_fields []models.BaseUserGroupFields) StoriesSearchRequest{
	s["fields"] = s_fields
	return s
}

func (s StoriesSearchRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/stories.search
func (s *Stories) StoriesSearch(params ...api.MethodParams) (resp models.StoriesGetV5113Response, err error) {
	req := api.NewRequest[models.StoriesGetV5113Response](s.api)

	res, err := req.Execute("stories.search", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

