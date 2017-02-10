package main

import (
	"github.com/MartinSahlen/go-cloud-fn/express-wrapper"
	"github.com/gopherjs/gopherjs/js"
)

//Handle is the main handler and entrypoint for the google cloud function
func Handle(req, res *js.Object) {
	request := express.ParseRequest(req)

	res.Call("set", "Content-Type", "application/json")
	res.Call("send", request.JSON())
}

func main() {
	js.Module.Get("exports").Set("helloGO", Handle)
}
