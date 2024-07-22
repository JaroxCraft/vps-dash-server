package system

import (
	"fmt"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/cpu"
	"math"
	"time"
)

func cpuUsage() (float64, error) {
	// Get CPU usage percentages
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, fmt.Errorf("error getting cpu usage: %w", err)
	}

	if len(percentages) == 0 {
		return 0, fmt.Errorf("no cpu data available")
	}
	return percentages[0], nil
}

func GetCpuUsage() uint8 {
	percentage, err := cpuUsage()
	if err != nil {
		logger.Get().Errorf("Error while getting cpu usage: %v", err)
	}
	logger.Get().Infof("CPU usage is %.2f", percentage)
	return uint8(math.Ceil(percentage))
}
