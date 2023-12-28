package model

type GetLongPollServerResponse struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     string `json:"ts"`
}
