package model

type Response struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
	Failed  int      `json:"failed"`
}
