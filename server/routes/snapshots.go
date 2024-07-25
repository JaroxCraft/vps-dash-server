package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/security"
)

func GetSnapshots(server *http.ServeMux) {
	server.HandleFunc("GET /snapshots/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		allowed := security.ValidateRequest(*r)

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

	server.HandleFunc("OPTIONS /snapshots/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	})
}
