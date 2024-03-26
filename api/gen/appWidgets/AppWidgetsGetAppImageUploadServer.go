// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AppWidgetsGetAppImageUploadServer Returns a URL for uploading a photo to the community collection for community app widgets
type AppWidgetsGetAppImageUploadServerRequest api.Params

func NewAppWidgetsGetAppImageUploadServerRequest() AppWidgetsGetAppImageUploadServerRequest {
	params := make(AppWidgetsGetAppImageUploadServerRequest, 2)
	return params
}

func (a AppWidgetsGetAppImageUploadServerRequest) WithImageType(a_image_type string) AppWidgetsGetAppImageUploadServerRequest{
	a["image_type"] = a_image_type
	return a
}

func (a AppWidgetsGetAppImageUploadServerRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/appWidgets.getAppImageUploadServer
func (a *AppWidgets) AppWidgetsGetAppImageUploadServer(params ...api.MethodParams) (resp models.AppWidgetsGetAppImageUploadServerResponse, err error) {
	req := api.NewRequest[models.AppWidgetsGetAppImageUploadServerResponse](a.api)

	res, err := req.Execute("appWidgets.getAppImageUploadServer", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

