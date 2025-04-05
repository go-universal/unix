package sysinfo

import (
	"github.com/shirou/gopsutil/v4/mem"
)

// MemoryInfo retrieves the system's memory statistics.
// It returns the total, used, and free memory in bytes, along with any error encountered.
func MemoryInfo() (total, used, free uint64, err error) {
	// Get virtual memory statistics using gopsutil
	stats, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}

	// Extract relevant memory statistics
	return stats.Total, stats.Used, stats.Free, nil
}
