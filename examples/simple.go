package main

import (
	"context"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"
	heargo "github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
	"github.com/defany/govk/vk"
	"log"
)

func vkInit() {
	vk, err := govk.NewVK("vk1.a.iOkLDJjKRydTL05R9Ve6SeIlDD5BiH-AwPYJ9L8x66J2slL-COSuNXdtwKs-mRghe9EgFDg0fqT2Y6YIPvP4GanOB3nfyAaFW60h-okf1wD6NDJYV2l7S9U43vVCthNmw-lA0IuNUb78p4F-8DESWA8R0hJDEAuC0uljr7MehGRB_wXDJPPfxuvEv6C0_ZfHT0rjhu0cDuSbZK63gonnsg")
	if err != nil {
		log.Fatal("failed to initialize api client")
	}

	hearManager := heargo.NewManager(vk)

	updates.Use(vk.Updates, "message_new", hearManager.Middleware)

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
	handler.WithMatchRules(heargo.WithMatchWord("123"), heargo.WithWordsIn("hello", "hi"))

	anotherOneHandler := hearManager.NewHandler(func(ctx context.Context, event msgmodel.MessagesNew) {
		params := requests.NewSendParams().
			WithMessage("Привет! Это я ответил на твое сообщение при помощи hear manager'а! Но уже от другого хендлера!")

		params.WithPeerID(event.Message.PeerID)

		_, err := vk.Api.Messages.Send(params)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// your own match rule
	anotherOneHandler.WithMatchRules(func(event msgmodel.MessagesNew) bool {
		return event.Message.Text == "hey!"
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
