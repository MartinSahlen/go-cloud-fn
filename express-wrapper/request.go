package express

import (
	"encoding/json"
	"log"

	"github.com/gopherjs/gopherjs/js"
)

//Request is the idiomatic (!) go type for a http request
type Request struct {
	Body        []byte              `json:"body,omitempty"`
	Path        string              `json:"path,omitempty"`
	Method      string              `json:"method,omitempty"`
	Host        string              `json:"host,omitempty"`
	IPAddresses []string            `json:"ips,omitempty"`
	IPAddress   string              `json:"ip,omitempty"`
	Params      map[string][]string `json:"params,omitempty"`
	Headers     map[string]string   `json:"headers,omitempty"`
	Query       map[string][]string `json:"query,omitempty"`
	Cookies     map[string][]string `json:"cookies,omitempty"`
	Raw         *js.Object          `json:"cookies,omit"`
}

//ParseRequest wraps a (express request) javascript object into what we need
func NewRequest(req *js.Object) Request {
	params, err := convertToMapOfStringSlices(nil)

	if err != nil {
		log.Println("params: " + err.Error())
	}

	query, err := convertToMapOfStringSlices(req.Get("query").Interface())

	if err != nil {
		log.Println("query: " + err.Error())
	}

	cookies, err := convertToMapOfStringSlices(req.Get("cookies").Interface())

	if err != nil {
		log.Println("cookies: " + err.Error())
	}

	body, err := convertToBytes(req.Get("body").Interface())

	if err != nil {
		log.Println("body: " + err.Error())
	}

	ips, err := convertToStringSlice(req.Get("ips").Interface())

	if err != nil {
		log.Println("ips: " + err.Error())
	}

	headers, err := convertToMapOfStrings(req.Get("headers").Interface())

	if err != nil {
		log.Println("headers: " + err.Error())
	}

	var path string
	pathObject := req.Get("path")

	if pathObject == nil {
		path = "/"
	} else {
		path = pathObject.String()
	}

	return Request{
		IPAddress:   req.Get("ip").String(),
		Host:        req.Get("hostname").String(),
		Method:      req.Get("method").String(),
		Path:        path,
		IPAddresses: ips,
		Headers:     headers,
		Params:      params,
		Body:        body,
		Query:       query,
		Cookies:     cookies,
		Raw:         req,
	}
}

//JSON is returns a STRING; not bytes
func (r Request) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return byt
}
