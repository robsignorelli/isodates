package isodates

import (
	"errors"
	"time"

	"github.com/snabb/isoweek"
)

/*
 * We rely heavily on github.com/snabb/isoweek for this parsing. There are a ton of algorithms
 * of varying complexity and performance. His library is already amazing, so no need to
 * reinvent the wheel.
 */

// ParseWeek accepts an ISO-formatted year/week string (e.g. "2019-W04") and returns the
// year and week number that it represents.
func ParseWeek(input string) (year int, week int, err error) {
	if len(input) != 8 {
		return 0, 0, invalidFormat("YYYY-W##", input)
	}
	if input[4:6] != "-W" {
		return 0, 0, invalidFormat("YYYY-W##", input)
	}

	yearText := input[0:4]
	weekText := input[6:]

	year, err = parseYear(yearText)
	if err != nil {
		return 0, 0, err
	}
	week, err = parseWeek(weekText)
	if err != nil {
		return 0, 0, err
	}
	return year, week, nil
}

// ParseWeekStart returns midnight on Monday of the specified ISO week string. The resulting
// date/time will be in UTC. If you would like this to be midnight of some local time, use ParseWeekStartIn.
func ParseWeekStart(input string) (time.Time, error) {
	return ParseWeekStartIn(input, time.UTC)
}

// ParseWeekEndIn returns midnight on Monday of the specified ISO week string. This will be in the
// local time of the specified location.
func ParseWeekStartIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, errors.New("parse week: nil location")
	}
	isoYear, isoWeek, err := ParseWeek(input)
	if err != nil {
		return time.Time{}, err
	}

	year, month, day := isoweek.StartDate(isoYear, isoWeek)
	return time.Date(year, month, day, 0, 0, 0, 0, loc), nil
}

// ParseWeekStart returns 11:59:59pm (one nanosecond before midnight) on Sunday of the specified ISO week
// string. The resulting date/time will be in UTC. If you would like this to be almost-midnight of some local
// time, use ParseWeekEndIn.
func ParseWeekEnd(input string) (time.Time, error) {
	return ParseWeekEndIn(input, time.UTC)
}

// ParseWeekEndIn returns 11:59:59pm (one nanosecond before midnight) on Sunday of the specified ISO week
// string. This will be in the local time of the specified location.
func ParseWeekEndIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, errors.New("parse week: nil location")
	}
	isoYear, isoWeek, err := ParseWeek(input)
	if err != nil {
		return time.Time{}, err
	}

	year, month, day := isoweek.StartDate(isoYear, isoWeek)
	return time.Date(year, month, day + 6, 23, 59, 59, 999999999, loc), nil
}