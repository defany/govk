// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type CommentThread struct {
	// Information whether current user can comment the post
	CanPost *bool `json:"can_post,omitempty"`
	// Comments number
	//  Minimum: 0
	Count int `json:"count"`
	// Information whether groups can comment the post
	GroupsCanPost *bool `json:"groups_can_post,omitempty"`
	Items *[]WallWallComment `json:"items,omitempty"`
	// Information whether recommended to display reply button
	ShowReplyButton *bool `json:"show_reply_button,omitempty"`
}

