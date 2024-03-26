// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// SecureGetSMSHistory Shows a list of SMS notifications sent by the application using [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
type SecureGetSMSHistoryRequest api.Params

func NewSecureGetSMSHistoryRequest() SecureGetSMSHistoryRequest {
	params := make(SecureGetSMSHistoryRequest, 5)
	return params
}

func (s SecureGetSMSHistoryRequest) WithUserId(s_user_id int) SecureGetSMSHistoryRequest{
	s["user_id"] = s_user_id
	return s
}

func (s SecureGetSMSHistoryRequest) WithDateFrom(s_date_from int) SecureGetSMSHistoryRequest{
	s["date_from"] = s_date_from
	return s
}

func (s SecureGetSMSHistoryRequest) WithDateTo(s_date_to int) SecureGetSMSHistoryRequest{
	s["date_to"] = s_date_to
	return s
}

func (s SecureGetSMSHistoryRequest) WithLimit(s_limit int) SecureGetSMSHistoryRequest{
	s["limit"] = s_limit
	return s
}

func (s SecureGetSMSHistoryRequest) Params() api.Params {
	return api.Params(s)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/secure.getSMSHistory
func (s *Secure) SecureGetSMSHistory(params ...api.MethodParams) (resp models.SecureGetSMSHistoryResponse, err error) {
	req := api.NewRequest[models.SecureGetSMSHistoryResponse](s.api)

	res, err := req.Execute("secure.getSMSHistory", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
