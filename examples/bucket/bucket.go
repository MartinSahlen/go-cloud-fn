package main

import (
	"log"

	"github.com/MartinSahlen/go-cloud-fn/shim/v2"
)

func bucketHandlerFunc(object shimV2.Object) {
	if object.TimeDeleted == "" {
		log.Println("object " + object.Name + " was created or updated")
	} else {
		log.Println("object " + object.Name + " was deleted")
	}
}

func main() {
	shimV2.HandleBucketEvent(bucketHandlerFunc)
}
