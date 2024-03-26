// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AdsGetSuggestions Returns a set of auto-suggestions for various targeting parameters.
type AdsGetSuggestionsRequest api.Params

func NewAdsGetSuggestionsRequest() AdsGetSuggestionsRequest {
	params := make(AdsGetSuggestionsRequest, 7)
	return params
}

func (a AdsGetSuggestionsRequest) WithSection(a_section string) AdsGetSuggestionsRequest{
	a["section"] = a_section
	return a
}

func (a AdsGetSuggestionsRequest) WithIds(a_ids string) AdsGetSuggestionsRequest{
	a["ids"] = a_ids
	return a
}

func (a AdsGetSuggestionsRequest) WithQ(a_q string) AdsGetSuggestionsRequest{
	a["q"] = a_q
	return a
}

func (a AdsGetSuggestionsRequest) WithCountry(a_country int) AdsGetSuggestionsRequest{
	a["country"] = a_country
	return a
}

func (a AdsGetSuggestionsRequest) WithCities(a_cities string) AdsGetSuggestionsRequest{
	a["cities"] = a_cities
	return a
}

func (a AdsGetSuggestionsRequest) WithLang(a_lang string) AdsGetSuggestionsRequest{
	a["lang"] = a_lang
	return a
}

func (a AdsGetSuggestionsRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/ads.getSuggestions
func (a *Ads) AdsGetSuggestions(params ...api.MethodParams) (resp models.AdsGetSuggestionsResponse, err error) {
	req := api.NewRequest[models.AdsGetSuggestionsResponse](a.api)

	res, err := req.Execute("ads.getSuggestions", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

