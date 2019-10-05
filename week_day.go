package isodates

import (
	"errors"
	"time"
)

// ParseWeekDayStart accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at midnight in UTC.
func ParseWeekDayStart(input string) (time.Time, error) {
	if len(input) != 10 {
		return ZeroTime, invalidFormat("YYYY-W##-#", input)
	}
	if input[4:6] != "-W" {
		return ZeroTime, invalidFormat("YYYY-W##-#", input)
	}
	if input[8] != '-' {
		return ZeroTime, invalidFormat("YYYY-W##-#", input)
	}

	weekStart, err := ParseWeekStart(input[0:8])
	if err != nil {
		return ZeroTime, err
	}

	switch input[9] {
	case '1':
		return weekStart, nil
	}
	offset, err := parseWeekOffset(input[9:])
	if err != nil {
		return ZeroTime, err
	}

	// A suffix of "-1" means the first day of the week, so you should actually add zero days.
	return weekStart.AddDate(0, 0, offset-1), nil
}

// ParseWeekDayStart accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at midnight in UTC.
func ParseWeekDayStartIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse week day start: nil location")
	}
	date, err := ParseWeekDayStart(input)
	if err != nil {
		return ZeroTime, err
	}
	return Midnight(date.Year(), date.Month(), date.Day(), loc), nil
}

// ParseWeekDayEnd accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at 11:59:59pm in UTC.
func ParseWeekDayEnd(input string) (time.Time, error) {
	date, err := ParseWeekDayStart(input)
	if err != nil {
		return ZeroTime, err
	}
	return AlmostMidnight(date.Year(), date.Month(), date.Day(), time.UTC), nil
}

// ParseWeekDayEndIn accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at 11:59:59pm in the given time zone.
func ParseWeekDayEndIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse week day start: nil location")
	}
	date, err := ParseWeekDayStart(input)
	if err != nil {
		return ZeroTime, err
	}
	return AlmostMidnight(date.Year(), date.Month(), date.Day(), loc), nil
}
