package isodates

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const ZeroMonth = time.Month(0)

var ZeroTime = time.Time{}

func parseYear(input string) (int, error) {
	year, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, errors.New("invalid year: " + input)
	}
	return int(year), nil
}

func parseMonth(input string) (time.Month, error) {
	month, err := strconv.ParseInt(input, 10, 64)
	if err != nil || month < 1 || month > 12 {
		return ZeroMonth, errors.New("invalid month: " + input)
	}
	return time.Month(month), nil
}

func parseDayOfMonth(input string) (int, error) {
	day, err := strconv.ParseInt(input, 10, 64)
	if err != nil || day < 1 {
		return 0, errors.New("invalid day of month: " + input)
	}
	return int(day), nil
}

func parseWeek(input string) (int, error) {
	week, err := strconv.ParseInt(input, 10, 64)
	if err != nil || week < 1 || week > 53 {
		return 0, errors.New("invalid week number: " + input)
	}
	return int(week), nil
}

func parseWeekOffset(input string) (int, error) {
	offset, err := strconv.ParseInt(input, 10, 64)
	if err != nil || offset < 1 || offset > 7 {
		return 0, errors.New("invalid week offset: " + input)
	}
	return int(offset), nil
}

func invalidFormat(format string, input string) error {
	return fmt.Errorf("invalid %s format: %s", format, input)
}
