package server

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/snapshots", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "text/json")

		// TODO: write custom marshaller for better control
		err := json.NewEncoder(writer).Encode(scheduler.GetSnapshots())
		if err != nil {
			logger.Get().Errorf("Error encoding snapshots: %s", err)
		}
	})
}
