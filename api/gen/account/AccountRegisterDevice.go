// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// AccountRegisterDevice Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
type AccountRegisterDeviceRequest api.Params

func NewAccountRegisterDeviceRequest() AccountRegisterDeviceRequest {
	params := make(AccountRegisterDeviceRequest, 8)
	return params
}

func (a AccountRegisterDeviceRequest) WithToken(a_token string) AccountRegisterDeviceRequest{
	a["token"] = a_token
	return a
}

func (a AccountRegisterDeviceRequest) WithDeviceModel(a_device_model string) AccountRegisterDeviceRequest{
	a["device_model"] = a_device_model
	return a
}

func (a AccountRegisterDeviceRequest) WithDeviceYear(a_device_year int) AccountRegisterDeviceRequest{
	a["device_year"] = a_device_year
	return a
}

func (a AccountRegisterDeviceRequest) WithDeviceId(a_device_id string) AccountRegisterDeviceRequest{
	a["device_id"] = a_device_id
	return a
}

func (a AccountRegisterDeviceRequest) WithSystemVersion(a_system_version string) AccountRegisterDeviceRequest{
	a["system_version"] = a_system_version
	return a
}

func (a AccountRegisterDeviceRequest) WithSettings(a_settings string) AccountRegisterDeviceRequest{
	a["settings"] = a_settings
	return a
}

func (a AccountRegisterDeviceRequest) WithSandbox(a_sandbox bool) AccountRegisterDeviceRequest{
	a["sandbox"] = a_sandbox
	return a
}

func (a AccountRegisterDeviceRequest) Params() api.Params {
	return api.Params(a)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/account.registerDevice
func (a *Account) AccountRegisterDevice(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](a.api)

	res, err := req.Execute("account.registerDevice", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

