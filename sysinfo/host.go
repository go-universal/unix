package sysinfo

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/host"
)

// Uptime retrieves the system uptime and returns it as a time.Duration.
// It returns an error if the uptime cannot be determined.
func Uptime() (time.Duration, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return 0, err
	}

	// Convert uptime in seconds to time.Duration.
	return time.Duration(uptime) * time.Second, nil
}

// UptimeParts retrieves the system uptime and returns it as days, hours, and minutes.
// It also returns an error if the uptime cannot be determined.
func UptimeParts() (days, hours, minutes uint64, err error) {
	uptime, err := host.Uptime()
	if err != nil {
		return 0, 0, 0, err
	}

	// Calculate days, hours, and minutes from the uptime in seconds.
	days = uptime / (60 * 60 * 24)
	hours = (uptime % (60 * 60 * 24)) / (60 * 60)
	minutes = (uptime % (60 * 60)) / 60
	return days, hours, minutes, nil
}

// UptimeI18n retrieves the system uptime and formats it as a localized string.
// The titles for days, hours, and minutes, as well as the separator, are provided as arguments.
func UptimeI18n(dayTitle, hourTitle, minuteTitle, separator string) (string, error) {
	days, hours, minutes, err := UptimeParts()
	if err != nil {
		return "", err
	}

	// Build the localized uptime string.
	var res []string
	if days > 0 {
		res = append(res, fmt.Sprintf("%d %s", days, dayTitle))
	}
	if hours > 0 {
		res = append(res, fmt.Sprintf("%d %s", hours, hourTitle))
	}
	res = append(res, fmt.Sprintf("%d %s", minutes, minuteTitle))

	return strings.Join(res, separator), nil
}
