// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// VideoGet Returns detailed information about videos.
type VideoGetRequest api.Params

func NewVideoGetRequest() VideoGetRequest {
	params := make(VideoGetRequest, 8)
	return params
}

func (v VideoGetRequest) WithOwnerId(v_owner_id int) VideoGetRequest{
	v["owner_id"] = v_owner_id
	return v
}

func (v VideoGetRequest) WithVideos(v_videos []string) VideoGetRequest{
	v["videos"] = v_videos
	return v
}

func (v VideoGetRequest) WithAlbumId(v_album_id int) VideoGetRequest{
	v["album_id"] = v_album_id
	return v
}

func (v VideoGetRequest) WithCount(v_count int) VideoGetRequest{
	v["count"] = v_count
	return v
}

func (v VideoGetRequest) WithOffset(v_offset int) VideoGetRequest{
	v["offset"] = v_offset
	return v
}

func (v VideoGetRequest) WithExtended(v_extended bool) VideoGetRequest{
	v["extended"] = v_extended
	return v
}

func (v VideoGetRequest) WithFields(v_fields []string) VideoGetRequest{
	v["fields"] = v_fields
	return v
}

func (v VideoGetRequest) Params() api.Params {
	return api.Params(v)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AccessVideo ]
//
// https://dev.vk.com/method/video.get
func (v *Video) VideoGet(params ...api.MethodParams) (resp models.VideoGetResponse, err error) {
	req := api.NewRequest[models.VideoGetResponse](v.api)

	res, err := req.Execute("video.get", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

