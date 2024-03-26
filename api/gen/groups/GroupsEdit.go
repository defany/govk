// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// GroupsEdit Edits a community.
type GroupsEditRequest api.Params

func NewGroupsEditRequest() GroupsEditRequest {
	params := make(GroupsEditRequest, 46)
	return params
}

func (g GroupsEditRequest) WithGroupId(g_group_id int) GroupsEditRequest{
	g["group_id"] = g_group_id
	return g
}

func (g GroupsEditRequest) WithTitle(g_title string) GroupsEditRequest{
	g["title"] = g_title
	return g
}

func (g GroupsEditRequest) WithDescription(g_description string) GroupsEditRequest{
	g["description"] = g_description
	return g
}

func (g GroupsEditRequest) WithScreenName(g_screen_name string) GroupsEditRequest{
	g["screen_name"] = g_screen_name
	return g
}

func (g GroupsEditRequest) WithAccess(g_access models.GroupsGroupAccess) GroupsEditRequest{
	g["access"] = g_access
	return g
}

func (g GroupsEditRequest) WithWebsite(g_website string) GroupsEditRequest{
	g["website"] = g_website
	return g
}

func (g GroupsEditRequest) WithSubject(g_subject models.GroupsGroupSubject) GroupsEditRequest{
	g["subject"] = g_subject
	return g
}

func (g GroupsEditRequest) WithEmail(g_email string) GroupsEditRequest{
	g["email"] = g_email
	return g
}

func (g GroupsEditRequest) WithPhone(g_phone string) GroupsEditRequest{
	g["phone"] = g_phone
	return g
}

func (g GroupsEditRequest) WithRss(g_rss string) GroupsEditRequest{
	g["rss"] = g_rss
	return g
}

func (g GroupsEditRequest) WithEventStartDate(g_event_start_date int) GroupsEditRequest{
	g["event_start_date"] = g_event_start_date
	return g
}

func (g GroupsEditRequest) WithEventFinishDate(g_event_finish_date int) GroupsEditRequest{
	g["event_finish_date"] = g_event_finish_date
	return g
}

func (g GroupsEditRequest) WithEventGroupId(g_event_group_id int) GroupsEditRequest{
	g["event_group_id"] = g_event_group_id
	return g
}

func (g GroupsEditRequest) WithPublicCategory(g_public_category int) GroupsEditRequest{
	g["public_category"] = g_public_category
	return g
}

func (g GroupsEditRequest) WithPublicSubcategory(g_public_subcategory int) GroupsEditRequest{
	g["public_subcategory"] = g_public_subcategory
	return g
}

func (g GroupsEditRequest) WithPublicDate(g_public_date string) GroupsEditRequest{
	g["public_date"] = g_public_date
	return g
}

func (g GroupsEditRequest) WithWall(g_wall models.GroupsGroupWall) GroupsEditRequest{
	g["wall"] = g_wall
	return g
}

func (g GroupsEditRequest) WithTopics(g_topics models.GroupsGroupTopics) GroupsEditRequest{
	g["topics"] = g_topics
	return g
}

func (g GroupsEditRequest) WithPhotos(g_photos models.GroupsGroupPhotos) GroupsEditRequest{
	g["photos"] = g_photos
	return g
}

func (g GroupsEditRequest) WithVideo(g_video models.GroupsGroupVideo) GroupsEditRequest{
	g["video"] = g_video
	return g
}

func (g GroupsEditRequest) WithAudio(g_audio models.GroupsGroupAudio) GroupsEditRequest{
	g["audio"] = g_audio
	return g
}

func (g GroupsEditRequest) WithLinks(g_links bool) GroupsEditRequest{
	g["links"] = g_links
	return g
}

func (g GroupsEditRequest) WithEvents(g_events bool) GroupsEditRequest{
	g["events"] = g_events
	return g
}

func (g GroupsEditRequest) WithPlaces(g_places bool) GroupsEditRequest{
	g["places"] = g_places
	return g
}

func (g GroupsEditRequest) WithContacts(g_contacts bool) GroupsEditRequest{
	g["contacts"] = g_contacts
	return g
}

func (g GroupsEditRequest) WithDocs(g_docs models.GroupsGroupDocs) GroupsEditRequest{
	g["docs"] = g_docs
	return g
}

func (g GroupsEditRequest) WithWiki(g_wiki models.GroupsGroupWiki) GroupsEditRequest{
	g["wiki"] = g_wiki
	return g
}

func (g GroupsEditRequest) WithMessages(g_messages bool) GroupsEditRequest{
	g["messages"] = g_messages
	return g
}

func (g GroupsEditRequest) WithArticles(g_articles bool) GroupsEditRequest{
	g["articles"] = g_articles
	return g
}

func (g GroupsEditRequest) WithAddresses(g_addresses bool) GroupsEditRequest{
	g["addresses"] = g_addresses
	return g
}

func (g GroupsEditRequest) WithAgeLimits(g_age_limits models.GroupsGroupAgeLimits) GroupsEditRequest{
	g["age_limits"] = g_age_limits
	return g
}

func (g GroupsEditRequest) WithMarket(g_market bool) GroupsEditRequest{
	g["market"] = g_market
	return g
}

func (g GroupsEditRequest) WithMarketComments(g_market_comments bool) GroupsEditRequest{
	g["market_comments"] = g_market_comments
	return g
}

func (g GroupsEditRequest) WithMarketCountry(g_market_country []int) GroupsEditRequest{
	g["market_country"] = g_market_country
	return g
}

func (g GroupsEditRequest) WithMarketCity(g_market_city []int) GroupsEditRequest{
	g["market_city"] = g_market_city
	return g
}

func (g GroupsEditRequest) WithMarketCurrency(g_market_currency models.GroupsGroupMarketCurrency) GroupsEditRequest{
	g["market_currency"] = g_market_currency
	return g
}

func (g GroupsEditRequest) WithMarketContact(g_market_contact int) GroupsEditRequest{
	g["market_contact"] = g_market_contact
	return g
}

func (g GroupsEditRequest) WithMarketWiki(g_market_wiki int) GroupsEditRequest{
	g["market_wiki"] = g_market_wiki
	return g
}

func (g GroupsEditRequest) WithObsceneFilter(g_obscene_filter bool) GroupsEditRequest{
	g["obscene_filter"] = g_obscene_filter
	return g
}

func (g GroupsEditRequest) WithObsceneStopwords(g_obscene_stopwords bool) GroupsEditRequest{
	g["obscene_stopwords"] = g_obscene_stopwords
	return g
}

func (g GroupsEditRequest) WithObsceneWords(g_obscene_words []string) GroupsEditRequest{
	g["obscene_words"] = g_obscene_words
	return g
}

func (g GroupsEditRequest) WithMainSection(g_main_section int) GroupsEditRequest{
	g["main_section"] = g_main_section
	return g
}

func (g GroupsEditRequest) WithSecondarySection(g_secondary_section int) GroupsEditRequest{
	g["secondary_section"] = g_secondary_section
	return g
}

func (g GroupsEditRequest) WithCountry(g_country int) GroupsEditRequest{
	g["country"] = g_country
	return g
}

func (g GroupsEditRequest) WithCity(g_city int) GroupsEditRequest{
	g["city"] = g_city
	return g
}

func (g GroupsEditRequest) Params() api.Params {
	return api.Params(g)
}

// May execute with listed access token types:
//    [ user, group ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_InvalidAddress ]
//
// https://dev.vk.com/method/groups.edit
func (g *Groups) GroupsEdit(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](g.api)

	res, err := req.Execute("groups.edit", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

