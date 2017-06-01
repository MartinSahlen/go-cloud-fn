# go-cloud-fn
`go-cloud-fn` is a tool that allows you to write and deploy [Google cloud functions](https://cloud.google.com/functions/) in pure go.

Run `go get github.com/MartinSahlen/go-cloud-fn` to get it.
You need to have [Google cloud SDK](https://cloud.google.com/sdk/downloads) installed, as well as
the [Cloud functions emulator](https://github.com/GoogleCloudPlatform/cloud-functions-emulator/).

# Usage
Usage is meant to be pretty idiomatic:

### Handling a http request (using [goa](https://github.com/goadesign/goa))
```go
package main

import (
	"os"

	"github.com/MartinSahlen/cloud-goa/app"
	"github.com/MartinSahlen/go-cloud-fn/shim"
	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
)

func main() {
	service := goa.New("adder")

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	service.WithLogger(goakit.New(logger))

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	adder := NewOperandsController(service)
	app.MountOperandsController(service, adder)
  //The key is to use the shim to serve the HTTP
	shim.ServeHTTP(service.Mux.ServeHTTP)
}
```

### Handling a bucket event
```go
package main

import (
	"log"

	"github.com/MartinSahlen/go-cloud-fn/shim"
	"google.golang.org/api/storage/v1beta2"
)

func bucketHandlerFunc(object storage.Object) {
	//Handle the bucket event
	if object.TimeDeleted == "" {
		log.Println("object " + object.Name + " was created or updated")
	} else {
		log.Println("object " + object.Name + " was deleted")
	}
}

func main() {
	shim.HandleBucketEvent(bucketHandlerFunc)
}
```

### Handling a pubsub message
```go
package main

import (
	"log"

	"github.com/MartinSahlen/go-cloud-fn/shim"
	pubsub "google.golang.org/api/pubsub/v1beta2"
)

func pubsubHandleFunc(message pubsub.PubsubMessage) {
	//print the message data (base64 encoded)
	log.Println(message.Data)
}

func main() {
	shim.HandlePubSubMessage(pubsubHandleFunc)
```


Run `go-cloud-fn deploy <function-name>` to deploy your finished function. the [options](https://cloud.google.com/sdk/gcloud/reference/beta/functions/deploy) are listed below:

```
This command lets you deploy your function with a given
set of parameters.

Usage:
  go-cloud-fn deploy <function-name> [flags]

Flags:
  -e, --emulator                Deploy to emulator
  -m, --memory string           Set function memory [MAX 2048MB] (default "1024MB")
	-r, --region string           Set gcloud region
  -s, --stage-bucket string     Set GCS bucket to upload zip bundle
  -o, --timeout string          Set function timeout [MAX 540s] (default "540s")
  -b, --trigger-bucket string   Set function to trigger on this GCS bucket event
  -j, --trigger-http            Set function to trigger on HTTP event
  -t, --trigger-topic string    Set function to trigger on this Pubsub topic
```

Inspired by https://github.com/kelseyhightower/google-cloud-functions-go

## License

Copyright © 2017 Martin Sahlen

Distributed under the MIT License
