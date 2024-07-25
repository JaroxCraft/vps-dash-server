package scheduler_test

import (
	"os"
	"testing"
	"time"

	"github.com/jaroxcraft/vps-dash-server/scheduler"
)

// Returns default snapshot interval when SNAPSHOT_INTERVAL is not set.
func TestReturnsDefaultSnapshotIntervalWhenNotSet(t *testing.T) {
	err := os.Unsetenv("SNAPSHOT_INTERVAL")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Parses and returns the interval from SNAPSHOT_INTERVAL when it is set correctly.
func TestParsesAndReturnsIntervalWhenSetCorrectly(t *testing.T) {
	err := os.Setenv("SNAPSHOT_INTERVAL", "2m")
	if err != nil {
		t.Fatal(err)
	}

	expected, _ := time.ParseDuration("2m")

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Logs an error and returns default snapshot interval when SNAPSHOT_INTERVAL is set incorrectly.
func TestLogsErrorAndReturnsDefaultWhenSetIncorrectly(t *testing.T) {
	err := os.Setenv("SNAPSHOT_INTERVAL", "invalid")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// SNAPSHOT_INTERVAL is set to an empty string.
func TestSnapshotIntervalSetToEmptyString(t *testing.T) {
	err := os.Setenv("SNAPSHOT_INTERVAL", "")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// SNAPSHOT_INTERVAL is set to a non-duration string.
func TestSnapshotIntervalSetToNonDurationString(t *testing.T) {
	err := os.Setenv("SNAPSHOT_INTERVAL", "invalid")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// SNAPSHOT_INTERVAL is set to a very large duration value.
func TestSnapshotIntervalSetToVeryLargeDurationValue(t *testing.T) {
	err := os.Setenv("SNAPSHOT_INTERVAL", "10000h")
	if err != nil {
		t.Fatal(err)
	}

	expected, _ := time.ParseDuration("10000h")

	result := scheduler.GetSnapshotInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Returns default snapshot interval when CACHE_INTERVAL is not set.
func TestReturnsDefaultSnapshotIntervalWhenCacheIntervalNotSet(t *testing.T) {
	err := os.Unsetenv("CACHE_INTERVAL")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Parses and returns the interval from CACHE_INTERVAL when it is set correctly.
func TestParsesAndReturnsIntervalFromCacheIntervalWhenSetCorrectly(t *testing.T) {
	err := os.Setenv("CACHE_INTERVAL", "5s")
	if err != nil {
		t.Fatal(err)
	}

	expected := 5 * time.Second

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Logs an error and returns default cache interval when CACHE_INTERVAL is set incorrectly.
func TestLogsErrorAndReturnsDefaultCacheIntervalWhenCacheIntervalSetIncorrectly(t *testing.T) {
	err := os.Setenv("CACHE_INTERVAL", "invalid")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.CacheInterval

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// CACHE_INTERVAL is set to an empty string.
func TestCacheIntervalSetToEmptyString(t *testing.T) {
	err := os.Setenv("CACHE_INTERVAL", "")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.SnapshotInterval

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// CACHE_INTERVAL is set to an invalid duration format.
func TestCacheIntervalSetToInvalidDurationFormat(t *testing.T) {
	err := os.Setenv("CACHE_INTERVAL", "invalid")
	if err != nil {
		t.Fatal(err)
	}

	expected := scheduler.CacheInterval

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// CACHE_INTERVAL is set to a very large duration value.
func TestCacheIntervalSetToVeryLargeDurationValue(t *testing.T) {
	err := os.Setenv("CACHE_INTERVAL", "1000h")
	if err != nil {
		t.Fatal(err)
	}

	expected := 1000 * time.Hour

	result := scheduler.GetCacheInterval()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
