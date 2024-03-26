// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/streaming"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func TestVKStreamingGetServerUrlSuccess(t *testing.T) {
	var expected models.StreamingGetServerUrlResponse
	fillRandomlyStreamingGetServerUrlResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "streaming.getServerUrl", nil, expectedJSON))
	resp, err := vk.Api.Streaming.StreamingGetServerUrl()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyStreamingSetSettingsRequest(r *requests.StreamingSetSettingsRequest) {
	r.WithMonthlyTier(randString())
}

func TestVKStreamingSetSettingsSuccess(t *testing.T) {
	params := requests.NewStreamingSetSettingsRequest()
	fillRandomlyStreamingSetSettingsRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "streaming.setSettings", params.Params(), expectedJSON))
	resp, err := vk.Api.Streaming.StreamingSetSettings(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
