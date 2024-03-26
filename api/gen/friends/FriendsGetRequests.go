// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// FriendsGetRequests Returns information about the current user's incoming and outgoing friend requests.
type FriendsGetRequestsRequest api.Params

func NewFriendsGetRequestsRequest() FriendsGetRequestsRequest {
	params := make(FriendsGetRequestsRequest, 12)
	return params
}

func (f FriendsGetRequestsRequest) WithOffset(f_offset int) FriendsGetRequestsRequest{
	f["offset"] = f_offset
	return f
}

func (f FriendsGetRequestsRequest) WithCount(f_count int) FriendsGetRequestsRequest{
	f["count"] = f_count
	return f
}

func (f FriendsGetRequestsRequest) WithExtended(f_extended bool) FriendsGetRequestsRequest{
	f["extended"] = f_extended
	return f
}

func (f FriendsGetRequestsRequest) WithNeedMutual(f_need_mutual bool) FriendsGetRequestsRequest{
	f["need_mutual"] = f_need_mutual
	return f
}

func (f FriendsGetRequestsRequest) WithOut(f_out bool) FriendsGetRequestsRequest{
	f["out"] = f_out
	return f
}

func (f FriendsGetRequestsRequest) WithSort(f_sort int) FriendsGetRequestsRequest{
	f["sort"] = f_sort
	return f
}

func (f FriendsGetRequestsRequest) WithNeedViewed(f_need_viewed bool) FriendsGetRequestsRequest{
	f["need_viewed"] = f_need_viewed
	return f
}

func (f FriendsGetRequestsRequest) WithSuggested(f_suggested bool) FriendsGetRequestsRequest{
	f["suggested"] = f_suggested
	return f
}

func (f FriendsGetRequestsRequest) WithRef(f_ref string) FriendsGetRequestsRequest{
	f["ref"] = f_ref
	return f
}

func (f FriendsGetRequestsRequest) WithFields(f_fields []models.UsersFields) FriendsGetRequestsRequest{
	f["fields"] = f_fields
	return f
}

func (f FriendsGetRequestsRequest) Params() api.Params {
	return api.Params(f)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/friends.getRequests
func (f *Friends) FriendsGetRequests(params ...api.MethodParams) (resp models.FriendsGetRequestsResponse, err error) {
	req := api.NewRequest[models.FriendsGetRequestsResponse](f.api)

	res, err := req.Execute("friends.getRequests", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

// FriendsGetRequestsExtended Returns information about the current user's incoming and outgoing friend requests.
// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global API errors.
//
// https://dev.vk.com/method/friends.getRequests
func (f *Friends) FriendsGetRequestsExtended(params ...api.MethodParams) (resp models.FriendsGetRequestsExtendedResponse, err error) {
	req := api.NewRequest[models.FriendsGetRequestsExtendedResponse](f.api)

	res, err := req.Execute("friends.getRequests", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

