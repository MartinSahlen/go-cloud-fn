package main

import (
	"github.com/MartinSahlen/go-cloud-fn/express-wrapper"
	"github.com/gopherjs/gopherjs/js"
)

//Handle is the main handler and entrypoint for the google cloud function
func RootHandler(req, res *js.Object) {
	request := express.NewRequest(req)
	response := express.NewResponse(res)
	response.Headers.Write("content-type", "application/json")
	response.Status = 201
	response.Write(request.JSON())
}

func main() {
	js.Module.Get("exports").Set("helloGO", RootHandler)
}
