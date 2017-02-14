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

type Website struct {
	URL string `json:"url"validate:"required"`
}

func websiteHandler(res express.Response, req express.Request) {
	validate := validator.New()
	var site Website
	err := json.Unmarshal(req.Body, &site)

	res.Headers.Write("content-type", "text/html")

	log.Println(string(req.Body))

	if err != nil {
		log.Println("json error" + err.Error())
		res.Write([]byte(err.Error()))
		return
	}

	err = validate.Struct(site)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, err := range errs {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		bytes, err := json.Marshal(errs)
		if err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		res.Status = 400
		res.Write(bytes)
		return
	}

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
