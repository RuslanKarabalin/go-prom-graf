package main

import (
	"log/slog"
	"os"

	"example/internal/app"
)

func main() {
	l := slog.Default()
	a, err := app.New()
	if err != nil {
		l.Error("Can't create App", "err", err)
		os.Exit(1)
	}
	err = a.Run()
	if err != nil {
		l.Error("Can't run App", "err", err)
		os.Exit(1)
	}
}
