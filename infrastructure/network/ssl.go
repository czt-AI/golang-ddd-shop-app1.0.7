package network

import (
	"crypto/tls"
	"net/http"
)

func SetupHTTPS(server *http.Server) {
	tlsConfig := &tls.Config{
		// Configure TLS settings here
	}

	server.TLSConfig = tlsConfig
}