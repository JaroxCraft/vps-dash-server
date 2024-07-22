package server

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/system"
	"net/http"
)

func RegisterRoutes() {
	Server().HandleFunc("GET /snapshots", func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "text/json; charset=utf-8")
		err := json.NewEncoder(writer).Encode(scheduler.GetSnapshots())
		if err != nil {
			logger.Get().Warnf("Failed encoding snapshots: %v", err)
			return
		}
	})

	Server().HandleFunc("GET /processes", func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Content-Type", "text/json; charset=utf-8")
		err := json.NewEncoder(writer).Encode(system.GetProcesses())
		if err != nil {
			logger.Get().Warnf("Failed encoding processes: %v", err)
			return
		}
	})
}
