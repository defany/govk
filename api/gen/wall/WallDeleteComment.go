// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// WallDeleteComment Deletes a comment on a post on a user wall or community wall.
type WallDeleteCommentRequest api.Params

func NewWallDeleteCommentRequest() WallDeleteCommentRequest {
	params := make(WallDeleteCommentRequest, 3)
	return params
}

func (w WallDeleteCommentRequest) WithOwnerId(w_owner_id int) WallDeleteCommentRequest{
	w["owner_id"] = w_owner_id
	return w
}

func (w WallDeleteCommentRequest) WithCommentId(w_comment_id int) WallDeleteCommentRequest{
	w["comment_id"] = w_comment_id
	return w
}

func (w WallDeleteCommentRequest) Params() api.Params {
	return api.Params(w)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallAccessComment ]
//
// https://dev.vk.com/method/wall.deleteComment
func (w *Wall) WallDeleteComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](w.api)

	res, err := req.Execute("wall.deleteComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
