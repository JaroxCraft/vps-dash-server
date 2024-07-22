package server

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/system"
	"net/http"
	"strconv"
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

	Server().HandleFunc("GET /processes/{pid}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/json; charset=utf-8")

		pidStr := request.PathValue("pid")
		if pidStr == "" {
			http.Error(writer, "Missing path param \"pid\"", http.StatusBadRequest)
			return
		}

		pid, err := strconv.ParseInt(pidStr, 0, 32)

		if err != nil {
			http.Error(writer, "Error parsing path param \"pid\"", http.StatusBadRequest|http.StatusInternalServerError)
			return
		}

		process := system.GetProcess(int32(pid))

		if process.Name == "" {
			http.Error(writer, "Process not found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(process)
		if err != nil {
			logger.Get().Warnf("Failed encoding process: %v", err)
			return
		}
	})
}
