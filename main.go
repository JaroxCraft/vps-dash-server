package main

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/server"
	"github.com/jaroxcraft/vps-dash-server/system"
)

func main() {
	defer logger.Get().Sync()

	tasks := scheduler.Start()
	defer tasks.Stop()

	system.GetCpuUsage()

	server.NewWithRoutes()
	server.Start()
}
