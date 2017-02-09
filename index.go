package main

import (
	"encoding/json"

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

func HelloWorld(req *js.Object, res *js.Object) {

	r := Result{"Hello, world", 1}

	res.Call("set", "Content-Type", "application/json")
	res.Call("send", string(r.JSON()))
}

func main() {
	js.Module.Get("exports").Set("helloGO", HelloWorld)
}
