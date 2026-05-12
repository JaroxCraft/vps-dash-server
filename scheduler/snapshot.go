package scheduler

import (
	"encoding/json"
	"time"

	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/system"
)

const (
	SnapshotCount = 8
)

var snapshots = make([]Snapshot, 0, SnapshotCount)

type Snapshot struct {
	CaptureTime   time.Time `json:"time"`
	CPUPercentage uint8     `json:"cpuPercent"`
	MemPercentage uint8     `json:"memPercent"`
	MemUsed       uint64    `json:"memUsed"`  // Bytes
	MemTotal      uint64    `json:"memTotal"` // Bytes
}

func GetSnapshots() []Snapshot {
	return snapshots
}

func AddSnapshot(snapshot *Snapshot) {
	if len(snapshots) < SnapshotCount {
		snapshots = append(snapshots, Snapshot{})
		copy(snapshots[1:], snapshots[0:len(snapshots)-1])
	} else {
		copy(snapshots[1:], snapshots[0:SnapshotCount-1])
	}
	snapshots[0] = *snapshot
}

func TakeSnapshot() error {
	mem := system.Memory()
	snapshot := Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: system.GetCPUUsage(),
		MemPercentage: system.GetCachedMemoryUsage(),
		MemTotal:      mem.Total,
		MemUsed:       mem.Used,
	}

	jsonSnapshot, err := json.Marshal(snapshot)
	if err != nil {
		logger.Get().Warnf("Failed marshaling snapshot: %v", err)
	} else {
		logger.Get().Debugf("Took Snapshot: %s", string(jsonSnapshot))
	}

	AddSnapshot(&snapshot)

	return nil
}
