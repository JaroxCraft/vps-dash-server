package system

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/cpu"
	"math"
)

func getCPUUsageF() (float64, error) {
	times, err := cpu.Times(false)
	if err != nil {
		return -1, err
	}
	logger.Get().Debugf("times: %v", times)
	usage := calculateCPUUsage(times[0])
	return usage, nil
}

func calculateCPUUsage(stat cpu.TimesStat) float64 {
	total := stat.User + stat.System + stat.Idle + stat.Nice + stat.Iowait + stat.Irq + stat.Softirq + stat.Steal + stat.Guest + stat.GuestNice
	usage := ((total - stat.Idle) / total) * 100
	return usage
}

func GetCpuUsage() uint8 {
	usageF, err := getCPUUsageF()
	if err != nil {
		logger.Get().Errorf("could not retrieve CPU usage: %s", err)
	}
	logger.Get().Debugf("CPU usage float: %f", usageF)
	return uint8(math.Ceil(usageF))
}
