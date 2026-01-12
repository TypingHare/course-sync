package ui

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ReadLine reads a line of input from standard input after displaying the
// given label. It trims any leading or trailing whitespace from the input
// before returning it.
func ReadLine(writer io.Writer, label string) (string, error) {
	colorFunc := color.New(color.FgCyan).SprintFunc()
	fmt.Fprint(writer, colorFunc(label))

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}

	return strings.TrimSpace(scanner.Text()), nil
}

// ParseUTCTime parses a date or date-time string interpreted in the local time
// zone and returns the corresponding time in UTC.
//
// Supported input formats are:
//
//   - "YYYY-MM-DD"
//     Interpreted as the given date at 23:59:00 local time.
//
//   - "YYYY-MM-DD hh:mm:ss"
//     Interpreted as the given date and time in the local time zone.
//
// If the input does not match one of the supported formats, ParseUTCTime
// returns an error. The returned time is always normalized to UTC.
func ParseUTCTime(input string) (time.Time, error) {
	const (
		dateLayout     = "2006-01-02"
		dateTimeLayout = "2006-01-02 15:04:05"
	)

	var (
		t   time.Time
		err error
	)

	switch len(input) {
	case len(dateLayout):
		// Date only → default to 23:59:59 local time
		input = input + " 23:59:59"
		t, err = time.ParseInLocation(dateTimeLayout, input, time.Local)

	case len(dateTimeLayout):
		// Date + time → parse as local time
		t, err = time.ParseInLocation(dateTimeLayout, input, time.Local)

	default:
		return time.Time{}, fmt.Errorf(
			"invalid time format %q "+
				"(expected YYYY-MM-DD or YYYY-MM-DD hh:mm:ss)",
			input,
		)
	}

	if err != nil {
		return time.Time{}, fmt.Errorf("parse local time %q: %w", input, err)
	}

	return t.UTC(), nil
}
