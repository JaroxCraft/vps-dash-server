package scheduler

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/system"
	"time"
)

var snapshots = make([]Snapshot, 8)

type Snapshot struct {
	CaptureTime   time.Time `json:"time"`
	CpuPercentage uint8     `json:"cpu"`
	MemPercentage uint8     `json:"mem"`
	MemUsed       uint64    `json:"memUsed"` // MB
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

	logger.Get().Debug("taking snapshot")

	// TODO: Get real values
	var snapshot = Snapshot{
		CaptureTime:   time.Now(),
		CpuPercentage: system.GetCpuUsage(),
		MemPercentage: system.MemoryPercent(),
		MemUsed:       system.Memory().Used,
	}

	AddSnapshot(&snapshot)
	return nil
}
