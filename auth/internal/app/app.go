package app

import (
	"context"
	"example/internal/config"
	"fmt"
	"time"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type App struct {
	Logger   *zap.Logger
	Config   *config.Config
	Postgres *pgxpool.Pool
	Fiber    *fiber.App
}

func New() (*App, error) {
	config := config.ReadConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("can't create zap logger %s", err)
	}

	poolConfig, err := pgxpool.ParseConfig(config.PostgresConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}
	poolConfig.MaxConns = 8
	poolConfig.MinConns = 2

	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pgPool.Ping(pingCtx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	goose.SetLogger(zap.NewStdLog(logger))

	fiber := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
		},
	)

	fiber.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	a := &App{
		Config:   config,
		Logger:   logger,
		Postgres: pgPool,
		Fiber:    fiber,
	}

	a.registerRoutes()

	return a, nil
}

func (a *App) Run() error {
	return a.Fiber.Listen(a.Config.Addr)
}
