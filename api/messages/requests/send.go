package requests

import (
	"fmt"
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/messages/model"
	"github.com/goccy/go-json"
	"reflect"
	"time"
)

// TODO: add a builder

var marshalingMessageKeys = []api.ParamMarshaler{
	{
		Key: msgmodel.KeyboardParam,
		Marshaler: func(v any) (any, error) {
			if reflect.TypeOf(v).Kind() == reflect.String {
				return v, nil
			}

			marshaler, ok := v.(json.Marshaler)
			if !ok {
				d, err := json.Marshal(v)
				if err != nil {
					return "", fmt.Errorf("cannot transform keyboard to json, recheck your keyboard code: %w", err)
				}

				return api.FmtValue(d, 0)
			}

			d, err := marshaler.MarshalJSON()
			if err != nil {
				return "", err
			}

			return string(d), nil
		},
	},
	{
		Key: msgmodel.TemplateParam,
		Marshaler: func(v any) (any, error) {
			if reflect.TypeOf(v).Kind() == reflect.String {
				return v, nil
			}

			marshaler, ok := v.(json.Marshaler)
			if !ok {
				d, err := json.Marshal(v)
				if err != nil {
					return "", fmt.Errorf("cannot cast keyboard to json, recheck your template code: %w", err)
				}

				return api.FmtValue(d, 0)
			}

			d, err := marshaler.MarshalJSON()
			if err != nil {
				return "", err
			}

			return string(d), nil
		},
	},
	{
		Key: msgmodel.PeerIDsParam,
		Marshaler: func(v any) (any, error) {
			return api.FmtValue(v, 0)
		},
	},
}

type SendParams api.Params

// NewSendParams - function to create and pass params for messages.Send()
//
// When created, passed a random id from this value: time.Now().UnixMilli()
//
// You can always pass your own params, like this:
// params["ur_unique_key"] = "your_data"
//
// Try to use builtin params, they all called as in vk but with postfix `Param`, for example:
// message param -> msgmodel.MessageParam
func NewSendParams() SendParams {
	params := make(SendParams, 10)

	params.WithRandomID(int(time.Now().UnixMilli()))

	return params
}

// WithKeyboard passes keyboard to params
//
// If u want to pass your own implementation of keyboard, just use this code:
// params[KeyboardParam] = yourOwnKeyboard
func (s SendParams) WithKeyboard(keyboard *msgmodel.MessageKeyboard) SendParams {
	s[msgmodel.KeyboardParam] = keyboard

	return s
}

// WithPeerID passes a peer ids to params
//
// Remember, that you can pass only 100 peer ids per message
func (s SendParams) WithPeerID(peerID ...int) SendParams {
	s[msgmodel.PeerIDsParam] = peerID

	return s
}

// WithMessage passes a text message to params
//
// You can also pass the params like for fmt.Sprintf
func (s SendParams) WithMessage(message string, args ...any) SendParams {
	s[msgmodel.MessageParam] = fmt.Sprintf(message, args...)

	return s
}

// WithRandomID u can set your own random id to params
func (s SendParams) WithRandomID(randomID int) SendParams {
	s[msgmodel.RandomIDParam] = randomID

	return s
}

// Params for compatability, nothing interesting
func (s SendParams) Params() api.Params {
	return api.Params(s)
}

type MessagesSendResponse struct {
	PeerID                int       `json:"peer_id"`
	MessageID             int       `json:"message_id"`
	ConversationMessageID int       `json:"conversation_message_id"`
	Error                 api.Error `json:"error"`
}

// Send - sending message to vk
//
// Be careful, because this method is using a peer_ids to send a message to users and do not return just a number from VK API
func (m *Messages) Send(params api.MethodParams) ([]MessagesSendResponse, error) {
	req := api.NewRequest[[]MessagesSendResponse](m.api)

	data := params.Params()

	for _, marshaling := range marshalingMessageKeys {
		v, ok := data[marshaling.Key]
		if !ok {
			continue
		}

		m, err := marshaling.Marshaler(v)
		if err != nil {
			return nil, err
		}

		data[marshaling.Key] = m
	}

	res, err := req.Execute(m.methodsGroup+"send", params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
