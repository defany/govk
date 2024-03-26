// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// DocsDelete Deletes a user or community document.
type DocsDeleteRequest api.Params

func NewDocsDeleteRequest() DocsDeleteRequest {
	params := make(DocsDeleteRequest, 3)
	return params
}

func (d DocsDeleteRequest) WithOwnerId(d_owner_id int) DocsDeleteRequest{
	d["owner_id"] = d_owner_id
	return d
}

func (d DocsDeleteRequest) WithDocId(d_doc_id int) DocsDeleteRequest{
	d["doc_id"] = d_doc_id
	return d
}

func (d DocsDeleteRequest) Params() api.Params {
	return api.Params(d)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_ParamDocDeleteAccess, Error_ParamDocId ]
//
// https://dev.vk.com/method/docs.delete
func (d *Docs) DocsDelete(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](d.api)

	res, err := req.Execute("docs.delete", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

