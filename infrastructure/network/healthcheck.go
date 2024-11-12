package network

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func SetupHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", HealthCheckHandler()).Methods("GET")
}