// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// DatabaseGetSchoolClasses Returns a list of school classes specified for the country.
type DatabaseGetSchoolClassesRequest api.Params

func NewDatabaseGetSchoolClassesRequest() DatabaseGetSchoolClassesRequest {
	params := make(DatabaseGetSchoolClassesRequest, 2)
	return params
}

func (d DatabaseGetSchoolClassesRequest) WithCountryId(d_country_id int) DatabaseGetSchoolClassesRequest{
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetSchoolClassesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getSchoolClasses
func (d *Database) DatabaseGetSchoolClasses(params ...api.MethodParams) (resp models.DatabaseGetSchoolClassesResponse, err error) {
	req := api.NewRequest[models.DatabaseGetSchoolClassesResponse](d.api)

	res, err := req.Execute("database.getSchoolClasses", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

