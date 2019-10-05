package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestDateSuite(t *testing.T) {
	suite.Run(t, new(DateSuite))
}

type DateSuite struct {
	ChronoSuite
}

func (suite DateSuite) TestParseDate() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseDate(input)
		suite.AssertMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseDate(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("01-2019-21")
	fails("2019/01/02")

	// Invalid year
	fails("$G33-04-03")
	fails("-04-03")
	fails("999-04-03")
	fails("99-04-03")
	fails("9-04-03")

	// Invalid month
	fails("2019-2-11")
	fails("2019--11")
	fails("2019-XX-11")
	fails("2019-00-11")

	// Invalid day
	fails("2019-04-9")
	fails("2019-04-XX")
	fails("2019-04-")
	fails("2019-04-00")

	succeeds("0123-01-01", 123, time.January, 1)
	succeeds("2000-01-01", 2000, time.January, 1)
	succeeds("2000-02-29", 2000, time.February, 29)
	succeeds("2004-02-29", 2004, time.February, 29)
	succeeds("2019-01-01", 2019, time.January, 1)
	succeeds("2319-12-31", 2319, time.December, 31)

	// Roll over to next month
	succeeds("2005-02-29", 2005, time.March, 1)
	succeeds("2005-01-33", 2005, time.February, 2)
}

func (suite *DateSuite) TestParseDateStart() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseDateStart(input)
		suite.AssertMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseDateStart(input)
		suite.Error(err)
	}

	// TestParseDate runs through all formats, so just make sure failure bubbles up.
	fails("")
	fails("not valid")

	succeeds("0123-01-01", 123, time.January, 1)
	succeeds("2000-01-01", 2000, time.January, 1)
	succeeds("2000-02-29", 2000, time.February, 29)
	succeeds("2004-02-29", 2004, time.February, 29)
	succeeds("2319-12-31", 2319, time.December, 31)
}

func (suite *DateSuite) TestParseDateStartIn() {
	succeeds := func(input string, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseDateStartIn(input, loc)
		suite.AssertMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseDateStartIn(input, loc)
		suite.Error(err)
	}

	// TestParseDate runs through all formats, so just make sure failure bubbles up.
	fails("", locationEDT)
	fails("not valid", locationPDT)

	succeeds("0123-01-01", 123, time.January, 1, time.UTC)
	succeeds("0123-01-01", 123, time.January, 1, locationEDT)
	succeeds("0123-01-01", 123, time.January, 1, locationPDT)

	succeeds("2000-01-01", 2000, time.January, 1, time.UTC)
	succeeds("2000-01-01", 2000, time.January, 1, locationEDT)
	succeeds("2000-01-01", 2000, time.January, 1, locationPDT)

	succeeds("2000-02-29", 2000, time.February, 29, time.UTC)
	succeeds("2000-02-29", 2000, time.February, 29, locationEDT)
	succeeds("2000-02-29", 2000, time.February, 29, locationPDT)

	succeeds("2004-02-29", 2004, time.February, 29, time.UTC)
	succeeds("2004-02-29", 2004, time.February, 29, locationEDT)
	succeeds("2004-02-29", 2004, time.February, 29, locationPDT)

	succeeds("2319-12-31", 2319, time.December, 31, time.UTC)
	succeeds("2319-12-31", 2319, time.December, 31, locationEDT)
	succeeds("2319-12-31", 2319, time.December, 31, locationPDT)
}

func (suite *DateSuite) TestParseDateEnd() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseDateEnd(input)
		suite.AssertAlmostMidnightUTC(date, err, year, month, day)
	}
	fails := func(input string) {
		_, err := isodates.ParseDateEnd(input)
		suite.Error(err)
	}

	// TestParseDate runs through all formats, so just make sure failure bubbles up.
	fails("")
	fails("not valid")

	succeeds("0123-01-01", 123, time.January, 1)
	succeeds("2000-01-01", 2000, time.January, 1)
	succeeds("2000-02-29", 2000, time.February, 29)
	succeeds("2004-02-29", 2004, time.February, 29)
	succeeds("2319-12-31", 2319, time.December, 31)
}

func (suite *DateSuite) TestParseDateEndIn() {
	succeeds := func(input string, year int, month time.Month, day int, loc *time.Location) {
		date, err := isodates.ParseDateEndIn(input, loc)
		suite.AssertAlmostMidnightIn(date, err, year, month, day, loc)
	}
	fails := func(input string, loc *time.Location) {
		_, err := isodates.ParseDateEndIn(input, loc)
		suite.Error(err)
	}

	// TestParseDate runs through all formats, so just make sure failure bubbles up.
	fails("", locationEDT)
	fails("not valid", locationPDT)

	succeeds("0123-01-01", 123, time.January, 1, time.UTC)
	succeeds("0123-01-01", 123, time.January, 1, locationEDT)
	succeeds("0123-01-01", 123, time.January, 1, locationPDT)

	succeeds("2000-01-01", 2000, time.January, 1, time.UTC)
	succeeds("2000-01-01", 2000, time.January, 1, locationEDT)
	succeeds("2000-01-01", 2000, time.January, 1, locationPDT)

	succeeds("2000-02-29", 2000, time.February, 29, time.UTC)
	succeeds("2000-02-29", 2000, time.February, 29, locationEDT)
	succeeds("2000-02-29", 2000, time.February, 29, locationPDT)

	succeeds("2004-02-29", 2004, time.February, 29, time.UTC)
	succeeds("2004-02-29", 2004, time.February, 29, locationEDT)
	succeeds("2004-02-29", 2004, time.February, 29, locationPDT)

	succeeds("2319-12-31", 2319, time.December, 31, time.UTC)
	succeeds("2319-12-31", 2319, time.December, 31, locationEDT)
	succeeds("2319-12-31", 2319, time.December, 31, locationPDT)
}

func ExampleParseDate() {
	date, err := isodates.ParseDate("2019-02-24")
	if err != nil {
		fmt.Printf("oops: %v\n", err)
	}
	fmt.Println(date.Format("Jan 2, 2006"))

	// Output: Feb 24, 2019
}

// BenchmarkParseDate typically runs about 70-80ns/op on a 2014 MacBook Pro
func BenchmarkParseDate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = isodates.ParseDate("2019-02-27")
	}
}
