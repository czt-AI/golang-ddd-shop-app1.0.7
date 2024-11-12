package network

import (
	"net/http"
	"time"
)

type Config struct {
	HTTPServer struct {
		Addr    string
		Timeout time.Duration
	}

	HTTPClient struct {
		Timeout time.Duration
	}

	LoadBalancer struct {
		Algorithm string
		Addr      string
	}

	CDN struct {
		Enabled bool
		Addr    string
	}
}

func NewConfig() *Config {
	return &Config{
		HTTPServer: struct {
			Addr    string
			Timeout time.Duration
		}{
			Addr:    ":8080",
			Timeout: 10 * time.Second,
		},
		HTTPClient: struct {
			Timeout time.Duration
		}{
			Timeout: 10 * time.Second,
		},
		LoadBalancer: struct {
			Algorithm string
			Addr      string
		}{
			Algorithm: "round_robin",
			Addr:      "localhost:8081",
		},
		CDN: struct {
			Enabled bool
			Addr    string
		}{
			Enabled: true,
			Addr:    "cdn.example.com",
		},
	}
}