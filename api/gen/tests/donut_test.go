// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/donut"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyDonutGetFriendsRequest(r *requests.DonutGetFriendsRequest) {
	r.WithOwnerId(randInt())
	r.WithOffset(randInt())
	r.WithCount(randInt())
	lFields := randIntn(maxArrayLength + 1)
	r.WithFields(randStringArr(lFields))
}

func TestVKDonutGetFriendsSuccess(t *testing.T) {
	params := requests.NewDonutGetFriendsRequest()
	fillRandomlyDonutGetFriendsRequest(&params)
	var expected models.GroupsGetMembersFieldsResponse
	fillRandomlyGroupsGetMembersFieldsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "donut.getFriends", params.Params(), expectedJSON))
	resp, err := vk.Api.Donut.DonutGetFriends(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDonutGetSubscriptionRequest(r *requests.DonutGetSubscriptionRequest) {
	r.WithOwnerId(randInt())
}

func TestVKDonutGetSubscriptionSuccess(t *testing.T) {
	params := requests.NewDonutGetSubscriptionRequest()
	fillRandomlyDonutGetSubscriptionRequest(&params)
	var expected models.DonutGetSubscriptionResponse
	fillRandomlyDonutGetSubscriptionResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "donut.getSubscription", params.Params(), expectedJSON))
	resp, err := vk.Api.Donut.DonutGetSubscription(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDonutGetSubscriptionsRequest(r *requests.DonutGetSubscriptionsRequest) {
	Fields := new([]models.BaseUserGroupFields)
	lFields := randIntn(maxArrayLength + 1)
	*Fields = make([]models.BaseUserGroupFields, lFields)
	for i0 := 0; i0 < lFields; i0++ {
		fillRandomlyBaseUserGroupFields(&(*Fields)[i0])
	}
	r.WithFields(*Fields)
	r.WithOffset(randInt())
	r.WithCount(randInt())
}

func TestVKDonutGetSubscriptionsSuccess(t *testing.T) {
	params := requests.NewDonutGetSubscriptionsRequest()
	fillRandomlyDonutGetSubscriptionsRequest(&params)
	var expected models.DonutGetSubscriptionsResponse
	fillRandomlyDonutGetSubscriptionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "donut.getSubscriptions", params.Params(), expectedJSON))
	resp, err := vk.Api.Donut.DonutGetSubscriptions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyDonutIsDonRequest(r *requests.DonutIsDonRequest) {
	r.WithOwnerId(randInt())
}

func TestVKDonutIsDonSuccess(t *testing.T) {
	params := requests.NewDonutIsDonRequest()
	fillRandomlyDonutIsDonRequest(&params)
	var expected models.BaseBoolResponse
	fillRandomlyBaseBoolResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "donut.isDon", params.Params(), expectedJSON))
	resp, err := vk.Api.Donut.DonutIsDon(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

