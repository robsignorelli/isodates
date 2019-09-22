package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestWeekSuite(t *testing.T) {
	suite.Run(t, new(WeekSuite))
}

type WeekSuite struct {
	suite.Suite
	edt *time.Location
	pdt *time.Location
}

func (suite *WeekSuite) SetupSuite() {
	suite.edt, _ = time.LoadLocation("America/New_York")
	suite.pdt, _ = time.LoadLocation("America/Los_Angeles")
}

func (suite *WeekSuite) TestParseWeek() {
	succeeds := func(input string, expectedYear int, expectedWeek int) {
		year, week, err := isodates.ParseWeek(input)
		_ = suite.NoError(err) &&
			suite.Equal(expectedYear, year) &&
			suite.Equal(expectedWeek, week)
	}
	fails := func(input string) {
		_, _, err := isodates.ParseWeek(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("W01-2019")
	fails("2019/W01")
	fails("1234-W04-")
	// Invalid weeks
	fails("2019-W-1")
	fails("2019-W73")
	fails("2019-W00")
	fails("2019-WJ4")

	// Invalid years
	fails("$G33-W04")
	fails("-W04")

	// Missing zero padding
	fails("123-W04")
	fails("23-W04")
	fails("3-W04")
	fails("1234-W4")

	succeeds("2000-W01", 2000, 1)
	succeeds("2000-W11", 2000, 11)
	succeeds("2019-W11", 2019, 11)
	succeeds("1215-W06", 1215, 6)
	succeeds("0123-W12", 123, 12)
	succeeds("0012-W12", 12, 12)
	succeeds("0001-W12", 1, 12)
	succeeds("0001-W01", 1, 1)
}

func (suite *WeekSuite) TestParseWeekStart() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekStart(input)
		_ = suite.NoError(err) &&
			suite.Equal(year, date.Year()) &&
			suite.Equal(month, date.Month()) &&
			suite.Equal(day, date.Day()) &&
			suite.Equal(0, date.Hour()) &&
			suite.Equal(0, date.Minute()) &&
			suite.Equal(0, date.Second()) &&
			suite.Equal(0, date.Nanosecond()) &&
			suite.Equal(time.UTC, date.Location())
	}
	fails := func(input string) {
		_, err := isodates.ParseWeekStart(input)
		suite.Error(err)
	}

	// TestParseWeek should handle all of the variants of bad input, so make sure we propagate some.
	fails("")
	fails("not valid")
	fails("W01-2019")
	fails("2019-WJ4")

	succeeds("2019-W01", 2018, time.December, 31)
	succeeds("2019-W02", 2019, time.January, 7)
	succeeds("2000-W01", 2000, time.January, 3)
	succeeds("1999-W52", 1999, time.December, 27)
	succeeds("2000-W09", 2000, time.February, 28)
	succeeds("1999-W53", 2000, time.January, 3)   // 53rd week rolls to next year
	succeeds("2004-W53", 2004, time.December, 27) // long year where still in that year
}

func (suite *WeekSuite) TestParseWeekStartIn() {
	succeeds := func(input string, loc *time.Location, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekStartIn(input, loc)
		_ = suite.NoError(err) &&
			suite.Equal(year, date.Year()) &&
			suite.Equal(month, date.Month()) &&
			suite.Equal(day, date.Day()) &&
			suite.Equal(0, date.Hour()) &&
			suite.Equal(0, date.Minute()) &&
			suite.Equal(0, date.Second()) &&
			suite.Equal(0, date.Nanosecond()) &&
			suite.Equal(loc, date.Location())
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseWeekStartIn(input, loc)
		suite.Error(err)
	}

	// TestParseWeek should handle all of the variants of bad input, so make sure we propagate something.
	fails("", suite.edt)
	fails("not valid", suite.edt)
	fails("W01-2019", suite.edt)
	fails("2019-WJ4", suite.edt)
	fails("2019-W04", nil)

	succeeds("2019-W01", suite.edt, 2018, time.December, 31)
	succeeds("2019-W02", suite.edt, 2019, time.January, 7)
	succeeds("2000-W01", suite.edt, 2000, time.January, 3)
	succeeds("1999-W52", suite.edt, 1999, time.December, 27)
	succeeds("2000-W09", suite.edt, 2000, time.February, 28)
	succeeds("1999-W53", suite.edt, 2000, time.January, 3)   // 53rd week rolls to next year
	succeeds("2004-W53", suite.edt, 2004, time.December, 27) // long year where still in that year

	// Make sure everything still works when given a different time zone.
	succeeds("2019-W01", suite.pdt, 2018, time.December, 31)
	succeeds("2019-W02", suite.pdt, 2019, time.January, 7)
	succeeds("2000-W01", suite.pdt, 2000, time.January, 3)
	succeeds("1999-W52", suite.pdt, 1999, time.December, 27)
	succeeds("2000-W09", suite.pdt, 2000, time.February, 28)
	succeeds("1999-W53", suite.pdt, 2000, time.January, 3)
	succeeds("2004-W53", suite.pdt, 2004, time.December, 27)
}

