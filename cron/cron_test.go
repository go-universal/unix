package cron_test

import (
	"testing"

	"github.com/go-universal/unix/cron"
	"github.com/stretchr/testify/assert"
)

func TestCronGenerator(t *testing.T) {
	data := map[string]cron.Cron{
		"@reboot do some": cron.New("do some", cron.RunAtReboot()),
		"30 20 * * 0 do some": cron.New(
			"do some",
			cron.WithTimezone(cron.NewTZ().SetHour(3).SetMinute(30)),
			cron.RunWeekly(cron.Auto),
		),
		"0 03 * * 3 do some": cron.New(
			"do some",
			cron.WithTimezone(cron.NewTZ().SetHour(0).SetMinute(0).SetWeekend(cron.Wednesday)),
			cron.RunWeekly(cron.Auto),
			cron.Hour(3),
			cron.Minute(0),
		),
	}

	for expected, cron := range data {
		result := cron.Raw()
		assert.Equal(t, expected, result, "Cron job did not match expected output")
	}
}
