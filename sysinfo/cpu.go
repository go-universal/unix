package sysinfo

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
)

// CPUInfo retrieves the number of CPU cores, the CPU usage percentage,
// and calculates the used and free CPU percentages.
// It returns the number of cores, used percentage, free percentage,
// and an error if any occurs.
func CPUInfo() (cores int, used, free float64, err error) {
	// Get the number of CPU cores available on the system.
	cores = runtime.NumCPU()

	// Get the CPU usage percentage for all CPUs combined (non-blocking call).
	// The first argument (0) specifies the interval in seconds (0 means immediate).
	// The second argument (false) specifies whether to retrieve per-CPU stats (false means combined stats).
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return cores, 0, 0, err
	}

	// Calculate used and free CPU percentages if data is available.
	if len(percent) > 0 {
		used = percent[0]
		free = 100 - used
	}

	return cores, used, free, nil
}
