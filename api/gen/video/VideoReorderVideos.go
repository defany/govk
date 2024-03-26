// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// VideoReorderVideos Reorders the video in the video album.
type VideoReorderVideosRequest api.Params

func NewVideoReorderVideosRequest() VideoReorderVideosRequest {
	params := make(VideoReorderVideosRequest, 9)
	return params
}

func (v VideoReorderVideosRequest) WithTargetId(v_target_id int) VideoReorderVideosRequest{
	v["target_id"] = v_target_id
	return v
}

func (v VideoReorderVideosRequest) WithAlbumId(v_album_id int) VideoReorderVideosRequest{
	v["album_id"] = v_album_id
	return v
}

func (v VideoReorderVideosRequest) WithOwnerId(v_owner_id int) VideoReorderVideosRequest{
	v["owner_id"] = v_owner_id
	return v
}

func (v VideoReorderVideosRequest) WithVideoId(v_video_id int) VideoReorderVideosRequest{
	v["video_id"] = v_video_id
	return v
}

func (v VideoReorderVideosRequest) WithBeforeOwnerId(v_before_owner_id int) VideoReorderVideosRequest{
	v["before_owner_id"] = v_before_owner_id
	return v
}

func (v VideoReorderVideosRequest) WithBeforeVideoId(v_before_video_id int) VideoReorderVideosRequest{
	v["before_video_id"] = v_before_video_id
	return v
}

func (v VideoReorderVideosRequest) WithAfterOwnerId(v_after_owner_id int) VideoReorderVideosRequest{
	v["after_owner_id"] = v_after_owner_id
	return v
}

func (v VideoReorderVideosRequest) WithAfterVideoId(v_after_video_id int) VideoReorderVideosRequest{
	v["after_video_id"] = v_after_video_id
	return v
}

func (v VideoReorderVideosRequest) Params() api.Params {
	return api.Params(v)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AccessVideo ]
//
// https://dev.vk.com/method/video.reorderVideos
func (v *Video) VideoReorderVideos(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](v.api)

	res, err := req.Execute("video.reorderVideos", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
