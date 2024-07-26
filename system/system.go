package system

import "github.com/jaroxcraft/vps-dash-server/logger"

func Cache() error {
	logger.Get().Debug("Updating Cache")

	cacheCPUUsage(GetCPUUsage())
	cacheMemoryUsage(GetMemoryUsage())

	return nil
}
