package govk

import (
	"github.com/defany/govk/api"
	apiModel "github.com/defany/govk/api/apiModel"
	accountReqs "github.com/defany/govk/api/gen/account"
	adsReqs "github.com/defany/govk/api/gen/ads"
	adswebReqs "github.com/defany/govk/api/gen/adsweb"
	appsReqs "github.com/defany/govk/api/gen/apps"
	appwidgetsReqs "github.com/defany/govk/api/gen/appwidgets"
	authReqs "github.com/defany/govk/api/gen/auth"
	boardReqs "github.com/defany/govk/api/gen/board"
	databaseReqs "github.com/defany/govk/api/gen/database"
	docsReqs "github.com/defany/govk/api/gen/docs"
	donutReqs "github.com/defany/govk/api/gen/donut"
	downloadedgamesReqs "github.com/defany/govk/api/gen/downloadedgames"
	faveReqs "github.com/defany/govk/api/gen/fave"
	friendsReqs "github.com/defany/govk/api/gen/friends"
	giftsReqs "github.com/defany/govk/api/gen/gifts"
	groupsReqs "github.com/defany/govk/api/gen/groups"
	leadformsReqs "github.com/defany/govk/api/gen/leadforms"
	likesReqs "github.com/defany/govk/api/gen/likes"
	marketReqs "github.com/defany/govk/api/gen/market"
	messagesReqs "github.com/defany/govk/api/gen/messages"
	newsfeedReqs "github.com/defany/govk/api/gen/newsfeed"
	notesReqs "github.com/defany/govk/api/gen/notes"
	notificationsReqs "github.com/defany/govk/api/gen/notifications"
	ordersReqs "github.com/defany/govk/api/gen/orders"
	pagesReqs "github.com/defany/govk/api/gen/pages"
	photosReqs "github.com/defany/govk/api/gen/photos"
	podcastsReqs "github.com/defany/govk/api/gen/podcasts"
	pollsReqs "github.com/defany/govk/api/gen/polls"
	prettycardsReqs "github.com/defany/govk/api/gen/prettycards"
	searchReqs "github.com/defany/govk/api/gen/search"
	secureReqs "github.com/defany/govk/api/gen/secure"
	statsReqs "github.com/defany/govk/api/gen/stats"
	statusReqs "github.com/defany/govk/api/gen/status"
	storageReqs "github.com/defany/govk/api/gen/storage"
	storeReqs "github.com/defany/govk/api/gen/store"
	storiesReqs "github.com/defany/govk/api/gen/stories"
	streamingReqs "github.com/defany/govk/api/gen/streaming"
	usersReqs "github.com/defany/govk/api/gen/users"
	utilsReqs "github.com/defany/govk/api/gen/utils"
	videoReqs "github.com/defany/govk/api/gen/video"
	wallReqs "github.com/defany/govk/api/gen/wall"
	widgetsReqs "github.com/defany/govk/api/gen/widgets"
	"github.com/defany/govk/updates"
)

type VK struct {
	Api     *apiModel.ApiProvider
	Updates *updates.Updates
}

func NewVK(tokens ...string) (*VK, error) {
	api, err := api.NewAPI(tokens...)
	if err != nil {
		return nil, err
	}
	apiProvider := &apiModel.ApiProvider{
		Api:             api,
		Account:         accountReqs.NewAccount(api),
		Ads:             adsReqs.NewAds(api),
		Adsweb:          adswebReqs.NewAdsweb(api),
		AppWidgets:      appwidgetsReqs.NewAppWidgets(api),
		Apps:            appsReqs.NewApps(api),
		Auth:            authReqs.NewAuth(api),
		Board:           boardReqs.NewBoard(api),
		Database:        databaseReqs.NewDatabase(api),
		Docs:            docsReqs.NewDocs(api),
		Donut:           donutReqs.NewDonut(api),
		DownloadedGames: downloadedgamesReqs.NewDownloadedGames(api),
		Fave:            faveReqs.NewFave(api),
		Friends:         friendsReqs.NewFriends(api),
		Gifts:           giftsReqs.NewGifts(api),
		Groups:          groupsReqs.NewGroups(api),
		LeadForms:       leadformsReqs.NewLeadForms(api),
		Likes:           likesReqs.NewLikes(api),
		Market:          marketReqs.NewMarket(api),
		Messages:        messagesReqs.NewMessages(api),
		Newsfeed:        newsfeedReqs.NewNewsfeed(api),
		Notes:           notesReqs.NewNotes(api),
		Notifications:   notificationsReqs.NewNotifications(api),
		Orders:          ordersReqs.NewOrders(api),
		Pages:           pagesReqs.NewPages(api),
		Photos:          photosReqs.NewPhotos(api),
		Podcasts:        podcastsReqs.NewPodcasts(api),
		Polls:           pollsReqs.NewPolls(api),
		PrettyCards:     prettycardsReqs.NewPrettyCards(api),
		Search:          searchReqs.NewSearch(api),
		Secure:          secureReqs.NewSecure(api),
		Stats:           statsReqs.NewStats(api),
		Status:          statusReqs.NewStatus(api),
		Storage:         storageReqs.NewStorage(api),
		Store:           storeReqs.NewStore(api),
		Stories:         storiesReqs.NewStories(api),
		Streaming:       streamingReqs.NewStreaming(api),
		Users:           usersReqs.NewUsers(api),
		Utils:           utilsReqs.NewUtils(api),
		Video:           videoReqs.NewVideo(api),
		Wall:            wallReqs.NewWall(api),
		Widgets:         widgetsReqs.NewWidgets(api),
	}

	updates := updates.NewUpdates(apiProvider)

	return &VK{
		Api:     apiProvider,
		Updates: updates,
	}, nil
}
