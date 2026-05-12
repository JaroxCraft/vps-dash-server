package server_test

import (
	"testing"

	"github.com/jaroxcraft/vps-dash-server/server"
)

// Returns port from API_PORT environment variable when it is valid.
func TestReturnsPortFromValidEnv(t *testing.T) {
	t.Setenv("API_PORT", "9090")

	port := server.GetAPIPort()
	if port != 9090 {
		t.Errorf("Expected port 9090, but got %d", port)
	}
}

// Returns default port when API_PORT environment variable is not set.
func TestReturnsDefaultPortWhenEnvNotSet(t *testing.T) {
	t.Setenv("API_PORT", "")

	port := server.GetAPIPort()
	if port != server.DefaultPort {
		t.Errorf("Expected default port %d, but got %d", server.DefaultPort, port)
	}
}

// API_PORT environment variable is set to the maximum valid port number (65535).
func TestMaxValidEnvPort(t *testing.T) {
	t.Setenv("API_PORT", "65535")

	port := server.GetAPIPort()
	if port != 65535 {
		t.Errorf("Expected port 65535, but got %d", port)
	}
}

// API_PORT environment variable is set to 0, which falls back to the default.
func TestZeroEnvPortFallsBackToDefault(t *testing.T) {
	t.Setenv("API_PORT", "0")

	port := server.GetAPIPort()
	if port != server.DefaultPort {
		t.Errorf("Expected default port %d, but got %d", server.DefaultPort, port)
	}
}

// API_PORT environment variable is set to a non-numeric string.
func TestNonNumericEnvPort(t *testing.T) {
	t.Setenv("API_PORT", "invalid")

	port := server.GetAPIPort()
	if port != server.DefaultPort {
		t.Errorf("Expected default port %d, but got %d", server.DefaultPort, port)
	}
}
