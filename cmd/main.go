package main

import (
	"log"

	"groupie-tracker/internal/app"
	"groupie-tracker/internal/config"
)

func main() {
	config := config.NewConfig()

	app := app.New(*config)
	if err := app.Run(); err != nil {
		log.Printf("%v", err)
	}
}
