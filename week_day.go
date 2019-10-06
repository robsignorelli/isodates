package isodates

import (
	"errors"
	"time"

	"github.com/snabb/isoweek"
)

// ParseWeekDay extracts all 3 numeric components from an ISO Week-Day string (e.g. "2019-W02-3").
func ParseWeekDay(input string) (year int, weekNum int, day int, err error) {
	if len(input) != 10 {
		return 0, 0, 0, invalidFormat("YYYY-W##-#", input)
	}
	if input[4:6] != "-W" {
		return 0, 0, 0, invalidFormat("YYYY-W##-#", input)
	}
	if input[8] != '-' {
		return 0, 0, 0, invalidFormat("YYYY-W##-#", input)
	}

	year, weekNum, err = ParseWeek(input[0:8])
	if err != nil {
		return 0, 0, 0, err
	}
	day, err = parseWeekOffset(input[9:])
	if err != nil {
		return 0, 0, 0, err
	}
	return year, weekNum, day, nil
}

// ParseWeekDayStart accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at midnight in UTC.
func ParseWeekDayStart(input string) (time.Time, error) {
	return ParseWeekDayStartIn(input, time.UTC)
}

// ParseWeekDayStartIn accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at midnight in the given time zone.
func ParseWeekDayStartIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse week day start: nil location")
	}
	year, weekNum, day, err := ParseWeekDay(input)
	if err != nil {
		return ZeroTime, err
	}
	startYear, startMonth, startDay := isoweek.StartDate(year, weekNum)
	return Midnight(startYear, startMonth, startDay+day-1, loc), nil
}

// ParseWeekDayEnd accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at 11:59:59pm in UTC.
func ParseWeekDayEnd(input string) (time.Time, error) {
	return ParseWeekDayEndIn(input, time.UTC)
}

// ParseWeekDayEndIn accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at 11:59:59pm in the given time zone.
func ParseWeekDayEndIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse week day end: nil location")
	}
	year, weekNum, day, err := ParseWeekDay(input)
	if err != nil {
		return ZeroTime, err
	}
	startYear, startMonth, startDay := isoweek.StartDate(year, weekNum)
	return AlmostMidnight(startYear, startMonth, startDay+day-1, loc), nil
}
