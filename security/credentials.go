package security

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/jaroxcraft/vps-dash-server/logger"
)

var (
	errNoAuth = errors.New("no AUTH_TOKEN ENV")
	authToken = ""
)

const defaultAuthToken = "vps-dash"

func Initialize() {
	authToken = callAuthToken()
}

func GetAuthToken() string {
	if authToken == "" {
		authToken = callAuthToken()
	}

	return authToken
}

func callAuthToken() string {
	env, err := getAuthEnv()
	if err != nil && errors.Is(err, errNoAuth) {
		logger.Get().Warnf("no auth ENV set, using default of %s", defaultAuthToken)

		return defaultAuthToken
	}

	return env
}

func getAuthEnv() (string, error) {
	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		return "", errNoAuth
	}

	return authToken, nil
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
