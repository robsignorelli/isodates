package isodates

import (
	"time"
)

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
