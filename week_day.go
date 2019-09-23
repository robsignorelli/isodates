package isodates

import "time"

// ParseWeekDay accepts an ISO-formatted year/week/day string (e.g. "2019-W04-3") and returns the
// exact date that it represents. The resulting date/time will be at midnight in UTC.
func ParseWeekDay(input string) (time.Time, error) {
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
