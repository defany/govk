// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StoriesGetVideoUploadServer Allows to receive URL for uploading story with video.
type StoriesGetVideoUploadServerRequest api.Params

func NewStoriesGetVideoUploadServerRequest() StoriesGetVideoUploadServerRequest {
	params := make(StoriesGetVideoUploadServerRequest, 8)
	return params
}

func (s StoriesGetVideoUploadServerRequest) WithAddToNews(s_add_to_news bool) StoriesGetVideoUploadServerRequest{
	s["add_to_news"] = s_add_to_news
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithUserIds(s_user_ids []int) StoriesGetVideoUploadServerRequest{
	s["user_ids"] = s_user_ids
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithReplyToStory(s_reply_to_story string) StoriesGetVideoUploadServerRequest{
	s["reply_to_story"] = s_reply_to_story
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithLinkText(s_link_text models.StoriesUploadLinkText) StoriesGetVideoUploadServerRequest{
	s["link_text"] = s_link_text
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithLinkUrl(s_link_url string) StoriesGetVideoUploadServerRequest{
	s["link_url"] = s_link_url
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithGroupId(s_group_id int) StoriesGetVideoUploadServerRequest{
	s["group_id"] = s_group_id
	return s
}

func (s StoriesGetVideoUploadServerRequest) WithClickableStickers(s_clickable_stickers string) StoriesGetVideoUploadServerRequest{
	s["clickable_stickers"] = s_clickable_stickers
	return s
}

func (s StoriesGetVideoUploadServerRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_MessagesUserBlocked, Error_StoryIncorrectReplyPrivacy, Error_Blocked ]
//
// https://dev.vk.com/method/stories.getVideoUploadServer
func (s *Stories) StoriesGetVideoUploadServer(params ...api.MethodParams) (resp models.StoriesGetVideoUploadServerResponse, err error) {
	req := api.NewRequest[models.StoriesGetVideoUploadServerResponse](s.api)

	res, err := req.Execute("stories.getVideoUploadServer", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

