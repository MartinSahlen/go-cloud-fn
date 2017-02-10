package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
)

type Result struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

func (r Result) JSON() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return b
}

func HelloWorld(req, res *js.Object) {

	go func() {
		resp, err := http.Get("https://google.com/")
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		r := Result{string(body), rand.Int()}
		res.Call("set", "Content-Type", "application/json")
		res.Call("send", string(r.JSON()))
	}()
}

func main() {
	//js.Global.Get("require").Invoke("es6-promise").Call("polyfill")
	//js.Global.Get("require").Invoke("isomorphic-fetch")
	js.Module.Get("exports").Set("helloGO", HelloWorld)
}
