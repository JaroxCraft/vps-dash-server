package server

import "github.com/jaroxcraft/vps-dash-server/server/routes"

func RegisterRoutes() {
	routes.GetSnapshots(server)
	routes.GetProcesses(server)
	routes.GetProcess(server)
}
