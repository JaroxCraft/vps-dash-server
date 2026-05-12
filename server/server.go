package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jaroxcraft/vps-dash-server/logger"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 120 * time.Second
)

func New() *http.ServeMux {
	return http.NewServeMux()
}

func NewWithRoutes() *http.ServeMux {
	mux := New()
	RegisterRoutes(mux)

	return mux
}

func Start(mux *http.ServeMux) {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", GetAPIPort()),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	logger.Get().Infof("Starting server on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Get().Fatalf("Error while starting server: %v", err)
	}
}
