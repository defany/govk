// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PhotosSaveOwnerPhoto Saves a profile or community photo. Upload URL can be got with the [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
type PhotosSaveOwnerPhotoRequest api.Params

func NewPhotosSaveOwnerPhotoRequest() PhotosSaveOwnerPhotoRequest {
	params := make(PhotosSaveOwnerPhotoRequest, 4)
	return params
}

func (p PhotosSaveOwnerPhotoRequest) WithServer(p_server string) PhotosSaveOwnerPhotoRequest{
	p["server"] = p_server
	return p
}

func (p PhotosSaveOwnerPhotoRequest) WithHash(p_hash string) PhotosSaveOwnerPhotoRequest{
	p["hash"] = p_hash
	return p
}

func (p PhotosSaveOwnerPhotoRequest) WithPhoto(p_photo string) PhotosSaveOwnerPhotoRequest{
	p["photo"] = p_photo
	return p
}

func (p PhotosSaveOwnerPhotoRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_ParamPhoto ]
//
// https://dev.vk.com/method/photos.saveOwnerPhoto
func (p *Photos) PhotosSaveOwnerPhoto(params ...api.MethodParams) (resp models.PhotosSaveOwnerPhotoResponse, err error) {
	req := api.NewRequest[models.PhotosSaveOwnerPhotoResponse](p.api)

	res, err := req.Execute("photos.saveOwnerPhoto", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
