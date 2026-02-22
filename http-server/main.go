package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type someType struct {
	Text string `json:"text"`
}

var some someType

func getSomeHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Call 'getSomeHandler'")
	fmt.Fprintf(w, "%s", some.Text)
}

func postSomeHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Call 'postSomeHandler'")

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var tmp someType

	if err := dec.Decode(&tmp); err != nil {
		slog.Error("Can't decode body", slog.Any("error", err))
		http.Error(w, "Can't decode body", http.StatusBadRequest)
		return
	}
	some = tmp
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /some", getSomeHandler)
	mux.HandleFunc("POST /some", postSomeHandler)

	slog.Info("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
	}
}
