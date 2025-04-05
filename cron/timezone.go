package cron

// CronTZ represents the timezone configuration for a cron job.
type CronTZ struct {
	hour    int
	minute  int
	weekend Weekday
}

// NewTZ initializes and returns a new CronTZ instance.
func NewTZ() *CronTZ {
	return &CronTZ{
		hour:    0,
		minute:  0,
		weekend: Sunday,
	}
}

// SetHour sets the hour for the CronTZ configuration.
func (tz *CronTZ) SetHour(hour int) *CronTZ {
	tz.hour = hour
	return tz
}

// SetMinute sets the minute for the CronTZ configuration.
func (tz *CronTZ) SetMinute(minute int) *CronTZ {
	tz.minute = minute
	return tz
}

// SetWeekend sets the weekend day for the CronTZ configuration.
func (tz *CronTZ) SetWeekend(weekend Weekday) *CronTZ {
	tz.weekend = weekend
	return tz
}
