// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// BoardEditComment Edits a comment on a topic on a community's discussion board.
type BoardEditCommentRequest api.Params

func NewBoardEditCommentRequest() BoardEditCommentRequest {
	params := make(BoardEditCommentRequest, 6)
	return params
}

func (b BoardEditCommentRequest) WithGroupId(b_group_id int) BoardEditCommentRequest{
	b["group_id"] = b_group_id
	return b
}

func (b BoardEditCommentRequest) WithTopicId(b_topic_id int) BoardEditCommentRequest{
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardEditCommentRequest) WithCommentId(b_comment_id int) BoardEditCommentRequest{
	b["comment_id"] = b_comment_id
	return b
}

func (b BoardEditCommentRequest) WithMessage(b_message string) BoardEditCommentRequest{
	b["message"] = b_message
	return b
}

func (b BoardEditCommentRequest) WithAttachments(b_attachments []string) BoardEditCommentRequest{
	b["attachments"] = b_attachments
	return b
}

func (b BoardEditCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.editComment
func (b *Board) BoardEditComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](b.api)

	res, err := req.Execute("board.editComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