func (suite *WeekSuite) TestParseWeekEnd() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekEnd(input)
		_ = suite.NoError(err) &&
			suite.Equal(year, date.Year()) &&
			suite.Equal(month, date.Month()) &&
			suite.Equal(day, date.Day()) &&
			suite.Equal(23, date.Hour()) &&
			suite.Equal(59, date.Minute()) &&
			suite.Equal(59, date.Second()) &&
			suite.Equal(999999999, date.Nanosecond()) &&
			suite.Equal(time.UTC, date.Location())
	}
	fails := func(input string) {
		_, err := isodates.ParseWeekEnd(input)
		suite.Error(err)
	}

	// TestParseWeek should handle all of the variants of bad input, so make sure we propagate some.
	fails("")
	fails("not valid")
	fails("W01-2019")
	fails("2019-WJ4")

	succeeds("2019-W01", 2019, time.January, 6)
	succeeds("2019-W02", 2019, time.January, 13)
	succeeds("2000-W01", 2000, time.January, 9)
	succeeds("1999-W52", 2000, time.January, 2)
	succeeds("2000-W09", 2000, time.March, 5)
	succeeds("1999-W53", 2000, time.January, 9) // 53rd week rolls to next year
	succeeds("2004-W53", 2005, time.January, 2) // long year where still in that year
}

func (suite *WeekSuite) TestParseWeekEndIn() {
	succeeds := func(input string, loc *time.Location, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekEndIn(input, loc)
		_ = suite.NoError(err) &&
			suite.Equal(year, date.Year()) &&
			suite.Equal(month, date.Month()) &&
			suite.Equal(day, date.Day()) &&
			suite.Equal(23, date.Hour()) &&
			suite.Equal(59, date.Minute()) &&
			suite.Equal(59, date.Second()) &&
			suite.Equal(999999999, date.Nanosecond()) &&
			suite.Equal(loc, date.Location())
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseWeekEndIn(input, loc)
		suite.Error(err)
	}

	// TestParseWeek should handle all of the variants of bad input, so make sure we propagate some.
	fails("", suite.edt)
	fails("not valid", suite.edt)
	fails("W01-2019", suite.edt)
	fails("2019-WJ4", suite.edt)
	fails("2019-W04", nil)

	succeeds("2019-W01", suite.edt, 2019, time.January, 6)
	succeeds("2019-W02", suite.edt, 2019, time.January, 13)
	succeeds("2000-W01", suite.edt, 2000, time.January, 9)
	succeeds("1999-W52", suite.edt, 2000, time.January, 2)
	succeeds("2000-W09", suite.edt, 2000, time.March, 5)
	succeeds("1999-W53", suite.edt, 2000, time.January, 9) // 53rd week rolls to next year
	succeeds("2004-W53", suite.edt, 2005, time.January, 2) // long year where still in that year

	// Make sure everything still works when given a different time zone.
	succeeds("2019-W01", suite.pdt, 2019, time.January, 6)
	succeeds("2019-W02", suite.pdt, 2019, time.January, 13)
	succeeds("2000-W01", suite.pdt, 2000, time.January, 9)
	succeeds("1999-W52", suite.pdt, 2000, time.January, 2)
	succeeds("2000-W09", suite.pdt, 2000, time.March, 5)
	succeeds("1999-W53", suite.pdt, 2000, time.January, 9) // 53rd week rolls to next year
	succeeds("2004-W53", suite.pdt, 2005, time.January, 2) // long year where still in that year
}

func ExampleParseWeek() {
	year, weekNumber, err := isodates.ParseWeek("2019-W02")
	fmt.Println(fmt.Sprintf("%d %d %v", year, weekNumber, err == nil))

	// We don't support months outside of 1-53
	year, weekNumber, err = isodates.ParseWeek("2019-W72")
	fmt.Println(fmt.Sprintf("%d %d %v", year, weekNumber, err == nil))

	// Output: 2019 2 true
	// 0 0 false
}

// BenchmarkParseWeek typically runs about 35-40ns/op on a 2014 MacBook Pro
func BenchmarkParseWeek(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _ = isodates.ParseWeek("2019-W04")
	}
}
