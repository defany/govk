// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// LeadFormsUpdate ...
type LeadFormsUpdateRequest api.Params

func NewLeadFormsUpdateRequest() LeadFormsUpdateRequest {
	params := make(LeadFormsUpdateRequest, 16)
	return params
}

func (l LeadFormsUpdateRequest) WithGroupId(l_group_id int) LeadFormsUpdateRequest{
	l["group_id"] = l_group_id
	return l
}

func (l LeadFormsUpdateRequest) WithFormId(l_form_id int) LeadFormsUpdateRequest{
	l["form_id"] = l_form_id
	return l
}

func (l LeadFormsUpdateRequest) WithName(l_name string) LeadFormsUpdateRequest{
	l["name"] = l_name
	return l
}

func (l LeadFormsUpdateRequest) WithTitle(l_title string) LeadFormsUpdateRequest{
	l["title"] = l_title
	return l
}

func (l LeadFormsUpdateRequest) WithDescription(l_description string) LeadFormsUpdateRequest{
	l["description"] = l_description
	return l
}

func (l LeadFormsUpdateRequest) WithQuestions(l_questions string) LeadFormsUpdateRequest{
	l["questions"] = l_questions
	return l
}

func (l LeadFormsUpdateRequest) WithPolicyLinkUrl(l_policy_link_url string) LeadFormsUpdateRequest{
	l["policy_link_url"] = l_policy_link_url
	return l
}

func (l LeadFormsUpdateRequest) WithPhoto(l_photo string) LeadFormsUpdateRequest{
	l["photo"] = l_photo
	return l
}

func (l LeadFormsUpdateRequest) WithConfirmation(l_confirmation string) LeadFormsUpdateRequest{
	l["confirmation"] = l_confirmation
	return l
}

func (l LeadFormsUpdateRequest) WithSiteLinkUrl(l_site_link_url string) LeadFormsUpdateRequest{
	l["site_link_url"] = l_site_link_url
	return l
}

func (l LeadFormsUpdateRequest) WithActive(l_active bool) LeadFormsUpdateRequest{
	l["active"] = l_active
	return l
}

func (l LeadFormsUpdateRequest) WithOncePerUser(l_once_per_user bool) LeadFormsUpdateRequest{
	l["once_per_user"] = l_once_per_user
	return l
}

func (l LeadFormsUpdateRequest) WithPixelCode(l_pixel_code string) LeadFormsUpdateRequest{
	l["pixel_code"] = l_pixel_code
	return l
}

func (l LeadFormsUpdateRequest) WithNotifyAdmins(l_notify_admins []int) LeadFormsUpdateRequest{
	l["notify_admins"] = l_notify_admins
	return l
}

func (l LeadFormsUpdateRequest) WithNotifyEmails(l_notify_emails []string) LeadFormsUpdateRequest{
	l["notify_emails"] = l_notify_emails
	return l
}

func (l LeadFormsUpdateRequest) Params() api.Params {
	return api.Params(l)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_NotFound ]
//
// https://dev.vk.com/method/leadForms.update
func (l *LeadForms) LeadFormsUpdate(params ...api.MethodParams) (resp models.LeadFormsCreateResponse, err error) {
	req := api.NewRequest[models.LeadFormsCreateResponse](l.api)

	res, err := req.Execute("leadForms.update", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}
