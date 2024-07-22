package main

import (
	"github.com/jaroxcraft/vps-dash-server/banner"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/security"
	"github.com/jaroxcraft/vps-dash-server/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	banner.PrintBanner()

	defer func() {
		if logger.Get().Sync() != nil {
			logger.Get().Warnf("Failed Syncing logger: %v", logger.Get().Sync())
		}
	}()

	tasks := scheduler.Start()
	defer tasks.Stop()

	security.CallAuthToken()

	server.NewWithRoutes()
	server.Start()
}
