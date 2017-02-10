package main

import (
	"os"

	"github.com/MartinSahlen/go-cloud-fn/express-wrapper"
	"github.com/MartinSahlen/go-cloud-fn/router"
	"github.com/gopherjs/gopherjs/js"
)

var googleCloudFunctionName = os.Getenv("GOOGLE_CLOUD_FUNCTION_NAME")

func rootHandler(res express.Response, req express.Request) {
	res.Headers.Write("content-type", "application/json")
	res.Status = 404
	res.Write(req.JSON())
}

func helloHandler(res express.Response, req express.Request) {
	res.Headers.Write("content-type", "application/json")
	res.Status = 200
	res.Write(req.JSON())
}

//EntryPoint is the main handler and entrypoint for the google cloud function
func EntryPoint(req, res *js.Object) {

	r := router.New(rootHandler)
	r.Handle("GET", "/hello/:ergegr", helloHandler)

	r.Serve(express.NewResponse(res), express.NewRequest(req))
}

func main() {
	if googleCloudFunctionName == "" {
		googleCloudFunctionName = "helloGO"
	}
	js.Module.Get("exports").Set(googleCloudFunctionName, EntryPoint)
}
