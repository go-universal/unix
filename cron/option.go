package cron

import (
	"strconv"
	"strings"
	"time"
)

// option holds the configuration for a cron schedule.
type option struct {
	tz      *CronTZ
	reboot  bool
	minute  string
	hour    string
	day     string
	month   string
	weekday string
}

// Option defines a functional option for configuring settings.
type Option func(*option)

// WithTimezone sets the timezone for the cron schedule.
func WithTimezone(tz *CronTZ) Option {
	return func(o *option) {
		if tz != nil {
			o.tz = tz
		}
	}
}

// RunAtReboot schedules the cron to run at system reboot.
func RunAtReboot() Option {
	return func(o *option) {
		o.reboot = true
	}
}

// RunYearly schedules the cron to run once a year (January 1st at midnight).
func RunYearly() Option {
	return func(o *option) {
		o.set("0", "0", "1", "1", "*")
	}
}

// RunMonthly schedules the cron to run once a month (1st day at midnight).
func RunMonthly() Option {
	return func(o *option) {
		o.set("0", "0", "1", "*", "*")
	}
}

// RunWeekly schedules the cron to run once a week on the specified weekday.
func RunWeekly(wd Weekday) Option {
	return func(o *option) {
		o.set("0", "0", "*", "*", strconv.Itoa(o.weekend().Real()))
	}
}

// RunDaily schedules the cron to run once a day at midnight.
func RunDaily() Option {
	return func(o *option) {
		o.set("0", "0", "*", "*", "*")
	}
}

// EveryXMinutes sets the minute to run every X minutes for the cron schedule.
func EveryXMinutes(minutes int) Option {
	return func(o *option) {
		o.minute = "*/" + strconv.Itoa(minutes)
	}
}

// EveryXHours sets the hour to run every X hours for the cron schedule.
func EveryXHours(hours int) Option {
	return func(o *option) {
		o.hour = "*/" + strconv.Itoa(hours)
	}
}

// Minute sets the specific minute for the cron schedule.
func Minute(minute int) Option {
	return func(o *option) {
		if minute >= 0 && minute <= 59 {
			o.minute = strconv.Itoa(minute)
		}
	}
}

// Hour sets the specific hour for the cron schedule.
func Hour(hour int) Option {
	return func(o *option) {
		if hour >= 0 && hour <= 23 {
			o.hour = strconv.Itoa(hour)
		}
	}
}

// DayOfMonth sets the specific day of the month for the cron schedule.
func DayOfMonth(day int) Option {
	return func(o *option) {
		if day >= 1 && day <= 31 {
			o.day = strconv.Itoa(day)
		}
	}
}

// Month sets the specific month for the cron schedule.
func Month(month int) Option {
	return func(o *option) {
		if month >= 1 && month <= 12 {
			o.month = strconv.Itoa(month)
		}
	}
}

// DayOfWeek sets the specific day of the week for the cron schedule.
func DayOfWeek(wd Weekday) Option {
	return func(o *option) {
		if wd.IsValid() {
			o.weekday = strconv.Itoa(wd.Real())
		}
	}
}

// set configures the cron schedule with the provided values.
// .---------------- minute (0 - 59)
// |  .------------- hour (0 - 23)
// |  |  .---------- day of month (1 - 31)
// |  |  |  .------- month (1 - 12) OR jan, feb, mar, apr ...
// |  |  |  |  .---- day of week (0 - 6) (Sunday=0) OR sun, mon, tue, wed, thu, fri, sat
// |  |  |  |  |
// m h dom mon dow command
func (o *option) set(minute, hour, day, month, weekday string) {
	o.minute = minute
	o.hour = hour
	o.day = day
	o.month = month
	o.weekday = weekday
}

// tzHour returns the time zone hour offset as a duration.
func (o *option) tzHour() time.Duration {
	return time.Duration(o.tz.hour) * time.Hour
}

// tzMinute returns the time zone minute offset as a duration.
func (o *option) tzMinute() time.Duration {
	return time.Duration(o.tz.minute) * time.Minute
}

// weekend returns the time zone's weekend day.
func (o *option) weekend() Weekday {
	return o.tz.weekend
}

// interval generates the cron expression based on the schedule and time zone.
func (o *option) interval() string {
	defaultExpr := o.minute + " " + o.hour + " " + o.day + " " + o.month + " " + o.weekday

	// Return default expression if minute or hour is not specified
	if o.minute == "*" || strings.Contains(o.minute, "*/") ||
		o.hour == "*" || strings.Contains(o.hour, "*/") {
		return defaultExpr
	}

	// Parse cron time
	tz, err := time.Parse("15:4", o.hour+":"+o.minute)
	if err != nil {
		return defaultExpr
	}

	// Adjust time for the specified time zone
	duration := -(o.tzHour() + o.tzMinute())
	timeInTz := tz.Add(duration)

	return timeInTz.Format("4 ") +
		timeInTz.Format("15 ") +
		o.day + " " +
		o.month + " " +
		o.weekday
}
