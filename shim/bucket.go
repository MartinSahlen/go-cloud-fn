package shim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	storage "google.golang.org/api/storage/v1beta2"
)

type BucketHandlerFunc func(object storage.Object)

func HandleBucketEvent(h BucketHandlerFunc) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var object storage.Object
	err = json.Unmarshal(stdin, &object)
	if err != nil {
		log.Fatal(err)
	}
	h(object)
}
