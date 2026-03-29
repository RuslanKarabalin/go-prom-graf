package main

import (
	"log"
	"os"

	"example/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("Can't create App", err)
		os.Exit(1)
	}
	err = a.Run()
	if err != nil {
		log.Fatal("Can't run App", err)
		os.Exit(1)
	}
}
