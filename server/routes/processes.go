package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/security"
	"github.com/jaroxcraft/vps-dash-server/system"
)

func GetProcesses(server *http.ServeMux) {
	server.HandleFunc("GET /processes/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		allowed := security.ValidateRequest(*r)
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
}

func GetProcess(server *http.ServeMux) {
	server.HandleFunc("GET /processes/{pid}/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		allowed := security.ValidateRequest(*r)
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
			http.Error(w, "Error parsing path param \"pid\"", http.StatusBadRequest)

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
