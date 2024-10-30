package main

import (
	"context"
	requests2 "github.com/defany/govk/api/gen/messages"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"

	"github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
	"github.com/defany/govk/vk"
	"log"
)

func vkInit() {
	vk, err := govk.NewVK("test")
	if err != nil {
		log.Fatal("failed to initialize api client")
	}

	hearManager := hear.NewManager(vk)
	eventHearManager := hear.NewEventManager(vk)

	// register middleware for text commands
	updates.Use(vk.Updates, "message_new", hearManager.Middleware)

	// register middleware for commands from callback buttons
	updates.Use(vk.Updates, "message_event", eventHearManager.Middleware)

	updates.On(vk.Updates, "message_new", func(_ context.Context, msg msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Я ответил на твое сообщение с ID: %d", msg.Message.ID)

		params.WithPeerID(msg.Message.FromID)
		params.WithKeyboard(msgmodel.NewKeyboard().AddRow().AddCallbackButton("huy", "primary"))

		_, err := vk.Api.Messages.MessagesSend(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	updates.On(vk.Updates, "message_event", func(_ context.Context, msg msgmodel.MessagesEvent) {
		params := requests2.NewMessagesSendMessageEventAnswerRequest().WithUserId(msg.UserID).WithPeerId(msg.PeerID).WithEventId(msg.EventID)
		_, err := vk.Api.Messages.MessagesSendMessageEventAnswer(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	handler := hear.NewHandler[msgmodel.MessagesNew]()
	handler.WithCallback(func(ctx context.Context, event msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а!")

		params.WithPeerID(event.Message.PeerID)

		_, err := vk.Api.Messages.MessagesSend(params)
		if err != nil {
			log.Fatalln(err)
		}
	}).WithMatchRules()
	handler.WithMatchRules(hear.WithWordEqualTo("hello"))

	handlerChildren := hear.NewHandler[msgmodel.MessagesNew]()
	handlerChildren.WithCallback(func(ctx context.Context, event msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а и я - ребёнок!")

		params.WithPeerID(event.Message.PeerID)

		_, err := vk.Api.Messages.MessagesSend(params)
		if err != nil {
			log.Fatalln(err)
		}
	})
	handlerChildren.WithMatchRules(hear.WithWordsIn("epta"))
	handler.Group(handlerChildren)

	hearManager.AddHandlers(handler)

	log.Println("😻 bot started")

	err = vk.Updates.Run()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	vkInit()
}
