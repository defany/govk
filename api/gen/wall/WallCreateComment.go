// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// WallCreateComment Adds a comment to a post on a user wall or community wall.
type WallCreateCommentRequest api.Params

func NewWallCreateCommentRequest() WallCreateCommentRequest {
	params := make(WallCreateCommentRequest, 9)
	return params
}

func (w WallCreateCommentRequest) WithOwnerId(w_owner_id int) WallCreateCommentRequest{
	w["owner_id"] = w_owner_id
	return w
}

func (w WallCreateCommentRequest) WithPostId(w_post_id int) WallCreateCommentRequest{
	w["post_id"] = w_post_id
	return w
}

func (w WallCreateCommentRequest) WithFromGroup(w_from_group int) WallCreateCommentRequest{
	w["from_group"] = w_from_group
	return w
}

func (w WallCreateCommentRequest) WithMessage(w_message string) WallCreateCommentRequest{
	w["message"] = w_message
	return w
}

func (w WallCreateCommentRequest) WithReplyToComment(w_reply_to_comment int) WallCreateCommentRequest{
	w["reply_to_comment"] = w_reply_to_comment
	return w
}

func (w WallCreateCommentRequest) WithAttachments(w_attachments []string) WallCreateCommentRequest{
	w["attachments"] = w_attachments
	return w
}

func (w WallCreateCommentRequest) WithStickerId(w_sticker_id int) WallCreateCommentRequest{
	w["sticker_id"] = w_sticker_id
	return w
}

func (w WallCreateCommentRequest) WithGuid(w_guid string) WallCreateCommentRequest{
	w["guid"] = w_guid
	return w
}

func (w WallCreateCommentRequest) Params() api.Params {
	return api.Params(w)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_WallAccessAddReply, Error_WallReplyOwnerFlood, Error_WallLinksForbidden, Error_WallAccessReplies ]
//
// https://dev.vk.com/method/wall.createComment
func (w *Wall) WallCreateComment(params ...api.MethodParams) (resp models.WallCreateCommentResponse, err error) {
	req := api.NewRequest[models.WallCreateCommentResponse](w.api)

	res, err := req.Execute("wall.createComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
