package server

import (
	"fmt"
	"github.com/jaroxcraft/vps-dash-server/logger"
	"os"
	"strconv"
)

const defaultPort = 8080

func GetAPIPort() uint16 {
	port, err := getEnvPort()
	if err == nil {
		logger.Get().Infof("Using port %d from env", port)
		return port
	}
	logger.Get().Infof("API_PORT not set in env, using default port %d", defaultPort)
	return defaultPort
}

func getEnvPort() (uint16, error) {
	envPort := os.Getenv("API_PORT")
	if envPort == "" {
		return 0, fmt.Errorf("no API_PORT environment variable set")
	}
	port, err := strconv.ParseUint(envPort, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid API_PORT environment variable set")
	}
	return uint16(port), nil
}
