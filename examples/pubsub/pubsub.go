package main

import (
	"log"

	"github.com/MartinSahlen/go-cloud-fn/shim/v2"
)

func pubsubHandleFunc(message shimV2.PubsubMessage) {
	log.Println(message.Data)
}

func main() {
	shimV2.HandlePubSubMessage(pubsubHandleFunc)
}
