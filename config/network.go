package config

import (
	"net/http"
	"time"
)

var (
	HTTPServerConfig struct {
		Addr    string
		Timeout time.Duration
	}

	HTTPClientConfig struct {
		Timeout time.Duration
	}
)

func InitHTTPServerConfig() {
	HTTPServerConfig.Addr = ":8080"
	HTTPServerConfig.Timeout = 10 * time.Second
}

func InitHTTPClientConfig() {
	HTTPClientConfig.Timeout = 10 * time.Second
}