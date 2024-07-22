package system

import (
	"github.com/jaroxcraft/vps-dash-server/logger"
	"github.com/shirou/gopsutil/process"
)

type Process struct {
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
}

func getPSProcesses() []*process.Process {
	processes, err := process.Processes() // TODO: array of pointers?
	if err != nil {
		logger.Get().Errorf("Error while fetching processes: %s", err)
	}
	return processes
}

func GetProcesses() []Process {
	psProcesses := getPSProcesses()
	processes := make([]Process, len(psProcesses))

	for i, psProcess := range psProcesses {
		psName, _ := psProcess.Name()

		processes[i] = Process{
			Pid:  psProcess.Pid,
			Name: psName,
		}
	}

	return processes
}
