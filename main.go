package main

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/server"
	"github.com/jaroxcraft/vps-dash-server/system"
	"log"
	"net/http"
)

func main() {
	defer logger.Get().Sync()

	tasks := scheduler.Start()
	defer tasks.Stop()

	system.GetCpuUsage()

	server.RegisterRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
