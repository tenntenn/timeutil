package timeutil

import (
	"time"
)

// NearHour returns nearest time by given hour.
func NearHour(now time.Time, hour int) time.Time {
	y, m, d := now.Year(), now.Month(), now.Day()
	loc := now.Location()

	if hour > 12 {
		if hour > now.Hour() {
			return time.Date(y, m, d, hour, 0, 0, 0, loc)
		}
		return time.Date(y, m, d+1, hour, 0, 0, 0, loc)
	}

	if now.Hour() <= 12 {
		if hour > now.Hour() {
			return time.Date(y, m, d, hour, 0, 0, 0, loc)
		}
		return time.Date(y, m, d, hour+12, 0, 0, 0, loc)
	}

	if hour+12 > now.Hour() {
		return time.Date(y, m, d, hour+12, 0, 0, 0, loc)
	}

	return time.Date(y, m, d+1, hour, 0, 0, 0, loc)
}
