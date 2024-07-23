package system

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/mem"
)

func Memory() *mem.VirtualMemoryStat {
	memory, err := mem.VirtualMemory()
	if err != nil {
		logger.Get().Errorf("Error getting memory info: %v", err)
	}

	return memory
}

func MemoryUsage() uint8 {
	return uint8(Memory().UsedPercent)
}
