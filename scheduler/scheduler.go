package scheduler

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/madflojo/tasks"
	"os"
	"time"
)

const snapshotInterval = 10 * time.Second

func Start() *tasks.Scheduler {
	scheduler := tasks.New()

	id, err := scheduler.Add(&tasks.Task{
		StartAfter: time.Now(),
		Interval:   GetSnapshotInterval(),
		TaskFunc:   TakeSnapshot,
	})
	if err != nil {
		logger.Get().Fatalf("Error while Starting Snapshot Scheduler: %v", err)
	}

	logger.Get().Infof("Snapshot Scheduler started with id %s", id)

	return scheduler
}

func GetSnapshotInterval() time.Duration {
	intervalEnv := os.Getenv("SNAPSHOT_INTERVAL")
	if intervalEnv == "" {
		return snapshotInterval
	}

	// TODO: hint in documentation
	parsedInterval, err := time.ParseDuration(intervalEnv)
	if err != nil {
		logger.Get().Errorf("Failed to parse snapshot interval from env, falling back to default of %s: %v", snapshotInterval.String(), err)
		return snapshotInterval
	}
	return parsedInterval
}
