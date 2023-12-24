package model

import "github.com/goccy/go-json"

type MessagesEvent struct {
	UserID                int             `json:"user_id"`
	PeerID                int             `json:"peer_id"`
	EventID               string          `json:"event_id"`
	Payload               json.RawMessage `json:"payload"`
	ConversationMessageID int             `json:"conversation_message_id"`
}
