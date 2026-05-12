package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jaroxcraft/vps-dash-server/logger"
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
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Get().Infof("Starting server on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		logger.Get().Fatalf("Error while starting server: %v", err)
	}
}
