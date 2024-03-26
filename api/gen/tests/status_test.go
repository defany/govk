// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyStatusGetRequest(r *requests.StatusGetRequest) {
	r.WithUserId(randInt())
	r.WithGroupId(randInt())
}

func TestVKStatusGetSuccess(t *testing.T) {
	params := requests.NewStatusGetRequest()
	fillRandomlyStatusGetRequest(&params)
	var expected models.StatusGetResponse
	fillRandomlyStatusGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "status.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Status.StatusGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStatusSetRequest(r *requests.StatusSetRequest) {
	r.WithText(randString())
	r.WithGroupId(randInt())
}

func TestVKStatusSetSuccess(t *testing.T) {
	params := requests.NewStatusSetRequest()
	fillRandomlyStatusSetRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "status.set", params.Params(), expectedJSON))
	resp, err := vk.Api.Status.StatusSet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
