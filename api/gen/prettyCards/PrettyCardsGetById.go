// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// PrettyCardsGetById ...
type PrettyCardsGetByIdRequest api.Params

func NewPrettyCardsGetByIdRequest() PrettyCardsGetByIdRequest {
	params := make(PrettyCardsGetByIdRequest, 3)
	return params
}

func (p PrettyCardsGetByIdRequest) WithOwnerId(p_owner_id int) PrettyCardsGetByIdRequest{
	p["owner_id"] = p_owner_id
	return p
}

func (p PrettyCardsGetByIdRequest) WithCardIds(p_card_ids []int) PrettyCardsGetByIdRequest{
	p["card_ids"] = p_card_ids
	return p
}

func (p PrettyCardsGetByIdRequest) Params() api.Params {
	return api.Params(p)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/prettyCards.getById
func (p *PrettyCards) PrettyCardsGetById(params ...api.MethodParams) (resp models.PrettyCardsGetByIdResponse, err error) {
	req := api.NewRequest[models.PrettyCardsGetByIdResponse](p.api)

	res, err := req.Execute("prettyCards.getById", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

