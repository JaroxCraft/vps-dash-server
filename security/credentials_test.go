package security_test

import (
	"net/http"
	"testing"

	"github.com/jaroxcraft/vps-dash-server/security"
)

// Valid request with correct Bearer token returns true.
func TestValidRequestWithCorrectBearerToken(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "vps-dash")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer vps-dash")

	if err := security.Initialize(); err != nil {
		t.Fatal(err)
	}

	if !security.ValidateRequest(*req) {
		t.Errorf("Expected true, got false")
	}
}

// Request with no Authorization header returns false.
func TestRequestWithNoAuthorizationHeader(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "test-token")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := security.Initialize(); err != nil {
		t.Fatal(err)
	}

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but no Bearer token returns false.
func TestRequestWithAuthorizationHeaderButNoBearerToken(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "test-token")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Basic password")

	if err := security.Initialize(); err != nil {
		t.Fatal(err)
	}

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but empty Bearer token returns false.
func TestRequestWithAuthorizationHeaderButEmptyBearerToken(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "test-token")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer ")

	if err := security.Initialize(); err != nil {
		t.Fatal(err)
	}

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Request with Authorization header but incorrect token format returns false.
func TestRequestWithAuthorizationHeaderButIncorrectTokenFormat(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "test-token")

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer1234")

	if err := security.Initialize(); err != nil {
		t.Fatal(err)
	}

	if security.ValidateRequest(*req) {
		t.Errorf("Expected false, got true")
	}
}

// Initialize returns an error when AUTH_TOKEN is not set.
func TestInitializeReturnsErrorWhenNoAuthToken(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "")

	if err := security.Initialize(); err == nil {
		t.Errorf("Expected error when AUTH_TOKEN is not set, got nil")
	}
}
