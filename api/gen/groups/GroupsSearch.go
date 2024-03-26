// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsSearch Returns a list of communities matching the search criteria.
type GroupsSearchRequest api.Params

func NewGroupsSearchRequest() GroupsSearchRequest {
	params := make(GroupsSearchRequest, 10)
	return params
}

func (g GroupsSearchRequest) WithQ(g_q string) GroupsSearchRequest{
	g["q"] = g_q
	return g
}

func (g GroupsSearchRequest) WithType(g_type string) GroupsSearchRequest{
	g["type"] = g_type
	return g
}

func (g GroupsSearchRequest) WithCountryId(g_country_id int) GroupsSearchRequest{
	g["country_id"] = g_country_id
	return g
}

func (g GroupsSearchRequest) WithCityId(g_city_id int) GroupsSearchRequest{
	g["city_id"] = g_city_id
	return g
}

func (g GroupsSearchRequest) WithFuture(g_future bool) GroupsSearchRequest{
	g["future"] = g_future
	return g
}

func (g GroupsSearchRequest) WithMarket(g_market bool) GroupsSearchRequest{
	g["market"] = g_market
	return g
}

func (g GroupsSearchRequest) WithSort(g_sort int) GroupsSearchRequest{
	g["sort"] = g_sort
	return g
}

func (g GroupsSearchRequest) WithOffset(g_offset int) GroupsSearchRequest{
	g["offset"] = g_offset
	return g
}

func (g GroupsSearchRequest) WithCount(g_count int) GroupsSearchRequest{
	g["count"] = g_count
	return g
}

func (g GroupsSearchRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/groups.search
func (g *Groups) GroupsSearch(params ...api.MethodParams) (resp models.GroupsSearchResponse, err error) {
	req := api.NewRequest[models.GroupsSearchResponse](g.api)

	res, err := req.Execute("groups.search", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
