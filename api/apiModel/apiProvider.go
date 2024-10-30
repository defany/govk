package apiModel

import (
	"github.com/defany/govk/api"
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
	"net/http"
)

type ApiProvider struct {
	Api             *api.API
	Account         *accountReqs.Account
	Ads             *adsReqs.Ads
	Adsweb          *adswebReqs.Adsweb
	AppWidgets      *appwidgetsReqs.AppWidgets
	Apps            *appsReqs.Apps
	Auth            *authReqs.Auth
	Board           *boardReqs.Board
	Database        *databaseReqs.Database
	Docs            *docsReqs.Docs
	Donut           *donutReqs.Donut
	DownloadedGames *downloadedgamesReqs.DownloadedGames
	Fave            *faveReqs.Fave
	Friends         *friendsReqs.Friends
	Gifts           *giftsReqs.Gifts
	Groups          *groupsReqs.Groups
	LeadForms       *leadformsReqs.LeadForms
	Likes           *likesReqs.Likes
	Market          *marketReqs.Market
	Messages        *messagesReqs.Messages
	Newsfeed        *newsfeedReqs.Newsfeed
	Notes           *notesReqs.Notes
	Notifications   *notificationsReqs.Notifications
	Orders          *ordersReqs.Orders
	Pages           *pagesReqs.Pages
	Photos          *photosReqs.Photos
	Podcasts        *podcastsReqs.Podcasts
	Polls           *pollsReqs.Polls
	PrettyCards     *prettycardsReqs.PrettyCards
	Search          *searchReqs.Search
	Secure          *secureReqs.Secure
	Stats           *statsReqs.Stats
	Status          *statusReqs.Status
	Storage         *storageReqs.Storage
	Store           *storeReqs.Store
	Stories         *storiesReqs.Stories
	Streaming       *streamingReqs.Streaming
	Users           *usersReqs.Users
	Utils           *utilsReqs.Utils
	Video           *videoReqs.Video
	Wall            *wallReqs.Wall
	Widgets         *widgetsReqs.Widgets
}

func (a *ApiProvider) WithLimit(limit uint) *ApiProvider {
	a.Api.WithLimit(limit)

	return a
}

func (a *ApiProvider) WithUrl(url string) *ApiProvider {
	a.Api.WithApiURL(url)

	return a
}

func (a *ApiProvider) WithVersion(v string) *ApiProvider {
	a.Api.WithVersion(v)

	return a
}

func (a *ApiProvider) WithMaxRetries(retries uint) *ApiProvider {
	a.Api.WithMaxRetries(retries)

	return a
}

func (a *ApiProvider) WithHTTP(client *http.Client) *ApiProvider {
	a.Api.WithHTTP(client)

	return a
}
