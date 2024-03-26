// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StatsGetPostReach Returns stats for a wall post.
type StatsGetPostReachRequest api.Params

func NewStatsGetPostReachRequest() StatsGetPostReachRequest {
	params := make(StatsGetPostReachRequest, 3)
	return params
}

func (s StatsGetPostReachRequest) WithOwnerId(s_owner_id string) StatsGetPostReachRequest{
	s["owner_id"] = s_owner_id
	return s
}

func (s StatsGetPostReachRequest) WithPostIds(s_post_ids []int) StatsGetPostReachRequest{
	s["post_ids"] = s_post_ids
	return s
}

func (s StatsGetPostReachRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallAccessPost ]
//
// https://dev.vk.com/method/stats.getPostReach
func (s *Stats) StatsGetPostReach(params ...api.MethodParams) (resp models.StatsGetPostReachResponse, err error) {
	req := api.NewRequest[models.StatsGetPostReachResponse](s.api)

	res, err := req.Execute("stats.getPostReach", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

