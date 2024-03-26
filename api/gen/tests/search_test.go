// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlySearchGetHintsRequest(r *requests.SearchGetHintsRequest) {
	r.WithQ(randString())
	r.WithOffset(randInt())
	r.WithLimit(randInt())
	lFilters := randIntn(maxArrayLength + 1)
	r.WithFilters(randStringArr(lFilters))
	lFields := randIntn(maxArrayLength + 1)
	r.WithFields(randStringArr(lFields))
	r.WithSearchGlobal(randBool())
}

func TestVKSearchGetHintsSuccess(t *testing.T) {
	params := requests.NewSearchGetHintsRequest()
	fillRandomlySearchGetHintsRequest(&params)
	var expected models.SearchGetHintsResponse
	fillRandomlySearchGetHintsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "search.getHints", params.Params(), expectedJSON))
	resp, err := vk.Api.Search.SearchGetHints(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
