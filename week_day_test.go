package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestWeekDaySuite(t *testing.T) {
	suite.Run(t, new(WeekDaySuite))
}

type WeekDaySuite struct {
	ChronoSuite
}

func (suite *WeekDaySuite) TestParseWeekDay() {
	succeeds := func(input string, year, week, day int) {
		actualYear, actualWeek, actualDay, err := isodates.ParseWeekDay(input)
		_ = suite.NoError(err) &&
			suite.Equal(year, actualYear) &&
			suite.Equal(week, actualWeek) &&
			suite.Equal(day, actualDay)
	}
	fails := func(input string) {
		_, err := isodates.ParseWeekDayStart(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("W01-2019-1")
	fails("2019/W01/2")
	fails("1234-W04-")

	// Invalid years
	fails("$G33-W04-3")
	fails("-W04-3")

	// Invalid weeks
	fails("2019-W-1")
	fails("2019-W73-1")
	fails("2019-W00-3")
	fails("2019-WJ4-4")

	// Invalid offsets
	fails("2019-W04-9")
	fails("2019-W73-X")
	fails("2019-W00-0")
	fails("2019-WJ4-44")

	// Missing zero padding
	fails("123-W04-1")
	fails("23-W04-1")
	fails("3-W04-1")
	fails("1234-W4-1")
	fails("1234-W4-03") // day offset shouldn't be padded

	succeeds("2019-W01-1", 2019, 1, 1)
	succeeds("2019-W01-2", 2019, 1, 2)
	succeeds("2019-W01-3", 2019, 1, 3)
	succeeds("2019-W01-4", 2019, 1, 4)
	succeeds("2019-W01-5", 2019, 1, 5)
	succeeds("2019-W01-6", 2019, 1, 6)
	succeeds("2019-W01-7", 2019, 1, 7)

	succeeds("2019-W02-1", 2019, 2, 1)
	succeeds("2019-W02-2", 2019, 2, 2)
	succeeds("2019-W02-3", 2019, 2, 3)
	succeeds("2019-W02-4", 2019, 2, 4)
	succeeds("2019-W02-5", 2019, 2, 5)
	succeeds("2019-W02-6", 2019, 2, 6)
	succeeds("2019-W02-7", 2019, 2, 7)

	succeeds("2004-W09-1", 2004, 9, 1)
	succeeds("2004-W09-2", 2004, 9, 2)
	succeeds("2004-W09-3", 2004, 9, 3)
	succeeds("2004-W09-4", 2004, 9, 4)
	succeeds("2004-W09-5", 2004, 9, 5)
	succeeds("2004-W09-6", 2004, 9, 6)
	succeeds("2004-W09-7", 2004, 9, 7)

	succeeds("2004-W53-1", 2004, 53, 1)
	succeeds("2004-W53-2", 2004, 53, 2)
	succeeds("2004-W53-3", 2004, 53, 3)
	succeeds("2004-W53-4", 2004, 53, 4)
	succeeds("2004-W53-5", 2004, 53, 5)
	succeeds("2004-W53-6", 2004, 53, 6)
	succeeds("2004-W53-7", 2004, 53, 7)
}

func (suite *WeekDaySuite) TestParseWeekDayStart() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekDayStart(input)
		suite.AssertMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseWeekDayStart(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("W01-2019-1")
	fails("2019/W01/2")
	fails("1234-W04-")

	// Invalid years
	fails("$G33-W04-3")
	fails("-W04-3")

	// Invalid weeks
	fails("2019-W-1")
	fails("2019-W73-1")
	fails("2019-W00-3")
	fails("2019-WJ4-4")

	// Invalid offsets
	fails("2019-W04-9")
	fails("2019-W73-X")
	fails("2019-W00-0")
	fails("2019-WJ4-44")

	// Missing zero padding
	fails("123-W04-1")
	fails("23-W04-1")
	fails("3-W04-1")
	fails("1234-W4-1")
	fails("1234-W4-03") // day offset shouldn't be padded

	succeeds("2019-W01-1", 2018, time.December, 31)
	succeeds("2019-W01-2", 2019, time.January, 1)
	succeeds("2019-W01-3", 2019, time.January, 2)
	succeeds("2019-W01-4", 2019, time.January, 3)
	succeeds("2019-W01-5", 2019, time.January, 4)
	succeeds("2019-W01-6", 2019, time.January, 5)
	succeeds("2019-W01-7", 2019, time.January, 6)

	succeeds("2019-W02-1", 2019, time.January, 7)
	succeeds("2019-W02-2", 2019, time.January, 8)
	succeeds("2019-W02-3", 2019, time.January, 9)
	succeeds("2019-W02-4", 2019, time.January, 10)
	succeeds("2019-W02-5", 2019, time.January, 11)
	succeeds("2019-W02-6", 2019, time.January, 12)
	succeeds("2019-W02-7", 2019, time.January, 13)

	succeeds("2004-W09-1", 2004, time.February, 23)
	succeeds("2004-W09-2", 2004, time.February, 24)
	succeeds("2004-W09-3", 2004, time.February, 25)
	succeeds("2004-W09-4", 2004, time.February, 26)
	succeeds("2004-W09-5", 2004, time.February, 27)
	succeeds("2004-W09-6", 2004, time.February, 28)
	succeeds("2004-W09-7", 2004, time.February, 29)

	succeeds("2004-W53-1", 2004, time.December, 27)
	succeeds("2004-W53-2", 2004, time.December, 28)
	succeeds("2004-W53-3", 2004, time.December, 29)
	succeeds("2004-W53-4", 2004, time.December, 30)
	succeeds("2004-W53-5", 2004, time.December, 31)
	succeeds("2004-W53-6", 2005, time.January, 1)
	succeeds("2004-W53-7", 2005, time.January, 2)
}

func (suite *WeekDaySuite) TestParseWeekDayStartIn() {
	succeeds := func(input string, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseWeekDayStartIn(input, loc)
		suite.AssertMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseWeekDayStartIn(input, loc)
		suite.Error(err)
	}
	fails("", locationEDT)
	fails("", locationPDT)
	fails("not valid", locationEDT)
	fails("not valid", locationPDT)
	fails("2019-W01-1", nil)

	succeeds("2019-W01-1", 2018, time.December, 31, locationEDT)
	succeeds("2019-W01-2", 2019, time.January, 1, locationEDT)
	succeeds("2019-W01-3", 2019, time.January, 2, locationEDT)
	succeeds("2019-W01-4", 2019, time.January, 3, locationEDT)
	succeeds("2019-W01-5", 2019, time.January, 4, locationEDT)
	succeeds("2019-W01-6", 2019, time.January, 5, locationEDT)
	succeeds("2019-W01-7", 2019, time.January, 6, locationEDT)

	succeeds("2019-W01-1", 2018, time.December, 31, locationPDT)
	succeeds("2019-W01-2", 2019, time.January, 1, locationPDT)
	succeeds("2019-W01-3", 2019, time.January, 2, locationPDT)
	succeeds("2019-W01-4", 2019, time.January, 3, locationPDT)
	succeeds("2019-W01-5", 2019, time.January, 4, locationPDT)
	succeeds("2019-W01-6", 2019, time.January, 5, locationPDT)
	succeeds("2019-W01-7", 2019, time.January, 6, locationPDT)

	succeeds("2004-W09-1", 2004, time.February, 23, locationEDT)
	succeeds("2004-W09-2", 2004, time.February, 24, locationEDT)
	succeeds("2004-W09-3", 2004, time.February, 25, locationEDT)
	succeeds("2004-W09-4", 2004, time.February, 26, locationEDT)
	succeeds("2004-W09-5", 2004, time.February, 27, locationEDT)
	succeeds("2004-W09-6", 2004, time.February, 28, locationEDT)
	succeeds("2004-W09-7", 2004, time.February, 29, locationEDT)

	succeeds("2004-W09-1", 2004, time.February, 23, locationPDT)
	succeeds("2004-W09-2", 2004, time.February, 24, locationPDT)
	succeeds("2004-W09-3", 2004, time.February, 25, locationPDT)
	succeeds("2004-W09-4", 2004, time.February, 26, locationPDT)
	succeeds("2004-W09-5", 2004, time.February, 27, locationPDT)
	succeeds("2004-W09-6", 2004, time.February, 28, locationPDT)
	succeeds("2004-W09-7", 2004, time.February, 29, locationPDT)

	succeeds("2004-W53-1", 2004, time.December, 27, locationEDT)
	succeeds("2004-W53-2", 2004, time.December, 28, locationEDT)
	succeeds("2004-W53-3", 2004, time.December, 29, locationEDT)
	succeeds("2004-W53-4", 2004, time.December, 30, locationEDT)
	succeeds("2004-W53-5", 2004, time.December, 31, locationEDT)
	succeeds("2004-W53-6", 2005, time.January, 1, locationEDT)
	succeeds("2004-W53-7", 2005, time.January, 2, locationEDT)

	succeeds("2004-W53-1", 2004, time.December, 27, locationPDT)
	succeeds("2004-W53-2", 2004, time.December, 28, locationPDT)
	succeeds("2004-W53-3", 2004, time.December, 29, locationPDT)
	succeeds("2004-W53-4", 2004, time.December, 30, locationPDT)
	succeeds("2004-W53-5", 2004, time.December, 31, locationPDT)
	succeeds("2004-W53-6", 2005, time.January, 1, locationPDT)
	succeeds("2004-W53-7", 2005, time.January, 2, locationPDT)
}

func (suite *WeekDaySuite) TestParseWeekDayEnd() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekDayEnd(input)
		suite.AssertAlmostMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseWeekDayEnd(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("W01-2019-1")
	fails("2019/W01/2")
	fails("1234-W04-")

	// Invalid years
	fails("$G33-W04-3")
	fails("-W04-3")

	// Invalid weeks
	fails("2019-W-1")
	fails("2019-W73-1")
	fails("2019-W00-3")
	fails("2019-WJ4-4")

	// Invalid offsets
	fails("2019-W04-9")
	fails("2019-W73-X")
	fails("2019-W00-0")
	fails("2019-WJ4-44")

	// Missing zero padding
	fails("123-W04-1")
	fails("23-W04-1")
	fails("3-W04-1")
	fails("1234-W4-1")
	fails("1234-W4-03") // day offset shouldn't be padded

	succeeds("2019-W01-1", 2018, time.December, 31)
	succeeds("2019-W01-2", 2019, time.January, 1)
	succeeds("2019-W01-3", 2019, time.January, 2)
	succeeds("2019-W01-4", 2019, time.January, 3)
	succeeds("2019-W01-5", 2019, time.January, 4)
	succeeds("2019-W01-6", 2019, time.January, 5)
	succeeds("2019-W01-7", 2019, time.January, 6)

	succeeds("2019-W02-1", 2019, time.January, 7)
	succeeds("2019-W02-2", 2019, time.January, 8)
	succeeds("2019-W02-3", 2019, time.January, 9)
	succeeds("2019-W02-4", 2019, time.January, 10)
	succeeds("2019-W02-5", 2019, time.January, 11)
	succeeds("2019-W02-6", 2019, time.January, 12)
	succeeds("2019-W02-7", 2019, time.January, 13)

	succeeds("2004-W09-1", 2004, time.February, 23)
	succeeds("2004-W09-2", 2004, time.February, 24)
	succeeds("2004-W09-3", 2004, time.February, 25)
	succeeds("2004-W09-4", 2004, time.February, 26)
	succeeds("2004-W09-5", 2004, time.February, 27)
	succeeds("2004-W09-6", 2004, time.February, 28)
	succeeds("2004-W09-7", 2004, time.February, 29)

	succeeds("2004-W53-1", 2004, time.December, 27)
	succeeds("2004-W53-2", 2004, time.December, 28)
	succeeds("2004-W53-3", 2004, time.December, 29)
	succeeds("2004-W53-4", 2004, time.December, 30)
	succeeds("2004-W53-5", 2004, time.December, 31)
	succeeds("2004-W53-6", 2005, time.January, 1)
	succeeds("2004-W53-7", 2005, time.January, 2)
}

func (suite *WeekDaySuite) TestParseWeekDayEndIn() {
	succeeds := func(input string, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseWeekDayEndIn(input, loc)
		suite.AssertAlmostMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseWeekDayEndIn(input, loc)
		suite.Error(err)
	}
	fails("", locationEDT)
	fails("", locationPDT)
	fails("not valid", locationEDT)
	fails("not valid", locationPDT)
	fails("2019-W01-1", nil)

	succeeds("2019-W01-1", 2018, time.December, 31, locationEDT)
	succeeds("2019-W01-2", 2019, time.January, 1, locationEDT)
	succeeds("2019-W01-3", 2019, time.January, 2, locationEDT)
	succeeds("2019-W01-4", 2019, time.January, 3, locationEDT)
	succeeds("2019-W01-5", 2019, time.January, 4, locationEDT)
	succeeds("2019-W01-6", 2019, time.January, 5, locationEDT)
	succeeds("2019-W01-7", 2019, time.January, 6, locationEDT)

	succeeds("2019-W01-1", 2018, time.December, 31, locationPDT)
	succeeds("2019-W01-2", 2019, time.January, 1, locationPDT)
	succeeds("2019-W01-3", 2019, time.January, 2, locationPDT)
	succeeds("2019-W01-4", 2019, time.January, 3, locationPDT)
	succeeds("2019-W01-5", 2019, time.January, 4, locationPDT)
	succeeds("2019-W01-6", 2019, time.January, 5, locationPDT)
	succeeds("2019-W01-7", 2019, time.January, 6, locationPDT)

	succeeds("2004-W09-1", 2004, time.February, 23, locationEDT)
	succeeds("2004-W09-2", 2004, time.February, 24, locationEDT)
	succeeds("2004-W09-3", 2004, time.February, 25, locationEDT)
	succeeds("2004-W09-4", 2004, time.February, 26, locationEDT)
	succeeds("2004-W09-5", 2004, time.February, 27, locationEDT)
	succeeds("2004-W09-6", 2004, time.February, 28, locationEDT)
	succeeds("2004-W09-7", 2004, time.February, 29, locationEDT)

	succeeds("2004-W09-1", 2004, time.February, 23, locationPDT)
	succeeds("2004-W09-2", 2004, time.February, 24, locationPDT)
	succeeds("2004-W09-3", 2004, time.February, 25, locationPDT)
	succeeds("2004-W09-4", 2004, time.February, 26, locationPDT)
	succeeds("2004-W09-5", 2004, time.February, 27, locationPDT)
	succeeds("2004-W09-6", 2004, time.February, 28, locationPDT)
	succeeds("2004-W09-7", 2004, time.February, 29, locationPDT)

	succeeds("2004-W53-1", 2004, time.December, 27, locationEDT)
	succeeds("2004-W53-2", 2004, time.December, 28, locationEDT)
	succeeds("2004-W53-3", 2004, time.December, 29, locationEDT)
	succeeds("2004-W53-4", 2004, time.December, 30, locationEDT)
	succeeds("2004-W53-5", 2004, time.December, 31, locationEDT)
	succeeds("2004-W53-6", 2005, time.January, 1, locationEDT)
	succeeds("2004-W53-7", 2005, time.January, 2, locationEDT)

	succeeds("2004-W53-1", 2004, time.December, 27, locationPDT)
	succeeds("2004-W53-2", 2004, time.December, 28, locationPDT)
	succeeds("2004-W53-3", 2004, time.December, 29, locationPDT)
	succeeds("2004-W53-4", 2004, time.December, 30, locationPDT)
	succeeds("2004-W53-5", 2004, time.December, 31, locationPDT)
	succeeds("2004-W53-6", 2005, time.January, 1, locationPDT)
	succeeds("2004-W53-7", 2005, time.January, 2, locationPDT)
}

func ExampleParseWeekDay() {
	date, err := isodates.ParseWeekDayStart("2019-W02-2")
	if err != nil {
		fmt.Printf("oops: %v\n", err)
	}
	fmt.Println(date.Format("Jan 2, 2006"))

	// Output: Jan 8, 2019
}

// BenchmarkParseWeekDay typically runs about 170-185ns/op on a 2014 MacBook Pro
func BenchmarkParseWeekDay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = isodates.ParseWeekDayStart("2019-W04-3")
	}
}
