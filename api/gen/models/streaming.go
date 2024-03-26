// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type StreamingGetServerUrlResponse struct {
	Streaming struct {
		// Server host
		Endpoint *string `json:"endpoint,omitempty"`
		// Access key
		Key *string `json:"key,omitempty"`
	} `json:"streaming"`
}

