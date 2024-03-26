// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StatsGet Returns statistics of a community or an application.
type StatsGetRequest api.Params

func NewStatsGetRequest() StatsGetRequest {
	params := make(StatsGetRequest, 10)
	return params
}

func (s StatsGetRequest) WithGroupId(s_group_id int) StatsGetRequest{
	s["group_id"] = s_group_id
	return s
}

func (s StatsGetRequest) WithAppId(s_app_id int) StatsGetRequest{
	s["app_id"] = s_app_id
	return s
}

func (s StatsGetRequest) WithTimestampFrom(s_timestamp_from int) StatsGetRequest{
	s["timestamp_from"] = s_timestamp_from
	return s
}

func (s StatsGetRequest) WithTimestampTo(s_timestamp_to int) StatsGetRequest{
	s["timestamp_to"] = s_timestamp_to
	return s
}

func (s StatsGetRequest) WithInterval(s_interval string) StatsGetRequest{
	s["interval"] = s_interval
	return s
}

func (s StatsGetRequest) WithIntervalsCount(s_intervals_count int) StatsGetRequest{
	s["intervals_count"] = s_intervals_count
	return s
}

func (s StatsGetRequest) WithFilters(s_filters []string) StatsGetRequest{
	s["filters"] = s_filters
	return s
}

func (s StatsGetRequest) WithStatsGroups(s_stats_groups []string) StatsGetRequest{
	s["stats_groups"] = s_stats_groups
	return s
}

func (s StatsGetRequest) WithExtended(s_extended bool) StatsGetRequest{
	s["extended"] = s_extended
	return s
}

func (s StatsGetRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/stats.get
func (s *Stats) StatsGet(params ...api.MethodParams) (resp models.StatsGetResponse, err error) {
	req := api.NewRequest[models.StatsGetResponse](s.api)

	res, err := req.Execute("stats.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
