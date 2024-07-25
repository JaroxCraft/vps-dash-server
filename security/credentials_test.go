package security_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/jaroxcraft/vps-dash-server/security"
)

// Valid request with correct Bearer token returns true.
func TestValidRequestWithCorrectBearerToken(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer vps-dash")
	security.Initialize()

	if !security.ValidateRequest(*req) {
		t.Errorf("Expected true, got false")
	}
}

// Valid request with default token returns true when no AUTH_TOKEN ENV is set.
func TestValidRequestWithDefaultToken(t *testing.T) {
	err := os.Unsetenv("AUTH_TOKEN")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer vps-dash")
	security.Initialize()

	if !security.ValidateRequest(*req) {
		t.Errorf("Expected true, got false")
	}
}

// Request with no Authorization header returns false.
func TestRequestWithNoAuthorizationHeader(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	security.Initialize()

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but no Bearer token returns false.
func TestRequestWithAuthorizationHeaderButNoBearerToken(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Basic password")
	security.Initialize()

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but empty Bearer token returns false.
func TestRequestWithAuthorizationHeaderButEmptyBearerToken(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer ")
	security.Initialize()

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but incorrect token format returns false.
func TestRequestWithAuthorizationHeaderButIncorrectTokenFormat(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer1234")
	security.Initialize()

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}
