package system

import (
	"fmt"
	"math"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/cpu"
)

var cpuUsageCache int8 = -1

func cpuUsage() (float64, error) {
	// Get CPU usage percentages
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, fmt.Errorf("error getting cpu usage: %w", err)
	}

	if len(percentages) == 0 {
		return 0, fmt.Errorf("no cpu data available")
	}

	return percentages[0], nil
}

func GetCPUUsage() uint8 {
	percentage, err := cpuUsage()
	if err != nil {
		logger.Get().Errorf("Error while getting cpu usage: %v", err)
	}

	return uint8(math.Ceil(percentage))
}

func cacheCPUUsage(usage uint8) {
	cpuUsageCache = int8(usage)
}

func GetCachedCPUUsage() uint8 {
	if cpuUsageCache == -1 {
		cacheCPUUsage(GetCPUUsage())
	}

	return uint8(cpuUsageCache)
}
