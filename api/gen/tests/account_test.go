// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyAccountBanRequest(r *requests.AccountBanRequest) {
	r.WithOwnerId(randInt())
}

func TestVKAccountBanSuccess(t *testing.T) {
	params := requests.NewAccountBanRequest()
	fillRandomlyAccountBanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.ban", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountBan(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountChangePasswordRequest(r *requests.AccountChangePasswordRequest) {
	r.WithRestoreSid(randString())
	r.WithChangePasswordHash(randString())
	r.WithOldPassword(randString())
	r.WithNewPassword(randString())
}

func TestVKAccountChangePasswordSuccess(t *testing.T) {
	params := requests.NewAccountChangePasswordRequest()
	fillRandomlyAccountChangePasswordRequest(&params)
	var expected models.AccountChangePasswordResponse
	fillRandomlyAccountChangePasswordResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.changePassword", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountChangePassword(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetActiveOffersRequest(r *requests.AccountGetActiveOffersRequest) {
	r.WithOffset(randInt())
	r.WithCount(randInt())
}

func TestVKAccountGetActiveOffersSuccess(t *testing.T) {
	params := requests.NewAccountGetActiveOffersRequest()
	fillRandomlyAccountGetActiveOffersRequest(&params)
	var expected models.AccountGetActiveOffersResponse
	fillRandomlyAccountGetActiveOffersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getActiveOffers", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetActiveOffers(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetAppPermissionsRequest(r *requests.AccountGetAppPermissionsRequest) {
	r.WithUserId(randInt())
}

func TestVKAccountGetAppPermissionsSuccess(t *testing.T) {
	params := requests.NewAccountGetAppPermissionsRequest()
	fillRandomlyAccountGetAppPermissionsRequest(&params)
	var expected models.AccountGetAppPermissionsResponse
	fillRandomlyAccountGetAppPermissionsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getAppPermissions", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetAppPermissions(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetBannedRequest(r *requests.AccountGetBannedRequest) {
	r.WithOffset(randInt())
	r.WithCount(randInt())
}

func TestVKAccountGetBannedSuccess(t *testing.T) {
	params := requests.NewAccountGetBannedRequest()
	fillRandomlyAccountGetBannedRequest(&params)
	var expected models.AccountGetBannedResponse
	fillRandomlyAccountGetBannedResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getBanned", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetBanned(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetCountersRequest(r *requests.AccountGetCountersRequest) {
	lFilter := randIntn(maxArrayLength + 1)
	r.WithFilter(randStringArr(lFilter))
	r.WithUserId(randInt())
}

func TestVKAccountGetCountersSuccess(t *testing.T) {
	params := requests.NewAccountGetCountersRequest()
	fillRandomlyAccountGetCountersRequest(&params)
	var expected models.AccountGetCountersResponse
	fillRandomlyAccountGetCountersResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getCounters", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetCounters(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetInfoRequest(r *requests.AccountGetInfoRequest) {
	lFields := randIntn(maxArrayLength + 1)
	r.WithFields(randStringArr(lFields))
}

