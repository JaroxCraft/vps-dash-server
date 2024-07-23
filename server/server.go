package server

import (
	"fmt"
	"net/http"

	"github.com/jaroxcraft/vps-dash-server/logger"
)

var server *http.ServeMux

func New() *http.ServeMux {
	server = http.NewServeMux()

	return server
}

func NewWithRoutes() *http.ServeMux {
	New()
	RegisterRoutes()

	return server
}

func Start() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", GetAPIPort()), server)
	if err != nil {
		logger.Get().Fatalf("Error while starting server: %v", err)
	}
}
