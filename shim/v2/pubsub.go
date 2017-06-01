package shimV2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type PubsubMessage struct {
	Attributes  map[string]string `json:"attributes,omitempty"`
	Data        string            `json:"data,omitempty"`
	MessageId   string            `json:"messageId,omitempty"`
	PublishTime string            `json:"publishTime,omitempty"`
}

type PubsubHandlerFunc func(message PubsubMessage)

func HandlePubSubMessage(h PubsubHandlerFunc) {

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var message PubsubMessage
	err = json.Unmarshal(stdin, &message)
	if err != nil {
		log.Fatal(err)
	}

	h(message)
}
