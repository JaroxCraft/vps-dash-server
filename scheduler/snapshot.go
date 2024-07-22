package scheduler

import (
	"encoding/json"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/system"
	"time"
)

var snapshots = make([]Snapshot, 8)

type Snapshot struct {
	CaptureTime   time.Time `json:"time"`
	CPUPercentage uint8     `json:"cpu_percent"`
	MemPercentage uint8     `json:"mem_percent"`
	MemTotal      uint64    `json:"mem_total"` // Bytes
	MemUsed       uint64    `json:"mem_used"`  // Bytes
}

func GetSnapshots() []Snapshot {
	return snapshots
}

func AddSnapshot(snapshot *Snapshot) {
	var newSnapshots = make([]Snapshot, 8)
	newSnapshots[0] = *snapshot

	for i, s := range snapshots {
		newIndex := i + 1

		if newIndex >= len(newSnapshots) {
			break
		}

		newSnapshots[newIndex] = s
	}
	snapshots = newSnapshots
}

func TakeSnapshot() error {

	var snapshot = Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: system.GetCPUUsage(),
		MemPercentage: system.MemoryUsage(),
		MemTotal:      system.Memory().Total,
		MemUsed:       system.Memory().Used,
	}

	jsonSnapshot, _ := json.Marshal(snapshot)
	logger.Get().Debugf("Snapshot taken: %s", string(jsonSnapshot))

	AddSnapshot(&snapshot)
	return nil
}
