// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// VideoRestore Restores a previously deleted video.
type VideoRestoreRequest api.Params

func NewVideoRestoreRequest() VideoRestoreRequest {
	params := make(VideoRestoreRequest, 3)
	return params
}

func (v VideoRestoreRequest) WithVideoId(v_video_id int) VideoRestoreRequest{
	v["video_id"] = v_video_id
	return v
}

func (v VideoRestoreRequest) WithOwnerId(v_owner_id int) VideoRestoreRequest{
	v["owner_id"] = v_owner_id
	return v
}

func (v VideoRestoreRequest) Params() api.Params {
	return api.Params(v)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/video.restore
func (v *Video) VideoRestore(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](v.api)

	res, err := req.Execute("video.restore", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
