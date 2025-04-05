package sysinfo

import "github.com/shirou/gopsutil/v4/net"

// NetworkInfo retrieves the total bytes sent and received on the network interfaces.
// It returns the number of bytes sent, bytes received, and an error if any occurs.
func NetworkInfo() (sent, recv uint64, err error) {
	// Get network I/O statistics for all interfaces.
	stats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}

	// Get sent and received bytes if data is available.
	if len(stats) > 0 {
		sent = stats[0].BytesSent
		recv = stats[0].BytesRecv
	}

	return sent, recv, nil
}
