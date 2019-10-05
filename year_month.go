package isodates

import (
	"errors"
	"time"
)

// ParseYearMonth accepts an ISO string such as "2019-04" and returns the individual date
// components for the year and month (e.g. 2019 and time.April). We also support the variant
// where you can prefix the year with either "+" or "-".
func ParseYearMonth(input string) (int, time.Month, error) {
	inputLength := len(input)
	yearText := ""
	monthText := ""

	// Must either by "YYYY-MM", "+YYYY-MM", or "-YYYY-MM"
	if inputLength < 7 || inputLength > 8 {
		return 0, ZeroMonth, invalidFormat("[+-]YYYY-MM", input)
	}

	// Either "+YYYY-MM" or "-YYYY-MM"
	if inputLength == 8 {
		if input[5] != '-' {
			return 0, ZeroMonth, invalidFormat("[+-]YYYY-MM", input)
		}

		// For 8-character variant, the first character must be '+' or '-'
		switch input[0] {
		case '+':
			yearText = input[1:5]
			monthText = input[6:]
		case '-':
			yearText = input[0:5]
			monthText = input[6:]
		default:
			return 0, ZeroMonth, invalidFormat("[+-]YYYY-MM", input)
		}
	}

	// "YYYY-MM" format
	if inputLength == 7 {
		if input[4] != '-' {
			return 0, ZeroMonth, invalidFormat("YYYY-MM", input)
		}
		yearText = input[0:4]
		monthText = input[5:]
	}

	year, err := parseYear(yearText)
	if err != nil {
		return 0, ZeroMonth, err
	}
	month, err := parseMonth(monthText)
	if err != nil {
		return 0, ZeroMonth, err
	}
	return year, month, nil
}

// ParseYearMonthStart returns the first day of the year/month for the parsed input. The
// resulting date will be at midnight in UTC.
func ParseYearMonthStart(input string) (time.Time, error) {
	return ParseYearMonthStartIn(input, time.UTC)
}

// ParseYearMonthStartIn returns the first day of the year/month for the parsed input. The
// resulting date will be at midnight in the specified time zone.
func ParseYearMonthStartIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse year month start: nil location")
	}
	year, month, err := ParseYearMonth(input)
	if err != nil {
		return ZeroTime, err
	}
	return Midnight(year, month, 1, loc), nil
}

// ParseYearMonthEnd returns the last day of the year/month for the parsed input. The
// resulting date will be at 11:59:59pm in UTC.
func ParseYearMonthEnd(input string) (time.Time, error) {
	return ParseYearMonthEndIn(input, time.UTC)
}

// ParseYearMonthEndIn returns the last day of the year/month for the parsed input. The
// resulting date will be at 11:59:59pm in the specified time zone.
func ParseYearMonthEndIn(input string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return ZeroTime, errors.New("parse year month end: nil location")
	}
	year, month, err := ParseYearMonth(input)
	if err != nil {
		return ZeroTime, err
	}
	firstOfMonth := AlmostMidnight(year, month, 1, loc)
	return firstOfMonth.AddDate(0, 1, -1), nil
}
