// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// VideoAddAlbum Creates an empty album for videos.
type VideoAddAlbumRequest api.Params

func NewVideoAddAlbumRequest() VideoAddAlbumRequest {
	params := make(VideoAddAlbumRequest, 4)
	return params
}

func (v VideoAddAlbumRequest) WithGroupId(v_group_id int) VideoAddAlbumRequest{
	v["group_id"] = v_group_id
	return v
}

func (v VideoAddAlbumRequest) WithTitle(v_title string) VideoAddAlbumRequest{
	v["title"] = v_title
	return v
}

func (v VideoAddAlbumRequest) WithPrivacy(v_privacy []string) VideoAddAlbumRequest{
	v["privacy"] = v_privacy
	return v
}

func (v VideoAddAlbumRequest) Params() api.Params {
	return api.Params(v)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AccessVideo, Error_AlbumsLimit ]
//
// https://dev.vk.com/method/video.addAlbum
func (v *Video) VideoAddAlbum(params ...api.MethodParams) (resp models.VideoAddAlbumResponse, err error) {
	req := api.NewRequest[models.VideoAddAlbumResponse](v.api)

	res, err := req.Execute("video.addAlbum", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
