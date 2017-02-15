package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/MartinSahlen/go-cloud-fn/express-wrapper"
	"github.com/MartinSahlen/go-cloud-fn/router"
	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/go-playground/validator.v9"
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

type website struct {
	URL string `json:"url"validate:"required,url"`
}

func websiteHandler(res express.Response, req express.Request) {
	validate := validator.New()
	var site website
	err := json.Unmarshal(req.Body, &site)

	res.Headers.Write("content-type", "text/html")

	if err != nil {
		log.Println("json error" + err.Error())
		res.Write([]byte(err.Error()))
		return
	}

	err = validate.Struct(site)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
		}
		res.Status = 400
		res.Write([]byte(errs.Error()))
		return
	}

	go func() {

		packageFile, err := os.Open("./package.json")
		b, err := ioutil.ReadAll(packageFile)
		log.Println(string(b))

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
	r.Handle(http.MethodGet, "/hello/:ergegr", helloHandler)
	r.Handle(http.MethodPost, "/site", websiteHandler)

	r.Serve(express.NewResponse(res), express.NewRequest(req))
}

func main() {
	if googleCloudFunctionName == "" {
		googleCloudFunctionName = "helloGO"
	}
	js.Module.Get("exports").Set(googleCloudFunctionName, EntryPoint)
}
