package main

import (
	"context"
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/api/messages/requests"
	heargo "github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
<<<<<<< HEAD
	govk "github.com/defany/govk/vk"
=======
	"github.com/defany/govk/vk"
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1
	"log"
	"time"
)

func vkInit() {
<<<<<<< HEAD
	vk, err := govk.NewVK("vk1.a.iOkLDJjKRydTL05R9Ve6SeIlDD5BiH-AwPYJ9L8x66J2slL-COSuNXdtwKs-mRghe9EgFDg0fqT2Y6YIPvP4GanOB3nfyAaFW60h-okf1wD6NDJYV2l7S9U43vVCthNmw-lA0IuNUb78p4F-8DESWA8R0hJDEAuC0uljr7MehGRB_wXDJPPfxuvEv6C0_ZfHT0rjhu0cDuSbZK63gonnsg")
=======
	vk, err := govk.NewVK("vk1.a.kLGED5Km5A4nkwuufWB9JYkCkh4XYnM3ttbre3xOpTaFu027ma_qFKiLFMkksJJIzdxA1itQY_fERu5_FSTJNb3-IjY29KstAgj9t_iekEWaxOCz6IbHTc-JuWVHnbgU0DQjxK2_kCpkov_BSdCbMAGaYHUuuqnKYkDhtgc30npyvD39Hh-H1ZnogII6rEhLiZ7Ll1QgPxzUnW5wQ83oNw")
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1
	if err != nil {
		log.Fatal("failed to initialize api client")
	}

<<<<<<< HEAD
	vk.WithApiLimit(1)
=======
	vk.Api.Api.WithLimit(1)
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1

	kb := msgmodel.NewKeyboard()

	kb.MakeInline()

	kb.AddRow()
	kb.AddTextButton("test", msgmodel.ButtonBlue)

	params := requests.NewSendParams().WithMessage("Привет!")

	params.WithPeerID(222856843, 620893364)
	params.WithKeyboard(kb)

<<<<<<< HEAD
	res, err := vk.Api.Messages.Send(params)
=======
	_, err = vk.Api.Messages.Send(params)
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1
	if err != nil {
		log.Fatalln(err)
	}

<<<<<<< HEAD
	log.Println(res)

	updates.On(vk.Updates, "messages_new", func(msg msgmodel.MessagesNew) {
=======
	updates.On(vk.Updates, "message_new", func(msg msgmodel.MessagesNew) {
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1
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
<<<<<<< HEAD
=======

	err = vk.Updates.Run()
	if err != nil {
		log.Println(err)
	}
>>>>>>> 85c1ed48c7bab2e18d8484e800a82d509b607ae1
}

func main() {
	vkInit()
}
