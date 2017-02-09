package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
)

type Request struct{}

type Response struct{}

func (res Response) send(data []byte) {}

func HelloWorld(req Request, res Response) {
	log.Println(req, res)
	res.send([]byte("Hello world!"))
}

func main() {
	js.Module.Get("exports").Set("helloyolo", HelloWorld)
}
