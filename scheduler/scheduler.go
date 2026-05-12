package scheduler

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/jaroxcraft/vps-dash-server/colors"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/jaroxcraft/vps-dash-server/system"
	"github.com/madflojo/tasks"
)

const (
	snapshotTaskName = "snapshot"
	SnapshotInterval = 15 * time.Second
	cacheTaskName    = "cache"
	CacheInterval    = 5 * time.Second
)

func Start() *tasks.Scheduler {
	scheduler := tasks.New()

	err := scheduler.AddWithID(snapshotTaskName, &tasks.Task{
		StartAfter: time.Now(),
		Interval:   GetSnapshotInterval(),
		TaskFunc:   TakeSnapshot,
	})
	if err != nil {
		logger.Get().Fatalf("Error while Starting Snapshot Scheduler: %v", err)
	}

	logger.Get().Infof("Snapshot-Scheduler started with id %s", colors.S(color.FgYellow, snapshotTaskName))

	err = scheduler.AddWithID(cacheTaskName, &tasks.Task{
		StartAfter: time.Now(),
		Interval:   GetCacheInterval(),
		TaskFunc:   system.Cache,
	})
	if err != nil {
		logger.Get().Fatalf("Error while Starting Cache Scheduler: %v", err)
	}

	logger.Get().Infof("Cache-Scheduler started with id %s", colors.S(color.FgYellow, cacheTaskName))

	return scheduler
}

func GetCacheInterval() time.Duration {
	intervalEnv := os.Getenv("CACHE_INTERVAL")
	if intervalEnv == "" {
		return SnapshotInterval
	}

	// TODO: hint in documentation
	parsedInterval, err := time.ParseDuration(intervalEnv)
	if err != nil {
		logger.Get().Errorf("Failed to parse cache interval from ENV, falling back to default of %s: %v", CacheInterval.String(), err)

		return CacheInterval
	}

	return parsedInterval
}

func GetSnapshotInterval() time.Duration {
	intervalEnv := os.Getenv("SNAPSHOT_INTERVAL")
	if intervalEnv == "" {
		return SnapshotInterval
	}

	// TODO: hint in documentation
	parsedInterval, err := time.ParseDuration(intervalEnv)
	if err != nil {
		logger.Get().Errorf("Failed to parse snapshot interval from ENV, falling back to default of %s: %v", SnapshotInterval.String(), err)

		return SnapshotInterval
	}

	return parsedInterval
}
