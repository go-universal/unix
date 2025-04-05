package sysinfo

import (
	"github.com/shirou/gopsutil/v4/disk"
)

// DiskInfo retrieves disk usage statistics for the root directory ("/").
// It returns the total, used, and free disk space in bytes, along with any error encountered.
func DiskInfo() (total, used, free uint64, err error) {
	// Get disk usage statistics for the root directory.
	stats, err := disk.Usage("/")
	if err != nil {
		return 0, 0, 0, err
	}

	// Return the total, used, and free disk space.
	return stats.Total, stats.Used, stats.Free, nil
}
