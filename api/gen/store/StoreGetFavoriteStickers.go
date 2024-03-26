// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StoreGetFavoriteStickers ...
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/store.getFavoriteStickers
func (s *Store) StoreGetFavoriteStickers(params ...api.MethodParams) (resp models.StoreGetFavoriteStickersResponse, err error) {
	req := api.NewRequest[models.StoreGetFavoriteStickersResponse](s.api)

	res, err := req.Execute("store.getFavoriteStickers", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
