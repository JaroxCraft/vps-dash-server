package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/security"
	"github.com/jaroxcraft/vps-dash-server/system"
)

func GetCPU(server *http.ServeMux) {
	type Response struct {
		Usage uint8 `json:"usage"`
	}

	server.HandleFunc("GET /stats/cpu/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		allowed := security.ValidateRequest(*r)
		if !allowed {
			http.Error(w, "Token invalid", http.StatusUnauthorized)

			return
		}

		response := Response{system.GetCachedCPUUsage()}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Get().Warnf("Failed encoding response: %v", err)

			return
		}
	})
}

func GetMemory(server *http.ServeMux) {
	type Response struct {
		Usage uint8 `json:"usage"`
	}

	server.HandleFunc("GET /stats/memory/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		allowed := security.ValidateRequest(*r)
		if !allowed {
			http.Error(w, "Token invalid", http.StatusUnauthorized)

			return
		}

		response := Response{system.MemoryUsage()}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Get().Warnf("Failed encoding response: %v", err)

			return
		}
	})
}
