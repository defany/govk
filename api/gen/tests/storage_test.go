// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyStorageGetRequest(r *requests.StorageGetRequest) {
	r.WithKey(randString())
	lKeys := randIntn(maxArrayLength + 1)
	r.WithKeys(randStringArr(lKeys))
	r.WithUserId(randInt())
}

func TestVKStorageGetSuccess(t *testing.T) {
	params := requests.NewStorageGetRequest()
	fillRandomlyStorageGetRequest(&params)
	var expected models.StorageGetResponse
	fillRandomlyStorageGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "storage.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Storage.StorageGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStorageGetKeysRequest(r *requests.StorageGetKeysRequest) {
	r.WithUserId(randInt())
	r.WithOffset(randInt())
	r.WithCount(randInt())
}

func TestVKStorageGetKeysSuccess(t *testing.T) {
	params := requests.NewStorageGetKeysRequest()
	fillRandomlyStorageGetKeysRequest(&params)
	var expected models.StorageGetKeysResponse
	fillRandomlyStorageGetKeysResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "storage.getKeys", params.Params(), expectedJSON))
	resp, err := vk.Api.Storage.StorageGetKeys(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStorageSetRequest(r *requests.StorageSetRequest) {
	r.WithKey(randString())
	r.WithValue(randString())
	r.WithUserId(randInt())
}

func TestVKStorageSetSuccess(t *testing.T) {
	params := requests.NewStorageSetRequest()
	fillRandomlyStorageSetRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "storage.set", params.Params(), expectedJSON))
	resp, err := vk.Api.Storage.StorageSet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

