package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"serv/internal/config"
	"serv/internal/handler"
	"serv/internal/model"
	"serv/internal/store"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	pgUrl := cfg.GetPgUrl()

	slog.Info("PgUrl:", slog.Any("pgUrl", pgUrl))

	conn, err := pgxpool.New(ctx, pgUrl)
	if err != nil {
		slog.Error("Cannot connect to PostgreSQL", slog.Any("error", err))
		os.Exit(1)
	}
	defer conn.Close()

	// bad migration for start
	createTableSomes := `
	create table if not exists somes(
		id int generated always as identity primary key,
		name text not null
	)
	`
	_, err = conn.Exec(ctx, createTableSomes)
	if err != nil {
		slog.Error("Cannot create somes table", slog.Any("error", err))
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

	addr := cfg.GetAddr()
	slog.Info("Server listening on", slog.String("address", addr))

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
		os.Exit(1)
	}
}
