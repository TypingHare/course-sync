package app

import "time"

// GetDateTimeString formats a time.Time into a human-readable string.
func GetDateTimeString(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05 MST")
}
