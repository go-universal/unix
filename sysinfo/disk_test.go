package sysinfo_test

import (
	"testing"

	"github.com/go-universal/unix/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestDisk(t *testing.T) {
	total, used, free, err := sysinfo.DiskInfo()
	assert.NoError(t, err)

	// Ensure total, used, and free are non-negative
	assert.GreaterOrEqual(t, total, uint64(0), "Total disk space should be non-negative")
	assert.GreaterOrEqual(t, used, uint64(0), "Used disk space should be non-negative")
	assert.GreaterOrEqual(t, free, uint64(0), "Free disk space should be non-negative")

	// Ensure total is at least the sum of used and free
	assert.Equal(t, total, used+free, "Total disk space should equal used + free")
}
