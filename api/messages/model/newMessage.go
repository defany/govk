package model

import "github.com/defany/govk/api/types"

// TODO: add missed params in message obj

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

// MessageKeyboard struct.
type MessageKeyboard struct {
	// Keyboard buttons
	Buttons [][]MessageKeyboardButton `json:"buttons"`
	// Should the keyboard disappear after using it
	OneTime types.ApiBool `json:"one_time,omitempty"`
	// Should the keyboard be displayed inline the message
	Inline types.ApiBool `json:"inline,omitempty"`
}

type Message struct {
	// Only for messages from community. Contains user ID of community admin,
	// who sent this message.
	AdminAuthorID int `json:"admin_author_id"`
	// FIXME: refactor handle attachments in vk-io style
	Attachments []any `json:"attachments"`

	// Unique auto-incremented number for all messages with this peer.
	ConversationMessageID int `json:"conversation_message_id"`

	// Date when the message has been sent in Unixtime.
	Date int `json:"date"`

	// Message author's ID.
	FromID int `json:"from_id"`

	// Forwarded messages.
	FwdMessages  []Message     `json:"fwd_Messages"`
	ReplyMessage *Message      `json:"reply_message"`
	PinnedAt     int           `json:"pinned_at,omitempty"`
	ID           int           `json:"id"`        // Message ID
	Deleted      types.ApiBool `json:"deleted"`   // Is it an deleted message
	Important    types.ApiBool `json:"important"` // Is it an important message
	IsHidden     types.ApiBool `json:"is_hidden"`
	IsCropped    types.ApiBool `json:"is_cropped"`
	IsSilent     types.ApiBool `json:"is_silent"`
	Out          types.ApiBool `json:"out"` // Information whether the message is outcoming
	WasListened  types.ApiBool `json:"was_listened,omitempty"`
	// FIXME: refactor handle keyboard and template in vk-io style
	Keyboard MessageKeyboard `json:"keyboard"`
	Template any             `json:"template"`
	Payload  string          `json:"payload"`
	PeerID   int             `json:"peer_id"` // Peer ID

	// ID used for sending messages. It returned only for outgoing messages.
	RandomID     int    `json:"random_id"`
	Ref          string `json:"ref"`
	RefSource    string `json:"ref_source"`
	Text         string `json:"text"`          // Message text
	UpdateTime   int    `json:"update_time"`   // Date when the message has been updated in Unixtime
	MembersCount int    `json:"members_count"` // Members number
	ExpireTTL    int    `json:"expire_ttl"`
	MessageTag   string `json:"message_tag"` // for https://notify.mail.ru/
}

type MessagesNew struct {
	Message    Message    `json:"message"`
	ClientInfo ClientInfo `json:"client_info"`
}
