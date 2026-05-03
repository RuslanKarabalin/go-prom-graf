package app

import (
	"context"
	"example/internal/config"
	"example/internal/db"
	"fmt"
	"time"

	middleware "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
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
	if err := db.RunMigrations(pgPool); err != nil {
		return nil, err
	}

	fiber := fiber.New()

	fiber.Use(middleware.New(middleware.Config{
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
	return a.Fiber.Listen(a.Config.Addr, fiber.ListenConfig{DisableStartupMessage: true})
}
