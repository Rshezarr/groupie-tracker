package app

import (
	"fmt"

	"groupie-tracker/internal/config"
	"groupie-tracker/internal/delivery"
	"groupie-tracker/internal/server"
)

type App struct {
	config *config.Config
}

func New(config config.Config) *App {
	return &App{
		config: &config,
	}
}

func (a App) Run() error {
	handler, err := delivery.NewHandler(a.config.TemplatePath)
	if err != nil {
		return err
	}
	server := server.NewServer(a.config, handler)

	fmt.Printf("Starting server at port %v...\nhttp://localhost%v\n", a.config.Addr, a.config.Addr)
	return server.Srv.ListenAndServe()
}
