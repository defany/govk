// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// DatabaseGetRegions Returns a list of regions.
type DatabaseGetRegionsRequest api.Params

func NewDatabaseGetRegionsRequest() DatabaseGetRegionsRequest {
	params := make(DatabaseGetRegionsRequest, 5)
	return params
}

func (d DatabaseGetRegionsRequest) WithCountryId(d_country_id int) DatabaseGetRegionsRequest{
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetRegionsRequest) WithQ(d_q string) DatabaseGetRegionsRequest{
	d["q"] = d_q
	return d
}

func (d DatabaseGetRegionsRequest) WithOffset(d_offset int) DatabaseGetRegionsRequest{
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetRegionsRequest) WithCount(d_count int) DatabaseGetRegionsRequest{
	d["count"] = d_count
	return d
}

func (d DatabaseGetRegionsRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getRegions
func (d *Database) DatabaseGetRegions(params ...api.MethodParams) (resp models.DatabaseGetRegionsResponse, err error) {
	req := api.NewRequest[models.DatabaseGetRegionsResponse](d.api)

	res, err := req.Execute("database.getRegions", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

