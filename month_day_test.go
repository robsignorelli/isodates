package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestMonthDaySuite(t *testing.T) {
	suite.Run(t, new(MonthDaySuite))
}

type MonthDaySuite struct {
	ChronoSuite
}

func (suite *MonthDaySuite) TestParseMonthDay() {
	succeeds := func(input string, expectedMonth time.Month, expectedDate int) {
		month, date, err := isodates.ParseMonthDay(input)
		_ = suite.NoError(err) &&
			suite.Equal(expectedMonth, month) &&
			suite.Equal(expectedDate, date)
	}
	fails := func(input string) {
		_, _, err := isodates.ParseMonthDay(input)
		suite.Error(err)
	}

	fails("")
	fails("not valid")
	fails("-----")

	// This is an "oh" not a zero....
	fails("--O1-1")

	// Will attempt to parse "-1" as the month
	fails("---1-1")

	// Will attempt to parse "-2" as the day
	fails("--1--2")

	// Not valid months or days
	fails("--00-01")
	fails("--01-00")

	succeeds("--1-1", time.January, 1)
	succeeds("--01-1", time.January, 1)
	succeeds("--1-01", time.January, 1)
	succeeds("--01-01", time.January, 1)
	succeeds("--01-31", time.January, 31)

	succeeds("--5-1", time.May, 1)
	succeeds("--05-1", time.May, 1)
	succeeds("--5-01", time.May, 1)
	succeeds("--05-01", time.May, 1)
	succeeds("--05-30", time.May, 30)

	succeeds("--12-1", time.December, 1)
	succeeds("--12-01", time.December, 1)
	succeeds("--12-27", time.December, 27)

	// Roll over to subsequent months and we assume leap years
	succeeds("--05-32", time.June, 1)
	succeeds("--05-65", time.July, 4)
	succeeds("--02-28", time.February, 28)
	succeeds("--02-29", time.February, 29)
	succeeds("--02-30", time.March, 1)
}

func (suite *MonthDaySuite) TestParseMonthDayStart() {
	succeeds := func(input string, inputYear int, year int, month time.Month, day int) {
		date, err := isodates.ParseMonthDayStart(input, inputYear)
		suite.AssertMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseMonthDayStart(input, 2000)
		suite.Error(err)
	}

	fails("")
	fails("not valid")

	succeeds("--01-01", 2000, 2000, time.January, 1)
	succeeds("--01-01", 2019, 2019, time.January, 1)
	succeeds("--01-31", 2000, 2000, time.January, 31)
	succeeds("--01-31", 2018, 2018, time.January, 31)
	succeeds("--05-30", 123, 123, time.May, 30)
	succeeds("--05-30", 23, 23, time.May, 30)
	succeeds("--05-30", 1, 1, time.May, 30)
}

func (suite *MonthDaySuite) TestParseMonthDayStartIn() {
	succeeds := func(input string, inputYear int, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseMonthDayStartIn(input, inputYear, loc)
		suite.AssertMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseMonthDayStartIn(input, 2000, loc)
		suite.Error(err)
	}

	fails("", locationEDT)
	fails("not valid", locationEDT)
	fails("--01-01", nil)

	succeeds("--01-01", 2000, 2000, time.January, 1, locationEDT)
	succeeds("--01-01", 2019, 2019, time.January, 1, locationEDT)
	succeeds("--01-31", 2000, 2000, time.January, 31, locationEDT)
	succeeds("--01-31", 2018, 2018, time.January, 31, locationEDT)
	succeeds("--05-30", 123, 123, time.May, 30, locationEDT)
	succeeds("--05-30", 23, 23, time.May, 30, locationEDT)
	succeeds("--05-30", 1, 1, time.May, 30, locationEDT)

	succeeds("--01-01", 2000, 2000, time.January, 1, locationPDT)
	succeeds("--01-01", 2019, 2019, time.January, 1, locationPDT)
	succeeds("--01-31", 2000, 2000, time.January, 31, locationPDT)
	succeeds("--01-31", 2018, 2018, time.January, 31, locationPDT)
	succeeds("--05-30", 123, 123, time.May, 30, locationPDT)
	succeeds("--05-30", 23, 23, time.May, 30, locationPDT)
	succeeds("--05-30", 1, 1, time.May, 30, locationPDT)
}

func (suite *MonthDaySuite) TestParseMonthDayEnd() {
	succeeds := func(input string, inputYear int, year int, month time.Month, day int) {
		date, err := isodates.ParseMonthDayEnd(input, inputYear)
		suite.AssertAlmostMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseMonthDayEnd(input, 2000)
		suite.Error(err)
	}

	fails("")
	fails("not valid")

	succeeds("--01-01", 2000, 2000, time.January, 1)
	succeeds("--01-01", 2019, 2019, time.January, 1)
	succeeds("--01-31", 2000, 2000, time.January, 31)
	succeeds("--01-31", 2018, 2018, time.January, 31)
	succeeds("--05-30", 123, 123, time.May, 30)
	succeeds("--05-30", 23, 23, time.May, 30)
	succeeds("--05-30", 1, 1, time.May, 30)
}

func (suite *MonthDaySuite) TestParseMonthDayEndIn() {
	succeeds := func(input string, inputYear int, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseMonthDayEndIn(input, inputYear, loc)
		suite.AssertAlmostMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseMonthDayEndIn(input, 2000, loc)
		suite.Error(err)
	}

	fails("", locationEDT)
	fails("not valid", locationEDT)
	fails("--01-01", nil)

	succeeds("--01-01", 2000, 2000, time.January, 1, locationEDT)
	succeeds("--01-01", 2019, 2019, time.January, 1, locationEDT)
	succeeds("--01-31", 2000, 2000, time.January, 31, locationEDT)
	succeeds("--01-31", 2018, 2018, time.January, 31, locationEDT)
	succeeds("--05-30", 123, 123, time.May, 30, locationEDT)
	succeeds("--05-30", 23, 23, time.May, 30, locationEDT)
	succeeds("--05-30", 1, 1, time.May, 30, locationEDT)

	succeeds("--01-01", 2000, 2000, time.January, 1, locationPDT)
	succeeds("--01-01", 2019, 2019, time.January, 1, locationPDT)
	succeeds("--01-31", 2000, 2000, time.January, 31, locationPDT)
	succeeds("--01-31", 2018, 2018, time.January, 31, locationPDT)
	succeeds("--05-30", 123, 123, time.May, 30, locationPDT)
	succeeds("--05-30", 23, 23, time.May, 30, locationPDT)
	succeeds("--05-30", 1, 1, time.May, 30, locationPDT)
}

func ExampleParseMonthDay() {
	// Standard usage
	month, day, err := isodates.ParseMonthDay("--04-01")
	fmt.Println(fmt.Sprintf("%d %d %v", month, day, err == nil))

	// Automatic rollover to subsequent months
	month, day, err = isodates.ParseMonthDay("--01-34")
	fmt.Println(fmt.Sprintf("%d %d %v", month, day, err == nil))

	// Output: 4 1 true
	// 2 3 true
}

// BenchmarkParseMonthDay typically runs about 35ns/op on a 2014 MacBook Pro
func BenchmarkParseMonthDay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _ = isodates.ParseMonthDay("--05-19")
	}
}
