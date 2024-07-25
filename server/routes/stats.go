package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/security"
	"github.com/jaroxcraft/vps-dash-server/system"
)

type Response struct {
	Usage uint8 `json:"usage"`
}

func usageHandler(server *http.ServeMux, path string, getCachedUsage func() uint8, getUsage func() uint8) {
	server.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if !security.ValidateRequest(*r) {
			http.Error(w, "Token invalid", http.StatusUnauthorized)

			return
		}

		response := Response{}

		cacheParam := r.URL.Query().Get("cache")
		if cacheParam == "" {
			cacheParam = "true"
		}

		cache, err := strconv.ParseBool(cacheParam)
		if err != nil {
			http.Error(w, "Cache param must be a boolean", http.StatusBadRequest)

			return
		}

		if cache {
			response.Usage = getCachedUsage()
		} else {
			response.Usage = getUsage()
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Get().Warnf("Failed encoding response: %v", err)

			return
		}
	})
}

func GetCPU(server *http.ServeMux) {
	usageHandler(server, "GET /stats/cpu/", system.GetCachedCPUUsage, system.GetCPUUsage)
}

func GetMemory(server *http.ServeMux) {
	usageHandler(server, "GET /stats/memory/", system.GetCachedMemoryUsage, system.GetMemoryUsage)
}
