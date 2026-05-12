package server

import (
	"net/http"

	"github.com/jaroxcraft/vps-dash-server/server/routes"
)

func RegisterRoutes(mux *http.ServeMux) {
	routes.GetSnapshots(mux)
	routes.GetProcesses(mux)
	routes.GetProcess(mux)
	routes.GetCPU(mux)
	routes.GetMemory(mux)
}
