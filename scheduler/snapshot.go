package scheduler

import (
	"encoding/json"
	"time"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/system"
)

const (
	snapshotCount = 8
)

var snapshots = make([]Snapshot, snapshotCount)

type Snapshot struct {
	CaptureTime   time.Time `json:"time"`
	CPUPercentage uint8     `json:"cpuPercent"`
	MemPercentage uint8     `json:"memPercentage"`
	MemUsed       uint64    `json:"memUsed"`  // Bytes
	MemTotal      uint64    `json:"memTotal"` // Bytes
}

func GetSnapshots() []Snapshot {
	return snapshots
}

func AddSnapshot(snapshot *Snapshot) {
	newSnapshots := make([]Snapshot, snapshotCount)
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
	snapshot := Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: system.GetCPUUsage(),
		MemPercentage: system.MemoryUsage(),
		MemTotal:      system.Memory().Total,
		MemUsed:       system.Memory().Used,
	}

	jsonSnapshot, _ := json.Marshal(snapshot)
	logger.Get().Debugf("Took Snapshot: %s", string(jsonSnapshot))

	AddSnapshot(&snapshot)

	return nil
}
