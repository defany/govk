// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// VideoReportComment Reports (submits a complaint about) a comment on a video.
type VideoReportCommentRequest api.Params

func NewVideoReportCommentRequest() VideoReportCommentRequest {
	params := make(VideoReportCommentRequest, 4)
	return params
}

func (v VideoReportCommentRequest) WithOwnerId(v_owner_id int) VideoReportCommentRequest{
	v["owner_id"] = v_owner_id
	return v
}

func (v VideoReportCommentRequest) WithCommentId(v_comment_id int) VideoReportCommentRequest{
	v["comment_id"] = v_comment_id
	return v
}

func (v VideoReportCommentRequest) WithReason(v_reason int) VideoReportCommentRequest{
	v["reason"] = v_reason
	return v
}

func (v VideoReportCommentRequest) Params() api.Params {
	return api.Params(v)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/video.reportComment
func (v *Video) VideoReportComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](v.api)

	res, err := req.Execute("video.reportComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
