package msgmodel

import (
	"github.com/defany/govk/api/types"
)

// TODO: add missed params in message obj

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
