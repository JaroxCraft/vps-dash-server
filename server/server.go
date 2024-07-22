package server

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"net/http"
)

var server *http.ServeMux

func New() *http.ServeMux {
	server = http.NewServeMux()
	return server
}

func Server() *http.ServeMux {
	if server == nil {
		return New()
	}
	return server
}

func NewWithRoutes() *http.ServeMux {
	New()
	RegisterRoutes()
	return server
}

func Start() {
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		logger.Get().Fatalf("Error while starting server: %s", err)
	}
}
