package main

import (
	"encoding/json"
	"net/http"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"

	"github.com/MartinSahlen/go-cloud-fn/shim"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client, err := bigquery.NewClient(r.Context(), os.Getenv("GCLOUD_PROJECT"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	datasetsIterator := client.Datasets(r.Context())
	datasets := []*bigquery.Dataset{}
	for {
		dataset, err := datasetsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		datasets = append(datasets, dataset)
	}

	byt, err := json.Marshal(datasets)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byt)
}

func main() {
	shim.ServeHTTP(handler)
}
