// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// WallGetById Returns a list of posts from user or community walls by their IDs.
type WallGetByIdRequest api.Params

func NewWallGetByIdRequest() WallGetByIdRequest {
	params := make(WallGetByIdRequest, 6)
	return params
}

func (w WallGetByIdRequest) WithPosts(w_posts []string) WallGetByIdRequest{
	w["posts"] = w_posts
	return w
}

func (w WallGetByIdRequest) WithExtended(w_extended bool) WallGetByIdRequest{
	w["extended"] = w_extended
	return w
}

func (w WallGetByIdRequest) WithCopyHistoryDepth(w_copy_history_depth int) WallGetByIdRequest{
	w["copy_history_depth"] = w_copy_history_depth
	return w
}

func (w WallGetByIdRequest) WithFields(w_fields []models.BaseUserGroupFields) WallGetByIdRequest{
	w["fields"] = w_fields
	return w
}

func (w WallGetByIdRequest) Params() api.Params {
	return api.Params(w)
}

// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/wall.getById
func (w *Wall) WallGetById(params ...api.MethodParams) (resp models.WallGetByIdLegacyResponse, err error) {
	req := api.NewRequest[models.WallGetByIdLegacyResponse](w.api)

	res, err := req.Execute("wall.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// WallGetByIdExtended Returns a list of posts from user or community walls by their IDs.
// May execute with listed access token types:
//    [ user, service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/wall.getById
func (w *Wall) WallGetByIdExtended(params ...api.MethodParams) (resp models.WallGetByIdExtendedResponse, err error) {
	req := api.NewRequest[models.WallGetByIdExtendedResponse](w.api)

	res, err := req.Execute("wall.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
