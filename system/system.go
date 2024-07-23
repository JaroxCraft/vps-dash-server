package system

import "github.com/jaroxcraft/vps-dash-server/logger"

func Cache() error {
	logger.Get().Debug("Updating Cache")

	cacheCPUUsage(GetCPUUsage())

	return nil
}
