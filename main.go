package main

import (
	"log"

	"github.com/MartinSahlen/go-cloud-fn/shim"
	"google.golang.org/api/pubsub/v1beta2"
	"google.golang.org/api/storage/v1beta2"
)

func PubsubHandler(message pubsub.PubsubMessage) {
	log.Println(message)
}

func BucketHandler(object storage.Object) {
	log.Println(object)
}

func main() {
	shim.HandleBucketEvent(BucketHandler)
}
