package msgmodel

import (
	"github.com/defany/govk/api/types"
	"github.com/goccy/go-json"
)

// Button action type.
const (
	// ButtonText A button that sends a message with text specified in the label.
	ButtonText = "text"

	// ButtonVKPay Opens the VK Pay window with predefined parameters. The button is called
	// “Pay with VK Pay” (VK Pay is displayed as a logo). This button always
	// stretches to the whole keyboard width.
	ButtonVKPay = "vkpay"

	// ButtonVKApp Opens a specified VK Apps app. This button always stretches to the whole
	// keyboard width.
	ButtonVKApp = "open_app"

	// ButtonLocation Sends the location to the chat. This button always stretches to the
	// whole keyboard width.
	ButtonLocation = "location"

	// ButtonOpenLink Opens the specified link.
	ButtonOpenLink = "open_link"

	// ButtonCallback Allows, without sending a message from the user, to receive a
	// notification about pressing the button and perform the necessary action.
	ButtonCallback = "callback"
)

// Button color. This parameter is used only for buttons with the text and callback types.
const (
	Primary = "primary" // Blue button, indicates the main action. #5181B8
	ButtonBlue

	Secondary = "secondary" // Default white button. #FFFFFF
	ButtonWhite

	Negative = "negative" // Dangerous or negative action (cancel, delete etc.) #E64646
	ButtonRed

	Positive = "positive" // Accept, agree. #4BB34B
	ButtonGreen
)

// MessageKeyboardButtonAction struct.
type MessageKeyboardButtonAction struct {
	// Type of button action
	Type string `json:"type"`
	// ID of the called application
	// with the VK Mini Apps type (for the button with the "open_app" action)
	AppID int `json:"app_id,omitempty"`
	// ID of the community
	// where the application is installed (for the button with the "open_app" action)
	OwnerID int `json:"owner_id,omitempty"`
	// Additional data along with the sent message (for the developer)
	Payload string `json:"payload,omitempty"`
	// Button text
	Label string `json:"label,omitempty"`
	// Contains VK Pay payment parameters (to see more, open https://dev.vk.com/ru/pay/getting-started)
	// and the application ID
	Hash string `json:"hash,omitempty"`
	// The link that needs to be opened
	// by clicking on the button
	Link string `json:"link,omitempty"`
}

// MessageKeyboardButton struct.
type MessageKeyboardButton struct {
	// A structure describing the type of button action and its parameters
	Action MessageKeyboardButtonAction `json:"action"`
	// Button color
	Color string `json:"color,omitempty"`
}

type KeyboardRow []MessageKeyboardButton

// MessageKeyboard struct.
type MessageKeyboard struct {
	// Keyboard buttons
	Buttons []KeyboardRow `json:"buttons"`
	// Should the keyboard disappear after using it
	OneTime types.ApiBool `json:"one_time,omitempty"`
	// Should the keyboard be displayed inline the message
	Inline types.ApiBool `json:"inline,omitempty"`
}

func NewKeyboard() *MessageKeyboard {
	return &MessageKeyboard{
		Buttons: make([]KeyboardRow, 0, 10),
		OneTime: false,
		Inline:  false,
	}
}

func (k *MessageKeyboard) MakeOneTime() *MessageKeyboard {
	k.OneTime = true

	return k
}

func (k *MessageKeyboard) MakeInline() *MessageKeyboard {
	k.Inline = true

	return k
}

func (k *MessageKeyboard) AddRow() *MessageKeyboard {
	if len(k.Buttons) == 0 {
		k.Buttons = make([]KeyboardRow, 1)
	} else {
		row := make([]MessageKeyboardButton, 0)

		k.Buttons = append(k.Buttons, row)
	}

	return k
}

func (k *MessageKeyboard) AddButton(btn MessageKeyboardButton) *MessageKeyboard {
	lastRow := len(k.Buttons) - 1

	k.Buttons[lastRow] = append(k.Buttons[lastRow], btn)

	return k
}

func (k *MessageKeyboard) AddTextButton(label string, color string, payload ...string) *MessageKeyboard {
	return k.AddButton(MessageKeyboardButton{
		Action: MessageKeyboardButtonAction{
			Type:    ButtonText,
			Label:   label,
			Payload: k.payload(payload),
		},
		Color: color,
	})
}

// AddOpenLinkButton add Open Link button in last row.
func (k *MessageKeyboard) AddOpenLinkButton(link, label string, payload string) *MessageKeyboard {
	return k.AddButton(MessageKeyboardButton{
		Action: MessageKeyboardButtonAction{
			Type:    ButtonOpenLink,
			Payload: payload,
			Label:   label,
			Link:    link,
		},
	})
}

func (k *MessageKeyboard) AddCallbackButton(label string, color string, payload ...string) *MessageKeyboard {
	var p string

	if len(payload) > 0 {
		p = payload[0]
	}

	return k.AddButton(MessageKeyboardButton{
		Action: MessageKeyboardButtonAction{
			Type:    ButtonCallback,
			Label:   label,
			Payload: p,
		},
		Color: color,
	})
}

func (k *MessageKeyboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(*k)
}

func (k *MessageKeyboard) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, k)
	if err != nil {
		return err
	}

	return nil
}

func (k *MessageKeyboard) payload(params []string) string {
	if len(params) > 0 {
		return params[0]
	}

	return ""
}
