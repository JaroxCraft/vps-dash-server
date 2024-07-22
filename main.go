package main

import (
	"github.com/jaroxcraft/vps-dash-server/banner"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"github.com/jaroxcraft/vps-dash-server/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	banner.PrintBanner()

	defer logger.Get().Sync()

	tasks := scheduler.Start()
	defer tasks.Stop()

	server.NewWithRoutes()
	server.Start()
}
