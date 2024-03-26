// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// StoreAddStickersToFavorite Adds given sticker IDs to the list of user's favorite stickers
type StoreAddStickersToFavoriteRequest api.Params

func NewStoreAddStickersToFavoriteRequest() StoreAddStickersToFavoriteRequest {
	params := make(StoreAddStickersToFavoriteRequest, 2)
	return params
}

func (s StoreAddStickersToFavoriteRequest) WithStickerIds(s_sticker_ids []int) StoreAddStickersToFavoriteRequest{
	s["sticker_ids"] = s_sticker_ids
	return s
}

func (s StoreAddStickersToFavoriteRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_StickersNotPurchased, Error_StickersTooManyFavorites ]
//
// https://dev.vk.com/method/store.addStickersToFavorite
func (s *Store) StoreAddStickersToFavorite(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](s.api)

	res, err := req.Execute("store.addStickersToFavorite", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

