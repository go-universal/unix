package sysinfo_test

import (
	"testing"

	"github.com/go-universal/unix/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {
	sent, recv, err := sysinfo.NetworkInfo()
	assert.NoError(t, err)

	// Ensure sent and recv are non-negative
	assert.GreaterOrEqual(t, sent, uint64(0), "Sent bytes should be non-negative")
	assert.GreaterOrEqual(t, recv, uint64(0), "Received bytes should be non-negative")
}
