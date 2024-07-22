package server

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"net/http"
)

func RegisterRoutes() {
	Server().HandleFunc("GET /snapshots", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/json; charset=utf-8")
		err := json.NewEncoder(writer).Encode(scheduler.GetSnapshots())
		if err != nil {
			logger.Get().Warnf("Failed encoding snapshots: %s", err)
			return
		}
	})
}
