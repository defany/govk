// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// BoardCreateComment Adds a comment on a topic on a community's discussion board.
type BoardCreateCommentRequest api.Params

func NewBoardCreateCommentRequest() BoardCreateCommentRequest {
	params := make(BoardCreateCommentRequest, 8)
	return params
}

func (b BoardCreateCommentRequest) WithGroupId(b_group_id int) BoardCreateCommentRequest{
	b["group_id"] = b_group_id
	return b
}

func (b BoardCreateCommentRequest) WithTopicId(b_topic_id int) BoardCreateCommentRequest{
	b["topic_id"] = b_topic_id
	return b
}

func (b BoardCreateCommentRequest) WithMessage(b_message string) BoardCreateCommentRequest{
	b["message"] = b_message
	return b
}

func (b BoardCreateCommentRequest) WithAttachments(b_attachments []string) BoardCreateCommentRequest{
	b["attachments"] = b_attachments
	return b
}

func (b BoardCreateCommentRequest) WithFromGroup(b_from_group bool) BoardCreateCommentRequest{
	b["from_group"] = b_from_group
	return b
}

func (b BoardCreateCommentRequest) WithStickerId(b_sticker_id int) BoardCreateCommentRequest{
	b["sticker_id"] = b_sticker_id
	return b
}

func (b BoardCreateCommentRequest) WithGuid(b_guid string) BoardCreateCommentRequest{
	b["guid"] = b_guid
	return b
}

func (b BoardCreateCommentRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.createComment
func (b *Board) BoardCreateComment(params ...api.MethodParams) (resp models.BoardCreateCommentResponse, err error) {
	req := api.NewRequest[models.BoardCreateCommentResponse](b.api)

	res, err := req.Execute("board.createComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

