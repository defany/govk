package model

import "github.com/defany/govk/api/types"

type ClientInfo struct {
	ButtonActions  []string      `json:"button_actions"`
	Keyboard       types.ApiBool `json:"keyboard"`
	InlineKeyboard types.ApiBool `json:"inline_keyboard"`
	Carousel       types.ApiBool `json:"carousel"`
	LangID         int           `json:"lang_id"`
}
