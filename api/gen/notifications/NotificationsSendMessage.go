// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// NotificationsSendMessage ...
type NotificationsSendMessageRequest api.Params

func NewNotificationsSendMessageRequest() NotificationsSendMessageRequest {
	params := make(NotificationsSendMessageRequest, 7)
	return params
}

func (n NotificationsSendMessageRequest) WithUserIds(n_user_ids []int) NotificationsSendMessageRequest{
	n["user_ids"] = n_user_ids
	return n
}

func (n NotificationsSendMessageRequest) WithMessage(n_message string) NotificationsSendMessageRequest{
	n["message"] = n_message
	return n
}

func (n NotificationsSendMessageRequest) WithFragment(n_fragment string) NotificationsSendMessageRequest{
	n["fragment"] = n_fragment
	return n
}

func (n NotificationsSendMessageRequest) WithGroupId(n_group_id int) NotificationsSendMessageRequest{
	n["group_id"] = n_group_id
	return n
}

func (n NotificationsSendMessageRequest) WithRandomId(n_random_id int) NotificationsSendMessageRequest{
	n["random_id"] = n_random_id
	return n
}

func (n NotificationsSendMessageRequest) WithSendingMode(n_sending_mode string) NotificationsSendMessageRequest{
	n["sending_mode"] = n_sending_mode
	return n
}

func (n NotificationsSendMessageRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//    [ service ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_GroupAppIsNotInstalledInCommunity ]
//
// https://dev.vk.com/method/notifications.sendMessage
func (n *Notifications) NotificationsSendMessage(params ...api.MethodParams) (resp models.NotificationsSendMessageResponse, err error) {
	req := api.NewRequest[models.NotificationsSendMessageResponse](n.api)

	res, err := req.Execute("notifications.sendMessage", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
