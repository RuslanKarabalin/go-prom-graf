package app

import (
	"example/internal/config"
	"fmt"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
	Fiber  *fiber.App
}

func New() (*App, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("can't create zap logger %s", err)
	}

	config := config.ReadConfig()

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
		Logger: logger,
		Config: config,
		Fiber:  fiber,
	}

	a.registerRoutes()

	return a, nil
}

func (a *App) Run() error {
	return a.Fiber.Listen(a.Config.Addr)
}
