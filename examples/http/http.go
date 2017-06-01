package main

import (
	"io"
	"net/http"

	"github.com/MartinSahlen/go-cloud-fn/shim"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	shim.ServeHTTP(handler)
}
