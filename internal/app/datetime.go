package app

import (
	"fmt"
	"time"
)

// GetDateTimeString formats a time.Time into a human-readable string.
func GetDateTimeString(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05 MST")
}

// ParseDateTimeString parses a local date or date-time string into UTC.
func ParseDateTimeString(dateTimeStr string) (time.Time, error) {
	const (
		dateLayout     = "2006-01-02"
		dateTimeLayout = "2006-01-02 15:04:05"
	)

	var (
		t   time.Time
		err error
	)

	switch len(dateTimeStr) {
	case len(dateLayout):
		// Date only → default to 23:59:59 local time
		dateTimeStr = dateTimeStr + " 23:59:59"
		t, err = time.ParseInLocation(dateTimeLayout, dateTimeStr, time.Local)

	case len(dateTimeLayout):
		// Date + time → parse as local time
		t, err = time.ParseInLocation(dateTimeLayout, dateTimeStr, time.Local)

	default:
		return time.Time{}, fmt.Errorf(
			"invalid time format %q "+
				"(expected YYYY-MM-DD or YYYY-MM-DD hh:mm:ss)",
			dateTimeStr,
		)
	}

	if err != nil {
		return time.Time{}, fmt.Errorf(
			"parse local time %q: %w",
			dateTimeStr,
			err,
		)
	}

	return t.UTC(), nil
}
