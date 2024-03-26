// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// DatabaseGetCountries Returns a list of countries.
type DatabaseGetCountriesRequest api.Params

func NewDatabaseGetCountriesRequest() DatabaseGetCountriesRequest {
	params := make(DatabaseGetCountriesRequest, 5)
	return params
}

func (d DatabaseGetCountriesRequest) WithNeedAll(d_need_all bool) DatabaseGetCountriesRequest{
	d["need_all"] = d_need_all
	return d
}

func (d DatabaseGetCountriesRequest) WithCode(d_code string) DatabaseGetCountriesRequest{
	d["code"] = d_code
	return d
}

func (d DatabaseGetCountriesRequest) WithOffset(d_offset int) DatabaseGetCountriesRequest{
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetCountriesRequest) WithCount(d_count int) DatabaseGetCountriesRequest{
	d["count"] = d_count
	return d
}

func (d DatabaseGetCountriesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getCountries
func (d *Database) DatabaseGetCountries(params ...api.MethodParams) (resp models.DatabaseGetCountriesResponse, err error) {
	req := api.NewRequest[models.DatabaseGetCountriesResponse](d.api)

	res, err := req.Execute("database.getCountries", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

