// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FaveReorderTags ...
type FaveReorderTagsRequest api.Params

func NewFaveReorderTagsRequest() FaveReorderTagsRequest {
	params := make(FaveReorderTagsRequest, 2)
	return params
}

func (f FaveReorderTagsRequest) WithIds(f_ids []int) FaveReorderTagsRequest{
	f["ids"] = f_ids
	return f
}

func (f FaveReorderTagsRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/fave.reorderTags
func (f *Fave) FaveReorderTags(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](f.api)

	res, err := req.Execute("fave.reorderTags", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

