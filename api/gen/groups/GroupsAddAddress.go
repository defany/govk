// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsAddAddress ...
type GroupsAddAddressRequest api.Params

func NewGroupsAddAddressRequest() GroupsAddAddressRequest {
	params := make(GroupsAddAddressRequest, 14)
	return params
}

func (g GroupsAddAddressRequest) WithGroupId(g_group_id int) GroupsAddAddressRequest{
	g["group_id"] = g_group_id
	return g
}

func (g GroupsAddAddressRequest) WithTitle(g_title string) GroupsAddAddressRequest{
	g["title"] = g_title
	return g
}

func (g GroupsAddAddressRequest) WithAddress(g_address string) GroupsAddAddressRequest{
	g["address"] = g_address
	return g
}

func (g GroupsAddAddressRequest) WithAdditionalAddress(g_additional_address string) GroupsAddAddressRequest{
	g["additional_address"] = g_additional_address
	return g
}

func (g GroupsAddAddressRequest) WithCountryId(g_country_id int) GroupsAddAddressRequest{
	g["country_id"] = g_country_id
	return g
}

func (g GroupsAddAddressRequest) WithCityId(g_city_id int) GroupsAddAddressRequest{
	g["city_id"] = g_city_id
	return g
}

func (g GroupsAddAddressRequest) WithMetroId(g_metro_id int) GroupsAddAddressRequest{
	g["metro_id"] = g_metro_id
	return g
}

func (g GroupsAddAddressRequest) WithLatitude(g_latitude float64) GroupsAddAddressRequest{
	g["latitude"] = g_latitude
	return g
}

func (g GroupsAddAddressRequest) WithLongitude(g_longitude float64) GroupsAddAddressRequest{
	g["longitude"] = g_longitude
	return g
}

func (g GroupsAddAddressRequest) WithPhone(g_phone string) GroupsAddAddressRequest{
	g["phone"] = g_phone
	return g
}

func (g GroupsAddAddressRequest) WithWorkInfoStatus(g_work_info_status models.GroupsAddressWorkInfoStatus) GroupsAddAddressRequest{
	g["work_info_status"] = g_work_info_status
	return g
}

func (g GroupsAddAddressRequest) WithTimetable(g_timetable string) GroupsAddAddressRequest{
	g["timetable"] = g_timetable
	return g
}

func (g GroupsAddAddressRequest) WithIsMainAddress(g_is_main_address bool) GroupsAddAddressRequest{
	g["is_main_address"] = g_is_main_address
	return g
}

func (g GroupsAddAddressRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AccessGroups, Error_NotFound, Error_GroupTooManyAddresses ]
//
// https://dev.vk.com/method/groups.addAddress
func (g *Groups) GroupsAddAddress(params ...api.MethodParams) (resp models.GroupsAddAddressResponse, err error) {
	req := api.NewRequest[models.GroupsAddAddressResponse](g.api)

	res, err := req.Execute("groups.addAddress", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
