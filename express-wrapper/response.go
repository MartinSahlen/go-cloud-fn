package express

import "github.com/gopherjs/gopherjs/js"

type Response struct {
	Raw     *js.Object
	Status  int
	Headers Headers
}

type Headers struct {
	Data map[string]string
}

func (h Headers) Write(name, value string) {
	h.Data[name] = value
}

func (r Response) Write(data []byte) {
	for k, v := range r.Headers.Data {
		r.Raw.Call("set", k, v)
	}
	r.Raw.Call("status", r.Status)
	r.Raw.Call("send", string(data))
}

func NewResponse(req *js.Object) Response {
	return Response{
		Raw:     req,
		Status:  200,
		Headers: Headers{make(map[string]string)},
	}
}
