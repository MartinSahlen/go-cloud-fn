package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

type Website struct {
	URL string `json:"url"`
}

func websiteHandler(res express.Response, req express.Request) {

	var site Website
	err := json.Unmarshal(req.Body, &site)

	res.Headers.Write("content-type", "text/html")

	log.Println(string(req.Body))

	if err != nil {
		log.Println("json error" + err.Error())
		res.Write([]byte(err.Error()))
		return
	}

	log.Println(site)

	go func() {
		r, err := http.DefaultClient.Get(site.URL)

		if err != nil {
			log.Println(err.Error())
			res.Write([]byte(err.Error()))
			return
		}

		byt, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			res.Write([]byte(err.Error()))
			return
		}

		res.Status = 200
		res.Write(byt)
	}()

}

//EntryPoint is the main handler and entrypoint for the google cloud function
func EntryPoint(req, res *js.Object) {

	r := router.New(rootHandler)
	r.Handle("GET", "/hello/:ergegr", helloHandler)
	r.Handle("POST", "/site", websiteHandler)

	r.Serve(express.NewResponse(res), express.NewRequest(req))
}

func main() {
	if googleCloudFunctionName == "" {
		googleCloudFunctionName = "helloGO"
	}
	js.Module.Get("exports").Set(googleCloudFunctionName, EntryPoint)
}
