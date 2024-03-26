// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// DatabaseGetCitiesById Returns information about cities by their IDs.
type DatabaseGetCitiesByIdRequest api.Params

func NewDatabaseGetCitiesByIdRequest() DatabaseGetCitiesByIdRequest {
	params := make(DatabaseGetCitiesByIdRequest, 2)
	return params
}

func (d DatabaseGetCitiesByIdRequest) WithCityIds(d_city_ids []int) DatabaseGetCitiesByIdRequest{
	d["city_ids"] = d_city_ids
	return d
}

func (d DatabaseGetCitiesByIdRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getCitiesById
func (d *Database) DatabaseGetCitiesById(params ...api.MethodParams) (resp models.DatabaseGetCitiesByIdResponse, err error) {
	req := api.NewRequest[models.DatabaseGetCitiesByIdResponse](d.api)

	res, err := req.Execute("database.getCitiesById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

