// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

type Database struct {
	api *api.API
}

func NewDatabase(api *api.API) *Database {
	return &Database{
		api: api,
	}
}

// DatabaseGetChairs Returns list of chairs on a specified faculty.
type DatabaseGetChairsRequest api.Params

func NewDatabaseGetChairsRequest() DatabaseGetChairsRequest {
	params := make(DatabaseGetChairsRequest, 4)
	return params
}

func (d DatabaseGetChairsRequest) WithFacultyId(d_faculty_id int) DatabaseGetChairsRequest {
	d["faculty_id"] = d_faculty_id
	return d
}

func (d DatabaseGetChairsRequest) WithOffset(d_offset int) DatabaseGetChairsRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetChairsRequest) WithCount(d_count int) DatabaseGetChairsRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetChairsRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getChairs
func (d *Database) DatabaseGetChairs(params ...api.MethodParams) (resp models.DatabaseGetChairsResponse, err error) {
	req := api.NewRequest[models.DatabaseGetChairsResponse](d.api)

	res, err := req.Execute("database.getChairs", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetCities Returns a list of cities.
type DatabaseGetCitiesRequest api.Params

func NewDatabaseGetCitiesRequest() DatabaseGetCitiesRequest {
	params := make(DatabaseGetCitiesRequest, 7)
	return params
}

func (d DatabaseGetCitiesRequest) WithCountryId(d_country_id int) DatabaseGetCitiesRequest {
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetCitiesRequest) WithRegionId(d_region_id int) DatabaseGetCitiesRequest {
	d["region_id"] = d_region_id
	return d
}

func (d DatabaseGetCitiesRequest) WithQ(d_q string) DatabaseGetCitiesRequest {
	d["q"] = d_q
	return d
}

func (d DatabaseGetCitiesRequest) WithNeedAll(d_need_all bool) DatabaseGetCitiesRequest {
	d["need_all"] = d_need_all
	return d
}

func (d DatabaseGetCitiesRequest) WithOffset(d_offset int) DatabaseGetCitiesRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetCitiesRequest) WithCount(d_count int) DatabaseGetCitiesRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetCitiesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getCities
func (d *Database) DatabaseGetCities(params ...api.MethodParams) (resp models.DatabaseGetCitiesResponse, err error) {
	req := api.NewRequest[models.DatabaseGetCitiesResponse](d.api)

	res, err := req.Execute("database.getCities", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetCitiesById Returns information about cities by their IDs.
type DatabaseGetCitiesByIdRequest api.Params

func NewDatabaseGetCitiesByIdRequest() DatabaseGetCitiesByIdRequest {
	params := make(DatabaseGetCitiesByIdRequest, 2)
	return params
}

func (d DatabaseGetCitiesByIdRequest) WithCityIds(d_city_ids []int) DatabaseGetCitiesByIdRequest {
	d["city_ids"] = d_city_ids
	return d
}

func (d DatabaseGetCitiesByIdRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
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

// DatabaseGetCountries Returns a list of countries.
type DatabaseGetCountriesRequest api.Params

func NewDatabaseGetCountriesRequest() DatabaseGetCountriesRequest {
	params := make(DatabaseGetCountriesRequest, 5)
	return params
}

func (d DatabaseGetCountriesRequest) WithNeedAll(d_need_all bool) DatabaseGetCountriesRequest {
	d["need_all"] = d_need_all
	return d
}

func (d DatabaseGetCountriesRequest) WithCode(d_code string) DatabaseGetCountriesRequest {
	d["code"] = d_code
	return d
}

func (d DatabaseGetCountriesRequest) WithOffset(d_offset int) DatabaseGetCountriesRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetCountriesRequest) WithCount(d_count int) DatabaseGetCountriesRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetCountriesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user ]
//
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

// DatabaseGetCountriesById Returns information about countries by their IDs.
type DatabaseGetCountriesByIdRequest api.Params

func NewDatabaseGetCountriesByIdRequest() DatabaseGetCountriesByIdRequest {
	params := make(DatabaseGetCountriesByIdRequest, 2)
	return params
}

func (d DatabaseGetCountriesByIdRequest) WithCountryIds(d_country_ids []int) DatabaseGetCountriesByIdRequest {
	d["country_ids"] = d_country_ids
	return d
}

func (d DatabaseGetCountriesByIdRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getCountriesById
func (d *Database) DatabaseGetCountriesById(params ...api.MethodParams) (resp models.DatabaseGetCountriesByIdResponse, err error) {
	req := api.NewRequest[models.DatabaseGetCountriesByIdResponse](d.api)

	res, err := req.Execute("database.getCountriesById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetFaculties Returns a list of faculties (i.e., university departments).
type DatabaseGetFacultiesRequest api.Params

func NewDatabaseGetFacultiesRequest() DatabaseGetFacultiesRequest {
	params := make(DatabaseGetFacultiesRequest, 4)
	return params
}

func (d DatabaseGetFacultiesRequest) WithUniversityId(d_university_id int) DatabaseGetFacultiesRequest {
	d["university_id"] = d_university_id
	return d
}

func (d DatabaseGetFacultiesRequest) WithOffset(d_offset int) DatabaseGetFacultiesRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetFacultiesRequest) WithCount(d_count int) DatabaseGetFacultiesRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetFacultiesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getFaculties
func (d *Database) DatabaseGetFaculties(params ...api.MethodParams) (resp models.DatabaseGetFacultiesResponse, err error) {
	req := api.NewRequest[models.DatabaseGetFacultiesResponse](d.api)

	res, err := req.Execute("database.getFaculties", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetMetroStations Get metro stations by city
type DatabaseGetMetroStationsRequest api.Params

func NewDatabaseGetMetroStationsRequest() DatabaseGetMetroStationsRequest {
	params := make(DatabaseGetMetroStationsRequest, 5)
	return params
}

func (d DatabaseGetMetroStationsRequest) WithCityId(d_city_id int) DatabaseGetMetroStationsRequest {
	d["city_id"] = d_city_id
	return d
}

func (d DatabaseGetMetroStationsRequest) WithOffset(d_offset int) DatabaseGetMetroStationsRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetMetroStationsRequest) WithCount(d_count int) DatabaseGetMetroStationsRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetMetroStationsRequest) WithExtended(d_extended bool) DatabaseGetMetroStationsRequest {
	d["extended"] = d_extended
	return d
}

func (d DatabaseGetMetroStationsRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getMetroStations
func (d *Database) DatabaseGetMetroStations(params ...api.MethodParams) (resp models.DatabaseGetMetroStationsResponse, err error) {
	req := api.NewRequest[models.DatabaseGetMetroStationsResponse](d.api)

	res, err := req.Execute("database.getMetroStations", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetMetroStationsById Get metro station by his id
type DatabaseGetMetroStationsByIdRequest api.Params

func NewDatabaseGetMetroStationsByIdRequest() DatabaseGetMetroStationsByIdRequest {
	params := make(DatabaseGetMetroStationsByIdRequest, 2)
	return params
}

func (d DatabaseGetMetroStationsByIdRequest) WithStationIds(d_station_ids []int) DatabaseGetMetroStationsByIdRequest {
	d["station_ids"] = d_station_ids
	return d
}

func (d DatabaseGetMetroStationsByIdRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getMetroStationsById
func (d *Database) DatabaseGetMetroStationsById(params ...api.MethodParams) (resp models.DatabaseGetMetroStationsByIdResponse, err error) {
	req := api.NewRequest[models.DatabaseGetMetroStationsByIdResponse](d.api)

	res, err := req.Execute("database.getMetroStationsById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetRegions Returns a list of regions.
type DatabaseGetRegionsRequest api.Params

func NewDatabaseGetRegionsRequest() DatabaseGetRegionsRequest {
	params := make(DatabaseGetRegionsRequest, 5)
	return params
}

func (d DatabaseGetRegionsRequest) WithCountryId(d_country_id int) DatabaseGetRegionsRequest {
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetRegionsRequest) WithQ(d_q string) DatabaseGetRegionsRequest {
	d["q"] = d_q
	return d
}

func (d DatabaseGetRegionsRequest) WithOffset(d_offset int) DatabaseGetRegionsRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetRegionsRequest) WithCount(d_count int) DatabaseGetRegionsRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetRegionsRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
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

// DatabaseGetSchoolClasses Returns a list of school classes specified for the country.
type DatabaseGetSchoolClassesRequest api.Params

func NewDatabaseGetSchoolClassesRequest() DatabaseGetSchoolClassesRequest {
	params := make(DatabaseGetSchoolClassesRequest, 2)
	return params
}

func (d DatabaseGetSchoolClassesRequest) WithCountryId(d_country_id int) DatabaseGetSchoolClassesRequest {
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetSchoolClassesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
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

// DatabaseGetSchools Returns a list of schools.
type DatabaseGetSchoolsRequest api.Params

func NewDatabaseGetSchoolsRequest() DatabaseGetSchoolsRequest {
	params := make(DatabaseGetSchoolsRequest, 5)
	return params
}

func (d DatabaseGetSchoolsRequest) WithQ(d_q string) DatabaseGetSchoolsRequest {
	d["q"] = d_q
	return d
}

func (d DatabaseGetSchoolsRequest) WithCityId(d_city_id int) DatabaseGetSchoolsRequest {
	d["city_id"] = d_city_id
	return d
}

func (d DatabaseGetSchoolsRequest) WithOffset(d_offset int) DatabaseGetSchoolsRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetSchoolsRequest) WithCount(d_count int) DatabaseGetSchoolsRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetSchoolsRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getSchools
func (d *Database) DatabaseGetSchools(params ...api.MethodParams) (resp models.DatabaseGetSchoolsResponse, err error) {
	req := api.NewRequest[models.DatabaseGetSchoolsResponse](d.api)

	res, err := req.Execute("database.getSchools", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// DatabaseGetUniversities Returns a list of higher education institutions.
type DatabaseGetUniversitiesRequest api.Params

func NewDatabaseGetUniversitiesRequest() DatabaseGetUniversitiesRequest {
	params := make(DatabaseGetUniversitiesRequest, 6)
	return params
}

func (d DatabaseGetUniversitiesRequest) WithQ(d_q string) DatabaseGetUniversitiesRequest {
	d["q"] = d_q
	return d
}

func (d DatabaseGetUniversitiesRequest) WithCountryId(d_country_id int) DatabaseGetUniversitiesRequest {
	d["country_id"] = d_country_id
	return d
}

func (d DatabaseGetUniversitiesRequest) WithCityId(d_city_id int) DatabaseGetUniversitiesRequest {
	d["city_id"] = d_city_id
	return d
}

func (d DatabaseGetUniversitiesRequest) WithOffset(d_offset int) DatabaseGetUniversitiesRequest {
	d["offset"] = d_offset
	return d
}

func (d DatabaseGetUniversitiesRequest) WithCount(d_count int) DatabaseGetUniversitiesRequest {
	d["count"] = d_count
	return d
}

func (d DatabaseGetUniversitiesRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//
//	[ user, service ]
//
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/database.getUniversities
func (d *Database) DatabaseGetUniversities(params ...api.MethodParams) (resp models.DatabaseGetUniversitiesResponse, err error) {
	req := api.NewRequest[models.DatabaseGetUniversitiesResponse](d.api)

	res, err := req.Execute("database.getUniversities", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
