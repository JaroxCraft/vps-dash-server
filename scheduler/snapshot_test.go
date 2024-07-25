package scheduler_test

import (
	"github.com/jaroxcraft/vps-dash-server/scheduler"
	"testing"
	"time"
)

// Adds a new snapshot to the beginning of the snapshot array
func TestAddSnapshotAddsToBeginning(t *testing.T) {
	snapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: 50,
		MemPercentage: 30,
		MemUsed:       1024,
		MemTotal:      2048,
	}

	scheduler.AddSnapshot(snapshot)

	if scheduler.GetSnapshots()[0] != *snapshot {
		t.Errorf("Expected snapshot at the beginning of the array")
	}
}

// Shifts existing snapshots to the right by one position
func TestAddSnapshotShiftsExistingSnapshots(t *testing.T) {
	initialSnapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: 50,
		MemPercentage: 30,
		MemUsed:       1024,
		MemTotal:      2048,
	}

	scheduler.AddSnapshot(initialSnapshot)

	newSnapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now().Add(time.Minute),
		CPUPercentage: 60,
		MemPercentage: 40,
		MemUsed:       2048,
		MemTotal:      4096,
	}

	scheduler.AddSnapshot(newSnapshot)

	if scheduler.GetSnapshots()[1] != *initialSnapshot {
		t.Errorf("Expected initial snapshot to be shifted to the second position")
	}
}

// Maintains the fixed size of the snapshot array
func TestAddSnapshotMaintainsFixedSize(t *testing.T) {
	for i := 0; i < 10; i++ {
		snapshot := &scheduler.Snapshot{
			CaptureTime:   time.Now().Add(time.Duration(i) * time.Minute),
			CPUPercentage: uint8(i * 10),
			MemPercentage: uint8(i * 5),
			MemUsed:       uint64(i * 1000),
			MemTotal:      uint64(i * 2000),
		}
		scheduler.AddSnapshot(snapshot)
	}

	if len(scheduler.GetSnapshots()) != 8 {
		t.Errorf("Expected snapshots array to maintain fixed size of 8")
	}
}

// Handles the addition of a snapshot when the array is already full
func TestAddSnapshotWhenArrayIsFull(t *testing.T) {
	for i := 0; i < scheduler.SnapshotCount; i++ {
		snapshot := &scheduler.Snapshot{
			CaptureTime:   time.Now().Add(time.Duration(i) * time.Minute),
			CPUPercentage: uint8(i * 10),
			MemPercentage: uint8(i * 5),
			MemUsed:       uint64(i * 1000),
			MemTotal:      uint64(i * 2000),
		}
		scheduler.AddSnapshot(snapshot)
	}

	newSnapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now().Add(time.Hour),
		CPUPercentage: 90,
		MemPercentage: 45,
		MemUsed:       9000,
		MemTotal:      18000,
	}

	scheduler.AddSnapshot(newSnapshot)

	if scheduler.GetSnapshots()[0] != *newSnapshot {
		t.Errorf("Expected new snapshot at the beginning of the array")
	}
}

// Correctly shifts snapshots when the array contains only one snapshot
func TestAddSnapshotWithOneExistingSnapshot(t *testing.T) {
	initialSnapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: 50,
		MemPercentage: 30,
		MemUsed:       1024,
		MemTotal:      2048,
	}

	scheduler.AddSnapshot(initialSnapshot)

	newSnapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now().Add(time.Minute),
		CPUPercentage: 60,
		MemPercentage: 40,
		MemUsed:       2048,
		MemTotal:      4096,
	}

	scheduler.AddSnapshot(newSnapshot)

	if scheduler.GetSnapshots()[1] != *initialSnapshot {
		t.Errorf("Expected initial snapshot to be shifted to the second position")
	}
}

// Handles the addition of a snapshot when the array is empty
func TestAddSnapshotWhenArrayIsEmpty(t *testing.T) {
	snapshot := &scheduler.Snapshot{
		CaptureTime:   time.Now(),
		CPUPercentage: 50,
		MemPercentage: 30,
		MemUsed:       1024,
		MemTotal:      2048,
	}

	scheduler.AddSnapshot(snapshot)

	if scheduler.GetSnapshots()[0] != *snapshot {
		t.Errorf("Expected snapshot at the beginning of the array")
	}
}
