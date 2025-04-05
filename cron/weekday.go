package cron

// Weekday represents a day of the week for a cron job.
type Weekday int

const (
	Auto      Weekday = 0 // Auto used with the Weekly method to determine the weekend based on timezone.
	Sunday    Weekday = 1
	Monday    Weekday = 2
	Tuesday   Weekday = 3
	Wednesday Weekday = 4
	Thursday  Weekday = 5
	Friday    Weekday = 6
	Saturday  Weekday = 7
)

// IsValid checks if the Weekday is within the valid range (Sunday to Saturday).
func (wd Weekday) IsValid() bool {
	return wd >= Sunday && wd <= Saturday
}

// Real returns the zero-based index of the Weekday (Sunday = 0, Monday = 1, etc.).
func (wd Weekday) Real() int {
	if !wd.IsValid() {
		return 0
	}
	return int(wd) - 1
}
