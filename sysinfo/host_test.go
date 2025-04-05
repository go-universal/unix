package sysinfo_test

import (
	"testing"

	"github.com/go-universal/unix/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestUptime(t *testing.T) {
	// Test Uptime function
	duration, err := sysinfo.Uptime()
	assert.NoError(t, err)
	assert.Greater(t, duration.Seconds(), 0.0, "Uptime should be greater than 0")

	// Test UptimeParts function
	days, hours, minutes, err := sysinfo.UptimeParts()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, days, uint64(0), "Days should be non-negative")
	assert.GreaterOrEqual(t, hours, uint64(0), "Hours should be non-negative")
	assert.GreaterOrEqual(t, minutes, uint64(0), "Minutes should be non-negative")
	assert.Less(t, hours, uint64(24), "Hours should be less than 24")
	assert.Less(t, minutes, uint64(60), "Minutes should be less than 60")

	// Test UptimeI18n function
	str, err := sysinfo.UptimeI18n("Day", "Hour", "Minute", ", ")
	assert.NoError(t, err)
	assert.NotEmpty(t, str, "UptimeI18n should return a non-empty string")
}
