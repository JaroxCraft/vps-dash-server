package server

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/security"
	"github.com/jaroxcraft/vps-dash-server/system"
	"net/http"
	"strconv"
)

func RegisterRoutes() {
	Server().HandleFunc("GET /snapshots/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		allowed := security.ValidateToken(*r)
		if !allowed {
			http.Error(w, "Token invalid", http.StatusUnauthorized)
			return
		}

		err := json.NewEncoder(w).Encode(scheduler.GetSnapshots())
		if err != nil {
			logger.Get().Warnf("Failed encoding snapshots: %v", err)
			return
		}
	})

	Server().HandleFunc("GET /processes/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		allowed := security.ValidateToken(*r)
		if !allowed {
			http.Error(w, "Token invalid", http.StatusUnauthorized)
			return
		}

		err := json.NewEncoder(w).Encode(system.GetProcesses())
		if err != nil {
			logger.Get().Warnf("Failed encoding processes: %v", err)
			return
		}
	})

	Server().HandleFunc("GET /processes/{pid}/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		allowed := security.ValidateToken(*r)
		if !allowed {
			http.Error(w, "Token invalid", http.StatusUnauthorized)
			return
		}

		pidStr := r.PathValue("pid")
		if pidStr == "" {
			http.Error(w, "Missing path param \"pid\"", http.StatusBadRequest)
			return
		}

		pid, err := strconv.ParseInt(pidStr, 0, 32)

		if err != nil {
			http.Error(w, "Error parsing path param \"pid\"", http.StatusBadRequest|http.StatusInternalServerError)
			return
		}

		process, err := system.GetProcess(int32(pid))
		if err != nil {
			http.Error(w, "Process not found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(process)
		if err != nil {
			logger.Get().Warnf("Failed encoding process: %v", err)
			return
		}
	})
}
