package main

import (
	"github.com/defany/govk/api"
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

	res, err := vk.Api.Messages.Send(api.Params{
		"peer_id":   222856843,
		"random_id": time.Now().UnixMilli(),
		"message":   "amogus",
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func main() {
	vkInit()
}
