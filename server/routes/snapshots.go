package routes

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/security"
	"net/http"
)

func GetSnapshots(server *http.ServeMux) {
	server.HandleFunc("GET /snapshots/", func(w http.ResponseWriter, r *http.Request) {
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
}
