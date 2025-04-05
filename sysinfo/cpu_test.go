package sysinfo_test

import (
	"testing"

	"github.com/go-universal/unix/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestCPU(t *testing.T) {
	cores, used, free, err := sysinfo.CPUInfo()
	assert.NoError(t, err)

	// Assert that the number of cores is greater than 0
	assert.Greater(t, cores, 0, "Number of CPU cores should be greater than 0")

	// Assert that used and free percentages are within valid ranges
	assert.GreaterOrEqual(t, used, 0.0, "Used CPU percentage should be >= 0")
	assert.LessOrEqual(t, used, 100.0, "Used CPU percentage should be <= 100")
	assert.GreaterOrEqual(t, free, 0.0, "Free CPU percentage should be >= 0")
	assert.LessOrEqual(t, free, 100.0, "Free CPU percentage should be <= 100")

	// Assert that used + free is approximately 100 (allowing for rounding errors)
	assert.InDelta(t, 100.0, used+free, 0.1, "Used + Free CPU percentage should be approximately 100")
}
