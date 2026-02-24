package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type someType struct {
	Text string `json:"text"`
}

func getSomeHandler(somes *map[int]someType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'getSomeHandler'")
		jsonBytes, err := json.MarshalIndent(*somes, "", "  ")
		if err != nil {
			slog.Error("Can't code somes", slog.Any("error", err))
		}
		fmt.Fprintf(w, "%s", jsonBytes)
	}
}

func postSomeHandler(id *int, somes *map[int]someType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'postSomeHandler'")

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var tmp someType

		if err := dec.Decode(&tmp); err != nil {
			slog.Error("Can't decode body", slog.Any("error", err))
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}
		(*somes)[*id] = tmp
		(*id)++
	}
}

func putSomeHandler(somes *map[int]someType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'putSomeHandler'")

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error("Can't parse id", slog.Any("error", err))
			http.Error(w, "Can't parse id", http.StatusBadRequest)
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var tmp someType

		if err := dec.Decode(&tmp); err != nil {
			slog.Error("Can't decode body", slog.Any("error", err))
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}
		(*somes)[id] = tmp
	}
}

func main() {
	mux := http.NewServeMux()

	id := new(int(0))
	somes := new(make(map[int]someType, 0))

	mux.HandleFunc("GET /some", getSomeHandler(somes))
	mux.HandleFunc("POST /some", postSomeHandler(id, somes))
	mux.HandleFunc("PUT /some/{id}", putSomeHandler(somes))

	slog.Info("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
	}
}
