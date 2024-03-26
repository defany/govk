// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FaveAddProduct ...
type FaveAddProductRequest api.Params

func NewFaveAddProductRequest() FaveAddProductRequest {
	params := make(FaveAddProductRequest, 4)
	return params
}

func (f FaveAddProductRequest) WithOwnerId(f_owner_id int) FaveAddProductRequest{
	f["owner_id"] = f_owner_id
	return f
}

func (f FaveAddProductRequest) WithId(f_id int) FaveAddProductRequest{
	f["id"] = f_id
	return f
}

func (f FaveAddProductRequest) WithAccessKey(f_access_key string) FaveAddProductRequest{
	f["access_key"] = f_access_key
	return f
}

func (f FaveAddProductRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/fave.addProduct
func (f *Fave) FaveAddProduct(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](f.api)

	res, err := req.Execute("fave.addProduct", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

