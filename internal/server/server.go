package server

import (
	"net/http"

	"groupie-tracker/internal/config"
	"groupie-tracker/internal/delivery"
)

type Server struct {
	Srv *http.Server
}

func NewServer(config *config.Config, handler *delivery.Handler) *Server {
	mux := http.NewServeMux()
	handler.SetEndpoints(mux)
	return &Server{
		Srv: &http.Server{
			Addr:           config.Addr,
			Handler:        mux,
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
			MaxHeaderBytes: config.MaxHeaderBytes,
		},
	}
}
