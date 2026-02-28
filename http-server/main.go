package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"serv/internal/config"
	"serv/internal/handler"
	"serv/internal/model"
	"serv/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", slog.Any("error", err))
	}

	mux := http.NewServeMux()

	server := &store.Server{
		Somes: make(map[int]model.SomeType),
	}

	mux.HandleFunc("GET /somes", handler.GetAllSomeHandler(server))
	mux.HandleFunc("GET /somes/{id}", handler.GetSomeByIdHandler(server))
	mux.HandleFunc("POST /somes", handler.PostSomeHandler(server))
	mux.HandleFunc("PUT /somes/{id}", handler.PutSomeHandler(server))
	mux.HandleFunc("DELETE /somes/{id}", handler.DeleteSomeByIdHandler(server))

	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	slog.Info("Server listening on", slog.String("address", addr))

	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
	}
}
