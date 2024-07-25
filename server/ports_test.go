package server_test

import (
	"os"
	"testing"

	"github.com/jaroxcraft/vps-dash-server/server"
)

// Returns port from API_PORT environment variable when it is valid.
func TestReturnsPortFromValidEnv(t *testing.T) {
	err := os.Setenv("API_PORT", "9090")
	if err != nil {
		t.Fatal(err)
	}

	port := server.GetAPIPort()

	if port != 9090 {
		t.Errorf("Expected port 9090, but got %d", port)
	}
}

// Returns default port when API_PORT environment variable is not set.
func TestReturnsDefaultPortWhenEnvNotSet(t *testing.T) {
	err := os.Unsetenv("API_PORT")
	if err != nil {
		t.Fatal(err)
	}

	port := server.GetAPIPort()

	if port != server.DefaultPort {
		t.Errorf("Expected default port %d, but got %d", server.DefaultPort, port)
	}
}

// API_PORT environment variable is set to the maximum valid port number (65535).
func TestMaxValidEnvPort(t *testing.T) {
	err := os.Setenv("API_PORT", "65535")
	if err != nil {
		t.Fatal(err)
	}

	port := server.GetAPIPort()

	if port != 65535 {
		t.Errorf("Expected port 65535, but got %d", port)
	}
}

// API_PORT environment variable is set to the minimum valid port number (0).
func TestMinValidEnvPort(t *testing.T) {
	err := os.Setenv("API_PORT", "0")
	if err != nil {
		t.Fatal(err)
	}

	port := server.GetAPIPort()

	if port != 0 {
		t.Errorf("Expected port 0, but got %d", port)
	}
}

// API_PORT environment variable is set to a non-numeric string.
func TestNonNumericEnvPort(t *testing.T) {
	err := os.Setenv("API_PORT", "invalid")
	if err != nil {
		t.Fatal(err)
	}

	port := server.GetAPIPort()

	if port != server.DefaultPort {
		t.Errorf("Expected default port %d, but got %d", server.DefaultPort, port)
	}
}
