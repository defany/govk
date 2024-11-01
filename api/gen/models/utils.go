// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type UtilsDomainResolved struct {
	// Group ID
	//  Format: int64
	GroupId *int `json:"group_id,omitempty"`
	// Object ID
	ObjectId *int `json:"object_id,omitempty"`
	Type *UtilsDomainResolvedType `json:"type,omitempty"`
}

// UtilsDomainResolvedType Object type
type UtilsDomainResolvedType string

const (
	UtilsDomainResolvedTypeUser UtilsDomainResolvedType = "user"
	UtilsDomainResolvedTypeGroup UtilsDomainResolvedType = "group"
	UtilsDomainResolvedTypeApplication UtilsDomainResolvedType = "application"
	UtilsDomainResolvedTypePage UtilsDomainResolvedType = "page"
	UtilsDomainResolvedTypeVkApp UtilsDomainResolvedType = "vk_app"
	UtilsDomainResolvedTypeCommunityApplication UtilsDomainResolvedType = "community_application"
)

type UtilsLastShortenedLink struct {
	// Access key for private stats
	AccessKey *string `json:"access_key,omitempty"`
	// Link key (characters after vk.cc/)
	Key *string `json:"key,omitempty"`
	// Short link URL
	//  Format: uri
	ShortUrl *string `json:"short_url,omitempty"`
	// Creation time in Unixtime
	Timestamp *int `json:"timestamp,omitempty"`
	// Full URL
	//  Format: uri
	Url *string `json:"url,omitempty"`
	// Total views number
	Views *int `json:"views,omitempty"`
}

type UtilsLinkChecked struct {
	// Link URL
	//  Format: uri
	Link *string `json:"link,omitempty"`
	Status *UtilsLinkCheckedStatus `json:"status,omitempty"`
}

// UtilsLinkCheckedStatus Link status
type UtilsLinkCheckedStatus string

const (
	UtilsLinkCheckedStatusNotBanned UtilsLinkCheckedStatus = "not_banned"
	UtilsLinkCheckedStatusBanned UtilsLinkCheckedStatus = "banned"
	UtilsLinkCheckedStatusProcessing UtilsLinkCheckedStatus = "processing"
)

type UtilsLinkStats struct {
	// Link key (characters after vk.cc/)
	Key *string `json:"key,omitempty"`
	Stats *[]UtilsStats `json:"stats,omitempty"`
}

type UtilsLinkStatsExtended struct {
	// Link key (characters after vk.cc/)
	Key *string `json:"key,omitempty"`
	Stats *[]UtilsStatsExtended `json:"stats,omitempty"`
}

type UtilsShortLink struct {
	// Access key for private stats
	AccessKey *string `json:"access_key,omitempty"`
	// Link key (characters after vk.cc/)
	Key *string `json:"key,omitempty"`
	// Short link URL
	//  Format: uri
	ShortUrl *string `json:"short_url,omitempty"`
	// Full URL
	//  Format: uri
	Url *string `json:"url,omitempty"`
}

type UtilsStats struct {
	// Start time
	Timestamp *int `json:"timestamp,omitempty"`
	// Total views number
	Views *int `json:"views,omitempty"`
}

type UtilsStatsCity struct {
	// City ID
	CityId *int `json:"city_id,omitempty"`
	// Views number
	Views *int `json:"views,omitempty"`
}

type UtilsStatsCountry struct {
	// Country ID
	CountryId *int `json:"country_id,omitempty"`
	// Views number
	Views *int `json:"views,omitempty"`
}

type UtilsStatsExtended struct {
	Cities *[]UtilsStatsCity `json:"cities,omitempty"`
	Countries *[]UtilsStatsCountry `json:"countries,omitempty"`
	SexAge *[]UtilsStatsSexAge `json:"sex_age,omitempty"`
	// Start time
	Timestamp *int `json:"timestamp,omitempty"`
	// Total views number
	Views *int `json:"views,omitempty"`
}

type UtilsStatsSexAge struct {
	// Age denotation
	AgeRange *string `json:"age_range,omitempty"`
	//  Views by female users
	Female *int `json:"female,omitempty"`
	//  Views by male users
	Male *int `json:"male,omitempty"`
}

type UtilsCheckLinkResponse struct {
	Response UtilsLinkChecked `json:"response"`
}

type UtilsGetLastShortenedLinksResponse struct {
	Utils struct {
		// Total number of available results
		//  Minimum: 0
		Count *int `json:"count,omitempty"`
		Items *[]UtilsLastShortenedLink `json:"items,omitempty"`
	} `json:"utils"`
}

type UtilsGetLinkStatsExtendedResponse struct {
	Response UtilsLinkStatsExtended `json:"response"`
}

type UtilsGetLinkStatsResponse struct {
	Response UtilsLinkStats `json:"response"`
}

type UtilsGetServerTimeResponse struct {
	// Time as Unixtime
	Response int `json:"response"`
}

type UtilsGetShortLinkResponse struct {
	Response UtilsShortLink `json:"response"`
}

type UtilsResolveScreenNameResponse struct {
	Response UtilsDomainResolved `json:"response"`
}

