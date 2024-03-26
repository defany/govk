// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PhotosRestoreComment Restores a deleted comment on a photo.
type PhotosRestoreCommentRequest api.Params

func NewPhotosRestoreCommentRequest() PhotosRestoreCommentRequest {
	params := make(PhotosRestoreCommentRequest, 3)
	return params
}

func (p PhotosRestoreCommentRequest) WithOwnerId(p_owner_id int) PhotosRestoreCommentRequest{
	p["owner_id"] = p_owner_id
	return p
}

func (p PhotosRestoreCommentRequest) WithCommentId(p_comment_id int) PhotosRestoreCommentRequest{
	p["comment_id"] = p_comment_id
	return p
}

func (p PhotosRestoreCommentRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/photos.restoreComment
func (p *Photos) PhotosRestoreComment(params ...api.MethodParams) (resp models.PhotosRestoreCommentResponse, err error) {
	req := api.NewRequest[models.PhotosRestoreCommentResponse](p.api)

	res, err := req.Execute("photos.restoreComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
