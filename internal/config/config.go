package config

import "time"

const (
	port     = ":8090"
	oneMB    = 1 << 20
	timeout  = 10 * time.Second
	tempPath = "ui/*.html"
)

type Config struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	TemplatePath   string
}

func NewConfig() *Config {
	return &Config{
		Addr:           port,
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: oneMB,
		TemplatePath:   tempPath,
	}
}
