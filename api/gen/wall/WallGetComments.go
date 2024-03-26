// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// WallGetComments Returns a list of comments on a post on a user wall or community wall.
type WallGetCommentsRequest api.Params

func NewWallGetCommentsRequest() WallGetCommentsRequest {
	params := make(WallGetCommentsRequest, 14)
	return params
}

func (w WallGetCommentsRequest) WithOwnerId(w_owner_id int) WallGetCommentsRequest{
	w["owner_id"] = w_owner_id
	return w
}

func (w WallGetCommentsRequest) WithPostId(w_post_id int) WallGetCommentsRequest{
	w["post_id"] = w_post_id
	return w
}

func (w WallGetCommentsRequest) WithNeedLikes(w_need_likes bool) WallGetCommentsRequest{
	w["need_likes"] = w_need_likes
	return w
}

func (w WallGetCommentsRequest) WithStartCommentId(w_start_comment_id int) WallGetCommentsRequest{
	w["start_comment_id"] = w_start_comment_id
	return w
}

func (w WallGetCommentsRequest) WithOffset(w_offset int) WallGetCommentsRequest{
	w["offset"] = w_offset
	return w
}

func (w WallGetCommentsRequest) WithCount(w_count int) WallGetCommentsRequest{
	w["count"] = w_count
	return w
}

func (w WallGetCommentsRequest) WithSort(w_sort string) WallGetCommentsRequest{
	w["sort"] = w_sort
	return w
}

func (w WallGetCommentsRequest) WithPreviewLength(w_preview_length int) WallGetCommentsRequest{
	w["preview_length"] = w_preview_length
	return w
}

func (w WallGetCommentsRequest) WithExtended(w_extended bool) WallGetCommentsRequest{
	w["extended"] = w_extended
	return w
}

func (w WallGetCommentsRequest) WithFields(w_fields []models.BaseUserGroupFields) WallGetCommentsRequest{
	w["fields"] = w_fields
	return w
}

func (w WallGetCommentsRequest) WithCommentId(w_comment_id int) WallGetCommentsRequest{
	w["comment_id"] = w_comment_id
	return w
}

func (w WallGetCommentsRequest) WithThreadItemsCount(w_thread_items_count int) WallGetCommentsRequest{
	w["thread_items_count"] = w_thread_items_count
	return w
}

func (w WallGetCommentsRequest) Params() api.Params {
	return api.Params(w)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallAccessReplies ]
//
// https://dev.vk.com/method/wall.getComments
func (w *Wall) WallGetComments(params ...api.MethodParams) (resp models.WallGetCommentsResponse, err error) {
	req := api.NewRequest[models.WallGetCommentsResponse](w.api)

	res, err := req.Execute("wall.getComments", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// WallGetCommentsExtended Returns a list of comments on a post on a user wall or community wall.
// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallAccessReplies ]
//
// https://dev.vk.com/method/wall.getComments
func (w *Wall) WallGetCommentsExtended(params ...api.MethodParams) (resp models.WallGetCommentsExtendedResponse, err error) {
	req := api.NewRequest[models.WallGetCommentsExtendedResponse](w.api)

	res, err := req.Execute("wall.getComments", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
