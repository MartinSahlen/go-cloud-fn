package main

import "github.com/gopherjs/gopherjs/js"

const Message = "Hello world!"

func HelloWorld(req *js.Object, res *js.Object) {
	res.Call("send", Message)
}

func main() {
	js.Module.Get("exports").Set("helloGO", HelloWorld)
}
