package system

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/mem"
)

var memoryUsageCache int8 = -1

func Memory() *mem.VirtualMemoryStat {
	memory, err := mem.VirtualMemory()
	if err != nil {
		logger.Get().Errorf("Error getting memory info: %v", err)
	}

	return memory
}

func GetMemoryUsage() uint8 {
	return uint8(Memory().UsedPercent)
}

func cacheMemoryUsage(usage uint8) {
	memoryUsageCache = int8(usage)
}

func GetCachedMemoryUsage() uint8 {
	if memoryUsageCache == -1 {
		cacheMemoryUsage(GetMemoryUsage())
	}

	return uint8(memoryUsageCache)
}
