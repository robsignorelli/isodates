package isodates

import "time"

// ParseDate accepts an ISO-formatted year-month-day string (e.g. "2019-05-22") and returns the
// exact date that it represents. The resulting date/time will be at midnight in UTC.
func ParseDate(input string) (time.Time, error) {
	// We could use the standard time package to parse this, but assuming this format
	// means that we can cut the execution time in half.
	if len(input) != 10 {
		return ZeroTime, invalidFormat("YYYY-MM-DD", input)
	}
	if input[4] != '-' {
		return ZeroTime, invalidFormat("YYYY-MM-DD", input)
	}
	if input[7] != '-' {
		return ZeroTime, invalidFormat("YYYY-MM-DD", input)
	}

	year, err := parseYear(input[0:4])
	if err != nil {
		return ZeroTime, err
	}
	month, err := parseMonth(input[5:7])
	if err != nil {
		return ZeroTime, err
	}
	day, err := parseDayOfMonth(input[8:])
	if err != nil {
		return ZeroTime, err
	}
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}
