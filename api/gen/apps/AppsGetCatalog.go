// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AppsGetCatalog Returns a list of applications (apps) available to users in the App Catalog.
type AppsGetCatalogRequest api.Params

func NewAppsGetCatalogRequest() AppsGetCatalogRequest {
	params := make(AppsGetCatalogRequest, 12)
	return params
}

func (a AppsGetCatalogRequest) WithSort(a_sort string) AppsGetCatalogRequest{
	a["sort"] = a_sort
	return a
}

func (a AppsGetCatalogRequest) WithOffset(a_offset int) AppsGetCatalogRequest{
	a["offset"] = a_offset
	return a
}

func (a AppsGetCatalogRequest) WithCount(a_count int) AppsGetCatalogRequest{
	a["count"] = a_count
	return a
}

func (a AppsGetCatalogRequest) WithPlatform(a_platform string) AppsGetCatalogRequest{
	a["platform"] = a_platform
	return a
}

func (a AppsGetCatalogRequest) WithExtended(a_extended bool) AppsGetCatalogRequest{
	a["extended"] = a_extended
	return a
}

func (a AppsGetCatalogRequest) WithReturnFriends(a_return_friends bool) AppsGetCatalogRequest{
	a["return_friends"] = a_return_friends
	return a
}

func (a AppsGetCatalogRequest) WithFields(a_fields []models.UsersFields) AppsGetCatalogRequest{
	a["fields"] = a_fields
	return a
}

func (a AppsGetCatalogRequest) WithNameCase(a_name_case string) AppsGetCatalogRequest{
	a["name_case"] = a_name_case
	return a
}

func (a AppsGetCatalogRequest) WithQ(a_q string) AppsGetCatalogRequest{
	a["q"] = a_q
	return a
}

func (a AppsGetCatalogRequest) WithGenreId(a_genre_id int) AppsGetCatalogRequest{
	a["genre_id"] = a_genre_id
	return a
}

func (a AppsGetCatalogRequest) WithFilter(a_filter string) AppsGetCatalogRequest{
	a["filter"] = a_filter
	return a
}

func (a AppsGetCatalogRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/apps.getCatalog
func (a *Apps) AppsGetCatalog(params ...api.MethodParams) (resp models.AppsGetCatalogResponse, err error) {
	req := api.NewRequest[models.AppsGetCatalogResponse](a.api)

	res, err := req.Execute("apps.getCatalog", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

