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
	log.Println("Got event for object: " + object.Name + " in bucket " + object.Bucket)
	if object.TimeDeleted == "" {
		log.Println("The object was created")
	} else {
		log.Println("The object was deleted")
	}
}

func main() {
	shim.HandleBucketEvent(BucketHandler)
}
