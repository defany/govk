package main

import (
	"context"
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/messages/model"
	heargo "github.com/defany/govk/hear"
	"github.com/defany/govk/updates"
	govk "github.com/defany/govk/vk"
	"log"
	"time"
)

func vkInit() {
	vk, err := govk.NewVK("vk1.a.iOkLDJjKRydTL05R9Ve6SeIlDD5BiH-AwPYJ9L8x66J2slL-COSuNXdtwKs-mRghe9EgFDg0fqT2Y6YIPvP4GanOB3nfyAaFW60h-okf1wD6NDJYV2l7S9U43vVCthNmw-lA0IuNUb78p4F-8DESWA8R0hJDEAuC0uljr7MehGRB_wXDJPPfxuvEv6C0_ZfHT0rjhu0cDuSbZK63gonnsg")
	if err != nil {
		log.Fatal("failed to initialize api client")
	}

	vk.WithApiLimit(1)

	updates.On(vk.Updates, "messages_new", func(msg model.MessagesNew) {
		log.Println(msg.Message.ID)
	})

	updates.On(vk.Updates, "messages_event", func(msg model.MessagesEvent) {
		res, err := vk.Api.Messages.Send(api.Params{
			"peer_id":   msg.UserID,
			"random_id": time.Now().UnixMilli(),
			"message":   "amogus",
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(res)
	})

	handler := heargo.NewHandler(func(ctx context.Context, messagesNew model.MessagesNew) {

	})

	handler.WithPreValidator(func(ctx context.Context, event model.MessagesNew) {

	})

	handler.WithMatchRules(heargo.WithMatchWord("123"), heargo.WithWordsIn("hello", "hi", "hey"))
}

func main() {
	vkInit()
}
