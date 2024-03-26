// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// BoardGetTopics Returns a list of topics on a community's discussion board.
type BoardGetTopicsRequest api.Params

func NewBoardGetTopicsRequest() BoardGetTopicsRequest {
	params := make(BoardGetTopicsRequest, 10)
	return params
}

func (b BoardGetTopicsRequest) WithGroupId(b_group_id int) BoardGetTopicsRequest{
	b["group_id"] = b_group_id
	return b
}

func (b BoardGetTopicsRequest) WithTopicIds(b_topic_ids []int) BoardGetTopicsRequest{
	b["topic_ids"] = b_topic_ids
	return b
}

func (b BoardGetTopicsRequest) WithOrder(b_order int) BoardGetTopicsRequest{
	b["order"] = b_order
	return b
}

func (b BoardGetTopicsRequest) WithOffset(b_offset int) BoardGetTopicsRequest{
	b["offset"] = b_offset
	return b
}

func (b BoardGetTopicsRequest) WithCount(b_count int) BoardGetTopicsRequest{
	b["count"] = b_count
	return b
}

func (b BoardGetTopicsRequest) WithExtended(b_extended bool) BoardGetTopicsRequest{
	b["extended"] = b_extended
	return b
}

func (b BoardGetTopicsRequest) WithPreview(b_preview int) BoardGetTopicsRequest{
	b["preview"] = b_preview
	return b
}

func (b BoardGetTopicsRequest) WithPreviewLength(b_preview_length int) BoardGetTopicsRequest{
	b["preview_length"] = b_preview_length
	return b
}

func (b BoardGetTopicsRequest) Params() api.Params {
	return api.Params(b)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getTopics
func (b *Board) BoardGetTopics(params ...api.MethodParams) (resp models.BoardGetTopicsResponse, err error) {
	req := api.NewRequest[models.BoardGetTopicsResponse](b.api)

	res, err := req.Execute("board.getTopics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// BoardGetTopicsExtended Returns a list of topics on a community's discussion board.
// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/board.getTopics
func (b *Board) BoardGetTopicsExtended(params ...api.MethodParams) (resp models.BoardGetTopicsExtendedResponse, err error) {
	req := api.NewRequest[models.BoardGetTopicsExtendedResponse](b.api)

	res, err := req.Execute("board.getTopics", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

