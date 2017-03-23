package shim

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type HTTPRequest struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	Method     string            `json:"method"`
	RemoteAddr string            `json:"remote_addr"`
	URL        string            `json:"url"`
}

type HTTPResponse struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
}

type RequestHandler func(http.ResponseWriter, *http.Request)

func DebugRequest(r *http.Request) {

	params := ""
	for k, v := range r.URL.Query() {
		params += k + ": " + strings.Join(v, "), (")
	}
	headers := ""
	for k, v := range r.Header {
		headers += k + ": " + strings.Join(v, "), (")
	}

	log.Println("Host: " + r.Host)

	log.Println("URL Hostname: " + r.URL.Hostname())
	log.Println("URL Query: " + params)
	log.Println("URL Headers: " + headers)
	log.Println("URL Path: " + r.URL.Path)
	log.Println("URL Port: " + r.URL.Port())
	log.Println("URL Request URI: " + r.URL.RequestURI())
}

func ServeHTTP(h RequestHandler) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var httpRequest HTTPRequest
	err = json.Unmarshal(stdin, &httpRequest)
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.NewRequest(httpRequest.Method, httpRequest.URL, bytes.NewBufferString(httpRequest.Body))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range httpRequest.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = httpRequest.RemoteAddr

	if r.URL.Path == "" {
		r.URL.Path = "/"
	}

	w := httptest.NewRecorder()

	h(w, r)

	resp := w.Result()

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}
	httpResponse := HTTPResponse{
		Body:       w.Body.String(),
		Header:     header,
		StatusCode: resp.StatusCode,
	}

	out, err := json.Marshal(&httpResponse)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(out)
}
