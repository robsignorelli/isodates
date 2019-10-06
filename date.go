package isodates

import (
	"errors"
	"time"
)

// ParseDate accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// year/month/day it represents.
func ParseDate(input string) (year int, month time.Month, day int, err error) {
	// We could use the standard time package to parse this, but assuming this format
	// means that we can cut the execution time in half.
	if len(input) != 10 {
		return 0, ZeroMonth, 0, invalidFormat("YYYY-MM-DD", input)
	}
	if input[4] != '-' {
		return 0, ZeroMonth, 0, invalidFormat("YYYY-MM-DD", input)
	}
	if input[7] != '-' {
		return 0, ZeroMonth, 0, invalidFormat("YYYY-MM-DD", input)
	}

	year, err = parseYear(input[0:4])
	if err != nil {
		return 0, ZeroMonth, 0, err
	}
	month, err = parseMonth(input[5:7])
	if err != nil {
		return 0, ZeroMonth, 0, err
	}
	day, err = parseDayOfMonth(input[8:])
	if err != nil {
		return 0, ZeroMonth, 0, err
	}
	return year, month, day, nil
}

// ParseDateStart accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// given date set to exactly midnight in UTC.
func ParseDateStart(input string) (time.Time, error) {
	return ParseDateStartIn(input, time.UTC)
}

// ParseDateStartIn accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// given date set to exactly midnight in the specified location.
func ParseDateStartIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse date start: nil location")
	}
	year, month, day, err := ParseDate(input)
	if err != nil {
		return ZeroTime, err
	}
	return Midnight(year, month, day, loc), nil
}

// ParseDateEnd accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// given date set to the last nanosecond of 11:59pm in UTC.
func ParseDateEnd(input string) (time.Time, error) {
	return ParseDateEndIn(input, time.UTC)
}

// ParseDateEndIn accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// given date set to the last nanosecond of 11:59pm in the specified location.
func ParseDateEndIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse date end: nil location")
	}
	year, month, day, err := ParseDate(input)
	if err != nil {
		return ZeroTime, err
	}
	return AlmostMidnight(year, month, day, loc), nil
}
