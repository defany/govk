// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PhotosSaveOwnerCoverPhoto Saves cover photo after successful uploading.
type PhotosSaveOwnerCoverPhotoRequest api.Params

func NewPhotosSaveOwnerCoverPhotoRequest() PhotosSaveOwnerCoverPhotoRequest {
	params := make(PhotosSaveOwnerCoverPhotoRequest, 3)
	return params
}

func (p PhotosSaveOwnerCoverPhotoRequest) WithHash(p_hash string) PhotosSaveOwnerCoverPhotoRequest{
	p["hash"] = p_hash
	return p
}

func (p PhotosSaveOwnerCoverPhotoRequest) WithPhoto(p_photo string) PhotosSaveOwnerCoverPhotoRequest{
	p["photo"] = p_photo
	return p
}

func (p PhotosSaveOwnerCoverPhotoRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_ParamPhoto ]
//
// https://dev.vk.com/method/photos.saveOwnerCoverPhoto
func (p *Photos) PhotosSaveOwnerCoverPhoto(params ...api.MethodParams) (resp models.PhotosSaveOwnerCoverPhotoResponse, err error) {
	req := api.NewRequest[models.PhotosSaveOwnerCoverPhotoResponse](p.api)

	res, err := req.Execute("photos.saveOwnerCoverPhoto", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

