package model

import "github.com/defany/govk/api/types"

type User struct {
	ID                     int           `json:"id"`
	FirstName              string        `json:"first_name"`
	LastName               string        `json:"last_name"`
	FirstNameNom           string        `json:"first_name_nom"`
	FirstNameGen           string        `json:"first_name_gen"`
	FirstNameDat           string        `json:"first_name_dat"`
	FirstNameAcc           string        `json:"first_name_acc"`
	FirstNameIns           string        `json:"first_name_ins"`
	FirstNameAbl           string        `json:"first_name_abl"`
	LastNameNom            string        `json:"last_name_nom"`
	LastNameGen            string        `json:"last_name_gen"`
	LastNameDat            string        `json:"last_name_dat"`
	LastNameAcc            string        `json:"last_name_acc"`
	LastNameIns            string        `json:"last_name_ins"`
	LastNameAbl            string        `json:"last_name_abl"`
	MaidenName             string        `json:"maiden_name"`
	Sex                    int           `json:"sex"`
	Nickname               string        `json:"nickname"`
	Domain                 string        `json:"domain"`
	ScreenName             string        `json:"screen_name"`
	Bdate                  string        `json:"bdate"`
	Photo50                string        `json:"photo_50"`
	Photo100               string        `json:"photo_100"`
	Photo200               string        `json:"photo_200"`
	PhotoMax               string        `json:"photo_max"`
	Photo200Orig           string        `json:"photo_200_orig"`
	Photo400Orig           string        `json:"photo_400_orig"`
	PhotoMaxOrig           string        `json:"photo_max_orig"`
	PhotoID                string        `json:"photo_id"`
	FriendStatus           int           `json:"friend_status"` // see FriendStatus const
	OnlineApp              int           `json:"online_app"`
	Online                 types.ApiBool `json:"online"`
	OnlineMobile           types.ApiBool `json:"online_mobile"`
	HasPhoto               types.ApiBool `json:"has_photo"`
	HasMobile              types.ApiBool `json:"has_mobile"`
	IsClosed               types.ApiBool `json:"is_closed"`
	IsFriend               types.ApiBool `json:"is_friend"`
	IsFavorite             types.ApiBool `json:"is_favorite"`
	IsHiddenFromFeed       types.ApiBool `json:"is_hidden_from_feed"`
	CanAccessClosed        types.ApiBool `json:"can_access_closed"`
	CanBeInvitedGroup      types.ApiBool `json:"can_be_invited_group"`
	CanPost                types.ApiBool `json:"can_post"`
	CanSeeAllPosts         types.ApiBool `json:"can_see_all_posts"`
	CanSeeAudio            types.ApiBool `json:"can_see_audio"`
	CanWritePrivateMessage types.ApiBool `json:"can_write_private_message"`
	CanSendFriendRequest   types.ApiBool `json:"can_send_friend_request"`
	CanCallFromGroup       types.ApiBool `json:"can_call_from_group"`
	Verified               types.ApiBool `json:"verified"`
	Trending               types.ApiBool `json:"trending"`
	Blacklisted            types.ApiBool `json:"blacklisted"`
	BlacklistedByMe        types.ApiBool `json:"blacklisted_by_me"`
	// Deprecated: Facebook и Instagram запрещены в России, Meta признана экстремистской организацией...
	Facebook string `json:"facebook"`
	// Deprecated: Facebook и Instagram запрещены в России, Meta признана экстремистской организацией...
	FacebookName string `json:"facebook_name"`
	// Deprecated: Facebook и Instagram запрещены в России, Meta признана экстремистской организацией...
	Instagram       string `json:"instagram"`
	Twitter         string `json:"twitter"`
	Site            string `json:"site"`
	Status          string `json:"status"`
	FollowersCount  int    `json:"followers_count"`
	CommonCount     int    `json:"common_count"`
	University      int    `json:"university"`
	UniversityName  string `json:"university_name"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Graduation      int    `json:"graduation"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
	HomeTown        string `json:"home_town"`
	Relation        int    `json:"relation"`
	Interests       string `json:"interests"`
	Music           string `json:"music"`
	Activities      string `json:"activities"`
	Movies          string `json:"movies"`
	Tv              string `json:"tv"`
	Books           string `json:"books"`
	Games           string `json:"games"`
	About           string `json:"about"`
	Quotes          string `json:"quotes"`
	Lists           []int  `json:"lists"`
	Deactivated     string `json:"deactivated"`
	WallDefault     string `json:"wall_default"`
	Timezone        int    `json:"timezone"`
	MobilePhone     string `json:"mobile_phone"`
	HomePhone       string `json:"home_phone"`
	TrackCode       string `json:"track_code"`
	Type            string `json:"type"`
	Skype           string `json:"skype"`
}
