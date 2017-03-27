package shim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/api/pubsub/v1beta2"
)

type PubsubHandlerFunc func(message pubsub.PubsubMessage)

func HandlePubSubMessage(h PubsubHandlerFunc) {

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var message pubsub.PubsubMessage
	err = json.Unmarshal(stdin, &message)
	if err != nil {
		log.Fatal(err)
	}

	h(message)
}
