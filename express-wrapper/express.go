package express

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gopherjs/gopherjs/js"
)

//Request is the idiomatic (!) go type for a http request
type Request struct {
	Body        []byte              `json:"body,omitempty"`
	Path        string              `json:"path,omitempty"`
	Host        string              `json:"host,omitempty"`
	IPAddresses []string            `json:"ips,omitempty"`
	IPAddress   string              `json:"ip,omitempty"`
	Params      map[string][]string `json:"params,omitempty"`
	Headers     map[string]string   `json:"headers,omitempty"`
	Query       map[string][]string `json:"query,omitempty"`
	Cookies     map[string][]string `json:"cookies,omitempty"`
}

func convertToMapOfStringSlices(i interface{}) (m map[string][]string, err error) {

	m = make(map[string][]string)

	if i == nil {
		err = errors.New("Got nil trying to convert interface to map of string slices")
		return
	}

	tempMap, isMap := i.(map[string]interface{})

	if isMap {

		for k, v := range tempMap {

			s, isString := v.(string)
			if isString {
				m[k] = []string{s}
				break
			}

			si, isSlice := v.([]interface{})

			if isSlice {

				sss, serr := convertToStringSlice(si)
				if serr == nil {
					m[k] = sss
				}
			}
		}
	} else {
		err = errors.New("Not a valid map")
	}
	return
}

func convertToStringSlice(i interface{}) (ss []string, err error) {

	ss = []string{}

	if i == nil {
		err = errors.New("Got nil trying to convert interface to string slice")
		return
	}

	si, isSlice := i.([]interface{})

	if isSlice {
		for _, v := range si {
			s, isString := v.(string)
			if isString {
				ss = append(ss, s)
			}
		}
	} else {
		err = errors.New("Not a valid slice")
	}
	return
}

func convertToBytes(i interface{}) (b []byte, err error) {
	if i == nil {
		err = errors.New("Got nil trying to convert interface to bytes")
		return
	}

	b, err = json.Marshal(i)
	return
}

//ParseRequest wraps a (express request) javascript object into what we need
func ParseRequest(req *js.Object) Request {
	params, err := convertToMapOfStringSlices(req.Get("params").Interface())

	if err != nil {
		log.Println("params: " + err.Error())
	}

	query, err := convertToMapOfStringSlices(req.Get("query").Interface())

	if err != nil {
		log.Println("query: " + err.Error())
	}

	cookies, err := convertToMapOfStringSlices(req.Get("cookies").Interface())

	if err != nil {
		log.Println("query: " + err.Error())
	}

	body, err := convertToBytes(req.Get("body").Interface())

	if err != nil {
		log.Println("body: " + err.Error())
	}

	ips, err := convertToStringSlice(req.Get("ips").Interface())

	if err != nil {
		log.Println("ips: " + err.Error())
	}

	headers := make(map[string]string)
	headers["authorization"] = req.Call("get", "authorization").String()
	headers["content-type"] = req.Call("get", "content-type").String()

	return Request{
		Path:        req.Get("path").String(),
		IPAddress:   req.Get("ip").String(),
		Host:        req.Get("host").String(),
		IPAddresses: ips,
		Headers:     headers,
		Params:      params,
		Body:        body,
		Query:       query,
		Cookies:     cookies,
	}
}

//JSON is returns a STRING; not bytes
func (r Request) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		panic(err)
	} else {
		return string(byt)
	}
}
