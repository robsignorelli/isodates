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

// ParseMonthDayStart parses the month/day string (e.g. "--12-24") and returns a date/time at
// midnight in the specified year. The resulting timestamp will be in UTC.
func ParseMonthDayStart(input string, year int) (time.Time, error) {
	return ParseMonthDayStartIn(input, year, time.UTC)
}

// ParseMonthDayStartIn parses the month/day string (e.g. "--12-24") and returns a date/time at
// midnight in the specified year. The resulting timestamp will be in specified time zone.
func ParseMonthDayStartIn(input string, year int, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse month day start: nil location")
	}
	month, day, err := ParseMonthDay(input)
	if err != nil {
		return ZeroTime, err
	}
	return Midnight(year, month, day, loc), nil
}

// ParseMonthDayEnd parses the month/day string (e.g. "--12-24") and returns a date/time at
// 11:59:59pm in the specified year. The resulting timestamp will be in UTC.
func ParseMonthDayEnd(input string, year int) (time.Time, error) {
	return ParseMonthDayEndIn(input, year, time.UTC)
}

// ParseMonthDayEndIn parses the month/day string (e.g. "--12-24") and returns a date/time at
// 11:59:59pm in the specified year. The resulting timestamp will be in the specified time zone.
func ParseMonthDayEndIn(input string, year int, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse month day end: nil location")
	}
	month, day, err := ParseMonthDay(input)
	if err != nil {
		return ZeroTime, err
	}
	return AlmostMidnight(year, month, day, loc), nil
}
