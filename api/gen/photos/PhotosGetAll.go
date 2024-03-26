// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PhotosGetAll Returns a list of photos belonging to a user or community, in reverse chronological order.
type PhotosGetAllRequest api.Params

func NewPhotosGetAllRequest() PhotosGetAllRequest {
	params := make(PhotosGetAllRequest, 10)
	return params
}

func (p PhotosGetAllRequest) WithOwnerId(p_owner_id int) PhotosGetAllRequest{
	p["owner_id"] = p_owner_id
	return p
}

func (p PhotosGetAllRequest) WithExtended(p_extended bool) PhotosGetAllRequest{
	p["extended"] = p_extended
	return p
}

func (p PhotosGetAllRequest) WithOffset(p_offset int) PhotosGetAllRequest{
	p["offset"] = p_offset
	return p
}

func (p PhotosGetAllRequest) WithCount(p_count int) PhotosGetAllRequest{
	p["count"] = p_count
	return p
}

func (p PhotosGetAllRequest) WithPhotoSizes(p_photo_sizes bool) PhotosGetAllRequest{
	p["photo_sizes"] = p_photo_sizes
	return p
}

func (p PhotosGetAllRequest) WithNoServiceAlbums(p_no_service_albums bool) PhotosGetAllRequest{
	p["no_service_albums"] = p_no_service_albums
	return p
}

func (p PhotosGetAllRequest) WithNeedHidden(p_need_hidden bool) PhotosGetAllRequest{
	p["need_hidden"] = p_need_hidden
	return p
}

func (p PhotosGetAllRequest) WithSkipHidden(p_skip_hidden bool) PhotosGetAllRequest{
	p["skip_hidden"] = p_skip_hidden
	return p
}

func (p PhotosGetAllRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_Blocked ]
//
// https://dev.vk.com/method/photos.getAll
func (p *Photos) PhotosGetAll(params ...api.MethodParams) (resp models.PhotosGetAllResponse, err error) {
	req := api.NewRequest[models.PhotosGetAllResponse](p.api)

	res, err := req.Execute("photos.getAll", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// PhotosGetAllExtended Returns a list of photos belonging to a user or community, in reverse chronological order.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_Blocked ]
//
// https://dev.vk.com/method/photos.getAll
func (p *Photos) PhotosGetAllExtended(params ...api.MethodParams) (resp models.PhotosGetAllExtendedResponse, err error) {
	req := api.NewRequest[models.PhotosGetAllExtendedResponse](p.api)

	res, err := req.Execute("photos.getAll", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

