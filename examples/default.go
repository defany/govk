package main

import (
	"context"
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"
	heargo "github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
	"github.com/defany/govk/vk"
	"log"
	"time"
)

func vkInit() {
	vk, err := govk.NewVK("vk1.a.kLGED5Km5A4nkwuufWB9JYkCkh4XYnM3ttbre3xOpTaFu027ma_qFKiLFMkksJJIzdxA1itQY_fERu5_FSTJNb3-IjY29KstAgj9t_iekEWaxOCz6IbHTc-JuWVHnbgU0DQjxK2_kCpkov_BSdCbMAGaYHUuuqnKYkDhtgc30npyvD39Hh-H1ZnogII6rEhLiZ7Ll1QgPxzUnW5wQ83oNw")
	if err != nil {
		log.Fatal("failed to initialize api client")
	}

	vk.WithApiLimit(1)

	kb := msgmodel.NewKeyboard()

	kb.MakeInline()

	kb.AddRow()
	kb.AddTextButton("test", msgmodel.ButtonBlue)

	params := requests.NewSendParams().WithMessage("Привет!")

	params.WithPeerID(222856843, 620893364)
	params.WithKeyboard(kb)

	_, err = vk.Api.Messages.Send(params)
	if err != nil {
		log.Fatalln(err)
	}

	updates.On(vk.Updates, "message_new", func(msg msgmodel.MessagesNew) {
		log.Println(msg.Message.ID)
	})

	updates.On(vk.Updates, "messages_event", func(msg msgmodel.MessagesEvent) {
		kb := msgmodel.NewKeyboard()

		kb.MakeInline()

		kb.AddTextButton("test", msgmodel.ButtonBlue)

		res, err := vk.Api.Messages.Send(api.Params{
			"peer_id":   msg.UserID,
			"random_id": time.Now().UnixMilli(),
			"message":   "amogus",
			"keyboard":  kb,
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(res)
	})

	handler := heargo.NewHandler(func(ctx context.Context, messagesNew msgmodel.MessagesNew) {

	})

	// builtin match validators
	handler.WithMatchRules(heargo.WithMatchWord("123"), heargo.WithWordsIn("hello", "hi", "hey"))

	// your own match rule
	handler.WithMatchRules(func(event msgmodel.MessagesNew) bool {
		return event.Message.Text == "hey!"
	})

	err = vk.Updates.Check()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	vkInit()
}
