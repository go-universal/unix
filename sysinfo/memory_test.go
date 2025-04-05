package sysinfo_test

import (
	"testing"

	"github.com/go-universal/unix/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	total, used, free, err := sysinfo.MemoryInfo()
	assert.NoError(t, err)

	// Ensure total memory is greater than zero
	assert.Greater(t, total, uint64(0), "Total memory should be greater than zero")

	// Ensure used memory is less than or equal to total memory
	assert.LessOrEqual(t, used, total, "Used memory should be less than or equal to total memory")

	// Ensure free memory is less than or equal to total memory
	assert.LessOrEqual(t, free, total, "Free memory should be less than or equal to total memory")

	// Ensure total memory equals used + free
	assert.Equal(t, total, used+free, "Total memory should equal used + free memory")
}
