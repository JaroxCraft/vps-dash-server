package security

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/jaroxcraft/vps-dash-server/logger"
)

var (
	errNoAuth = errors.New("no AUTH_TOKEN environment variable set")
	authToken = ""
)

func Initialize() error {
	token, err := getAuthEnv()
	if err != nil {
		return err
	}

	authToken = token
	logger.Get().Info("Security initialized with AUTH_TOKEN")

	return nil
}

func GetAuthToken() string {
	return authToken
}

func getAuthEnv() (string, error) {
	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		return "", errNoAuth
	}

	return token, nil
}

func ValidateRequest(r http.Request) bool {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return false
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 { //nolint:mnd
		return false
	}

	if authHeaderParts[0] != "Bearer" {
		return false
	}

	tokenString := authHeaderParts[1]
	if tokenString == "" {
		return false
	}

	return tokenString == GetAuthToken()
}
