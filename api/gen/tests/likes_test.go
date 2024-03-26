// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/likes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyLikesAddRequest(r *requests.LikesAddRequest) {
	Type := new(models.LikesType)
	fillRandomlyLikesType(Type)
	r.WithType(*Type)
	r.WithOwnerId(randInt())
	r.WithItemId(randInt())
	r.WithAccessKey(randString())
}

func TestVKLikesAddSuccess(t *testing.T) {
	params := requests.NewLikesAddRequest()
	fillRandomlyLikesAddRequest(&params)
	var expected models.LikesAddResponse
	fillRandomlyLikesAddResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "likes.add", params.Params(), expectedJSON))
	resp, err := vk.Api.Likes.LikesAdd(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyLikesDeleteRequest(r *requests.LikesDeleteRequest) {
	Type := new(models.LikesType)
	fillRandomlyLikesType(Type)
	r.WithType(*Type)
	r.WithOwnerId(randInt())
	r.WithItemId(randInt())
	r.WithAccessKey(randString())
}

func TestVKLikesDeleteSuccess(t *testing.T) {
	params := requests.NewLikesDeleteRequest()
	fillRandomlyLikesDeleteRequest(&params)
	var expected models.LikesDeleteResponse
	fillRandomlyLikesDeleteResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "likes.delete", params.Params(), expectedJSON))
	resp, err := vk.Api.Likes.LikesDelete(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyLikesGetListRequest(r *requests.LikesGetListRequest) {
	Type := new(models.LikesType)
	fillRandomlyLikesType(Type)
	r.WithType(*Type)
	r.WithOwnerId(randInt())
	r.WithItemId(randInt())
	r.WithPageUrl(randString())
	r.WithFilter(randString())
	r.WithFriendsOnly(randInt())
	r.WithExtended(randBool())
	r.WithOffset(randInt())
	r.WithCount(randInt())
	r.WithSkipOwn(randBool())
}

func TestVKLikesGetListSuccess(t *testing.T) {
	params := requests.NewLikesGetListRequest()
	fillRandomlyLikesGetListRequest(&params)
	params.WithExtended(false)
	var expected models.LikesGetListResponse
	fillRandomlyLikesGetListResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "likes.getList", params.Params(), expectedJSON))
	resp, err := vk.Api.Likes.LikesGetList(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKLikesGetListExtendedSuccess(t *testing.T) {
	params := requests.NewLikesGetListRequest()
	fillRandomlyLikesGetListRequest(&params)
	params.WithExtended(true)
	var expected models.LikesGetListExtendedResponse
	fillRandomlyLikesGetListExtendedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "likes.getList", params.Params(), expectedJSON))
	resp, err := vk.Api.Likes.LikesGetListExtended(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyLikesIsLikedRequest(r *requests.LikesIsLikedRequest) {
	r.WithUserId(randInt())
	Type := new(models.LikesType)
	fillRandomlyLikesType(Type)
	r.WithType(*Type)
	r.WithOwnerId(randInt())
	r.WithItemId(randInt())
}

func TestVKLikesIsLikedSuccess(t *testing.T) {
	params := requests.NewLikesIsLikedRequest()
	fillRandomlyLikesIsLikedRequest(&params)
	var expected models.LikesIsLikedResponse
	fillRandomlyLikesIsLikedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "likes.isLiked", params.Params(), expectedJSON))
	resp, err := vk.Api.Likes.LikesIsLiked(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
