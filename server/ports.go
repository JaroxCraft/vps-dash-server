package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/jaroxcraft/vps-dash-server/colors"
	"github.com/jaroxcraft/vps-dash-server/logger"
)

const DefaultPort = 8080

func GetAPIPort() uint16 {
	port, err := getEnvPort()
	if err == nil {
		logger.Get().Infof("Found valid API_PORT ENV: %s, using it", colors.S(color.FgBlue, strconv.Itoa(int(port))))

		return port
	}

	logger.Get().Warnf("%v, starting on default port %s", err, colors.S(color.FgBlue, strconv.Itoa(DefaultPort)))
	// TODO: link to docs/reference
	return DefaultPort
}

func getEnvPort() (uint16, error) {
	envPort := os.Getenv("API_PORT")
	if envPort == "" {
		return 0, fmt.Errorf("no API_PORT ENV")
	}

	port, err := strconv.ParseUint(envPort, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid API_PORT ENV")
	}

	return uint16(port), nil
}
