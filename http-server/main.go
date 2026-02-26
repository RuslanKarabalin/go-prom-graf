package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

type someType struct {
	Text string `json:"text"`
}

type server struct {
	nextId int
	somes  map[int]someType
	mtx    sync.RWMutex
}

func getSomeHandler(s *server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'getSomeHandler'")
		w.Header().Set("Content-Type", "application/json")
		s.mtx.RLock()
		defer s.mtx.RUnlock()
		if err := json.NewEncoder(w).Encode(s.somes); err != nil {
			slog.Error("Can't encode somes", slog.Any("error", err))
		}
	}
}

func postSomeHandler(s *server) http.HandlerFunc {
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

		s.mtx.Lock()
		defer s.mtx.Unlock()
		s.somes[s.nextId] = tmp
		s.nextId++

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]int{"id": s.nextId - 1})
	}
}

func putSomeHandler(s *server) http.HandlerFunc {
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

		s.mtx.Lock()
		defer s.mtx.Unlock()
		if _, exists := s.somes[id]; !exists {
			slog.Error("Can't find id")
			http.Error(w, "Can't find id", http.StatusNotFound)
			return
		}
		s.somes[id] = tmp
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	mux := http.NewServeMux()

	server := &server{
		somes: make(map[int]someType),
	}

	mux.HandleFunc("GET /somes", getSomeHandler(server))
	mux.HandleFunc("POST /somes", postSomeHandler(server))
	mux.HandleFunc("PUT /somes/{id}", putSomeHandler(server))

	slog.Info("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
	}
}
