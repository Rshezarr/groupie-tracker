package config

import (
	"os"
	"time"
)

const (
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
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: oneMB,
		TemplatePath:   tempPath,
	}
}
