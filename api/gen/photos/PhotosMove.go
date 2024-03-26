// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PhotosMove Moves a photo from one album to another.
type PhotosMoveRequest api.Params

func NewPhotosMoveRequest() PhotosMoveRequest {
	params := make(PhotosMoveRequest, 4)
	return params
}

func (p PhotosMoveRequest) WithOwnerId(p_owner_id int) PhotosMoveRequest{
	p["owner_id"] = p_owner_id
	return p
}

func (p PhotosMoveRequest) WithTargetAlbumId(p_target_album_id int) PhotosMoveRequest{
	p["target_album_id"] = p_target_album_id
	return p
}

func (p PhotosMoveRequest) WithPhotoIds(p_photo_ids int) PhotosMoveRequest{
	p["photo_ids"] = p_photo_ids
	return p
}

func (p PhotosMoveRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/photos.move
func (p *Photos) PhotosMove(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](p.api)

	res, err := req.Execute("photos.move", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
