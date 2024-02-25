package groupmodel

import "github.com/defany/govk/api/types"

type Group struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	ScreenName string        `json:"screen_name"`
	IsClosed   types.ApiBool `json:"is_closed"`
	Type       string        `json:"type"`
	Photo50    string        `json:"photo_50"`
	Photo100   string        `json:"photo_100"`
	Photo200   string        `json:"photo_200"`
}
