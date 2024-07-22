package security

import (
	"errors"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"net/http"
	"os"
	"strings"
)

var noAuthEnvError = errors.New("no auth ENV")
var authToken = ""

const defaultAuthToken = "vps-dash"

func GetAuthToken() string {
	if authToken == "" {
		authToken = CallAuthToken()
	}
	return authToken
}

func CallAuthToken() string {
	env, err := getAuthEnv()
	if err != nil && errors.Is(err, noAuthEnvError) {
		logger.Get().Warnf("No auth ENV set, using default of %s", defaultAuthToken)
		return defaultAuthToken
	}
	return env
}

func getAuthEnv() (string, error) {
	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		return "", noAuthEnvError
	}
	return authToken, nil
}

func ValidateToken(r http.Request) bool {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return false
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
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
