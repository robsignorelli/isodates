package isodates

import (
	"errors"
	"time"
)

// ParseMonthDay accepts an ISO-formatted month/day string (e.g. "--04-01" is April, 1) and returns the
// month and day that it represents.
func ParseMonthDay(input string) (time.Month, int, error) {
	var monthText, dayText string
	inputLength := len(input)

	switch {
	// All valid inputs are between 5 and 7 chars: "--3-1", "--03-1", "--03-01"
	case inputLength < 5 || inputLength > 7:
		return time.Month(0), 0, errors.New("invalid iso month/day: " + input)
	// Month not padded: e.g. "--3-27", "--3-05", or "--3-5"
	case input[3] == '-':
		monthText = input[2:3]
		dayText = input[4:]
	// Month *is* padded: e.g. "--03-27", "--03-05", or "--03-5"
	case input[4] == '-':
		monthText = input[2:4]
		dayText = input[5:]
	default:
		return time.Month(0), 0, errors.New("invalid iso month/day: " + input)
	}

	month, err := parseMonth(monthText)
	if err != nil {
		return ZeroMonth, 0, err
	}

	day, err := parseDayOfMonth(dayText)
	if err != nil {
		return ZeroMonth, 0, err
	}

	// You have a potential for something like January 32nd which is actually Feb 1st.
	if day > 28 {
		d := time.Date(2000, time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
		return d.Month(), d.Day(), nil
	}
	return time.Month(month), int(day), nil
}