func TestVKAccountGetInfoSuccess(t *testing.T) {
	params := requests.NewAccountGetInfoRequest()
	fillRandomlyAccountGetInfoRequest(&params)
	var expected models.AccountGetInfoResponse
	fillRandomlyAccountGetInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAccountGetProfileInfoSuccess(t *testing.T) {
	var expected models.AccountGetProfileInfoResponse
	fillRandomlyAccountGetProfileInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getProfileInfo", nil, expectedJSON))
	resp, err := vk.Api.Account.AccountGetProfileInfo()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountGetPushSettingsRequest(r *requests.AccountGetPushSettingsRequest) {
	r.WithDeviceId(randString())
}

func TestVKAccountGetPushSettingsSuccess(t *testing.T) {
	params := requests.NewAccountGetPushSettingsRequest()
	fillRandomlyAccountGetPushSettingsRequest(&params)
	var expected models.AccountGetPushSettingsResponse
	fillRandomlyAccountGetPushSettingsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.getPushSettings", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountGetPushSettings(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountRegisterDeviceRequest(r *requests.AccountRegisterDeviceRequest) {
	r.WithToken(randString())
	r.WithDeviceModel(randString())
	r.WithDeviceYear(randInt())
	r.WithDeviceId(randString())
	r.WithSystemVersion(randString())
	r.WithSettings(randString())
	r.WithSandbox(randBool())
}

func TestVKAccountRegisterDeviceSuccess(t *testing.T) {
	params := requests.NewAccountRegisterDeviceRequest()
	fillRandomlyAccountRegisterDeviceRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.registerDevice", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountRegisterDevice(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSaveProfileInfoRequest(r *requests.AccountSaveProfileInfoRequest) {
	r.WithFirstName(randString())
	r.WithLastName(randString())
	r.WithMaidenName(randString())
	r.WithScreenName(randString())
	r.WithCancelRequestId(randInt())
	r.WithSex(randInt())
	r.WithRelation(randInt())
	r.WithRelationPartnerId(randInt())
	r.WithBdate(randString())
	r.WithBdateVisibility(randInt())
	r.WithHomeTown(randString())
	r.WithCountryId(randInt())
	r.WithCityId(randInt())
	r.WithStatus(randString())
}

func TestVKAccountSaveProfileInfoSuccess(t *testing.T) {
	params := requests.NewAccountSaveProfileInfoRequest()
	fillRandomlyAccountSaveProfileInfoRequest(&params)
	var expected models.AccountSaveProfileInfoResponse
	fillRandomlyAccountSaveProfileInfoResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.saveProfileInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSaveProfileInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetInfoRequest(r *requests.AccountSetInfoRequest) {
	r.WithName(randString())
	r.WithValue(randString())
}

func TestVKAccountSetInfoSuccess(t *testing.T) {
	params := requests.NewAccountSetInfoRequest()
	fillRandomlyAccountSetInfoRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setInfo", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetInfo(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func TestVKAccountSetOfflineSuccess(t *testing.T) {
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setOffline", nil, expectedJSON))
	resp, err := vk.Api.Account.AccountSetOffline()
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetOnlineRequest(r *requests.AccountSetOnlineRequest) {
	r.WithVoip(randBool())
}

func TestVKAccountSetOnlineSuccess(t *testing.T) {
	params := requests.NewAccountSetOnlineRequest()
	fillRandomlyAccountSetOnlineRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setOnline", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetOnline(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetPushSettingsRequest(r *requests.AccountSetPushSettingsRequest) {
	r.WithDeviceId(randString())
	r.WithSettings(randString())
	r.WithKey(randString())
	lValue := randIntn(maxArrayLength + 1)
	r.WithValue(randStringArr(lValue))
}

func TestVKAccountSetPushSettingsSuccess(t *testing.T) {
	params := requests.NewAccountSetPushSettingsRequest()
	fillRandomlyAccountSetPushSettingsRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setPushSettings", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetPushSettings(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountSetSilenceModeRequest(r *requests.AccountSetSilenceModeRequest) {
	r.WithDeviceId(randString())
	r.WithTime(randInt())
	r.WithPeerId(randInt())
	r.WithSound(randInt())
}

func TestVKAccountSetSilenceModeSuccess(t *testing.T) {
	params := requests.NewAccountSetSilenceModeRequest()
	fillRandomlyAccountSetSilenceModeRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.setSilenceMode", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountSetSilenceMode(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountUnbanRequest(r *requests.AccountUnbanRequest) {
	r.WithOwnerId(randInt())
}

func TestVKAccountUnbanSuccess(t *testing.T) {
	params := requests.NewAccountUnbanRequest()
	fillRandomlyAccountUnbanRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.unban", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountUnban(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyAccountUnregisterDeviceRequest(r *requests.AccountUnregisterDeviceRequest) {
	r.WithDeviceId(randString())
	r.WithSandbox(randBool())
}

func TestVKAccountUnregisterDeviceSuccess(t *testing.T) {
	params := requests.NewAccountUnregisterDeviceRequest()
	fillRandomlyAccountUnregisterDeviceRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "account.unregisterDevice", params.Params(), expectedJSON))
	resp, err := vk.Api.Account.AccountUnregisterDevice(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}
