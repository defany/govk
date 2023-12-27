package model

import "github.com/goccy/go-json"

type Updates struct {
	Type    string          `json:"type"`
	EventID string          `json:"event_id"`
	Version string          `json:"v"`
	Object  json.RawMessage `json:"object"`
	GroupID int64           `json:"group_id"`
}

type Response struct {
	Ts      string    `json:"ts"`
	Updates []Updates `json:"updates"`
	Failed  int       `json:"failed"`
}
