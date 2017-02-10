package main

import (
	"net/url"

	"github.com/MartinSahlen/go-cloud-fn/express-wrapper"
	"github.com/MartinSahlen/go-cloud-fn/router"
	"github.com/gopherjs/gopherjs/js"
)

func RootHandler(res express.Response, req express.Request, params url.Values) {
	res.Headers.Write("content-type", "application/json")
	res.Status = 404
	res.Write(req.JSON())
}

func HelloHandler(res express.Response, req express.Request, params url.Values) {
	res.Headers.Write("content-type", "application/json")
	res.Status = 200
	res.Write(req.JSON())
}

//Handle is the main handler and entrypoint for the google cloud function
func EntryPoint(req, res *js.Object) {

	r := router.New(RootHandler)
	r.Handle("GET", "/hello/:ergegr", HelloHandler)

	r.Serve(express.NewResponse(res), express.NewRequest(req))
}

func main() {
	js.Module.Get("exports").Set("helloGO", EntryPoint)
}
