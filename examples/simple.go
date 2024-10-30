<<<<<<< HEAD
package main

import (
	"context"
	requests2 "github.com/defany/govk/api/gen/messages"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"

	hear "github.com/defany/govk/hear"
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
||||||| 06ea378
=======
package main

import (
	"context"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"
	hear "github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
	"github.com/defany/govk/vk"
	"log"
)

func vkInit() {
	vk, err := govk.NewVK("vk1.a.iOkLDJjKRydTL05R9Ve6SeIlDD5BiH-AwPYJ9L8x66J2slL-COSuNXdtwKs-mRghe9EgFDg0fqT2Y6YIPvP4GanOB3nfyAaFW60h-okf1wD6NDJYV2l7S9U43vVCthNmw-lA0IuNUb78p4F-8DESWA8R0hJDEAuC0uljr7MehGRB_wXDJPPfxuvEv6C0_ZfHT0rjhu0cDuSbZK63gonnsg")
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

		params.WithPeerID(222856843)

		_, err := vk.Api.Messages.Send(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// making a new handler
	handler := hearManager.NewHandler(func(ctx context.Context, event msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а!")

		params.WithPeerID(event.Message.PeerID)

		_, err := vk.Api.Messages.Send(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// builtin match validators
	handler.WithMatchRules(hear.WithMatchWord("123"), hear.WithWordsIn("hello", "hi"))

	anotherOneHandler := eventHearManager.NewHandler(func(ctx context.Context, event msgmodel.MessagesEvent) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а! Но уже от другого хендлера!")

		params.WithPeerID(event.PeerID)

		_, err := vk.Api.Messages.Send(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// your own match rule
	anotherOneHandler.WithMatchRules(func(event msgmodel.MessagesEvent) bool {
		return event.UserID == 222856843
	})

	log.Println("😻 bot started")

	err = vk.Updates.Run()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	vkInit()
}
>>>>>>> ec6cd9b5c35171d69d867df2d0a7cfd6f45a4d12
