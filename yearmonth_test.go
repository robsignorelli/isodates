package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestYearMonthSuite(t *testing.T) {
	suite.Run(t, new(YearMonthSuite))
}

type YearMonthSuite struct {
	suite.Suite
}

func (suite *YearMonthSuite) TestParseYearMonth() {
	succeeds := func(input string, expectedYear int, expectedMonth time.Month) {
		year, month, err := isodates.ParseYearMonth(input)
		_ = suite.NoError(err) &&
			suite.Equal(expectedYear, year) &&
			suite.Equal(expectedMonth, month)
	}
	fails := func(input string) {
		_, _, err := isodates.ParseYearMonth(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("01-2019")
	fails("2019/01")

	fails("2019--1")    // -1 not a valid month
	fails("2019-13")    // 13 not a valid month
	fails("2019-00")    // 00 not a valid month
	fails("2019-")      // month must be padded
	fails("2019-5")     // month must be padded
	fails("123-05")     // year must be padded
	fails("23-05")      // year must be padded
	fails("3-05")       // year must be padded
	fails("2019-O1")    // It's an "oh" not a zero...
	fails("02019-01")   // Only allow + or - as first char in BC/AD prefixed variant
	fails("2019--01")   // Only allow + or - as first char in BC/AD prefixed variant
	fails("_2019-01")   // Only allow + or - as first char in BC/AD prefixed variant
	fails("x2019-01")   // Only allow + or - as first char in BC/AD prefixed variant
	fails("2019-01-03") // Good ISO date. Not a good ISO year/month.

	succeeds("2000-01", 2000, time.January)
	succeeds("2000-11", 2000, time.November)
	succeeds("2019-11", 2019, time.November)
	succeeds("1215-06", 1215, time.June)
	succeeds("1215-06", 1215, time.June)
	succeeds("0123-12", 123, time.December)
	succeeds("0012-12", 12, time.December)
	succeeds("0001-12", 1, time.December)
	succeeds("0001-01", 1, time.January)
	succeeds("0000-01", 0, time.January)

	succeeds("+2000-01", 2000, time.January)
	succeeds("+2000-11", 2000, time.November)
	succeeds("+2019-11", 2019, time.November)
	succeeds("+1215-06", 1215, time.June)
	succeeds("+1215-06", 1215, time.June)
	succeeds("+0123-12", 123, time.December)
	succeeds("+0012-12", 12, time.December)
	succeeds("+0001-12", 1, time.December)
	succeeds("+0001-01", 1, time.January)
	succeeds("+0000-01", 0, time.January)

	succeeds("-2000-01", -2000, time.January)
	succeeds("-2000-11", -2000, time.November)
	succeeds("-2019-11", -2019, time.November)
	succeeds("-1215-06", -1215, time.June)
	succeeds("-1215-06", -1215, time.June)
	succeeds("-0123-12", -123, time.December)
	succeeds("-0012-12", -12, time.December)
	succeeds("-0001-12", -1, time.December)
	succeeds("-0001-01", -1, time.January)
	succeeds("-0000-01", 0, time.January)
}

func ExampleParseYearMonth() {
	year, month, err := isodates.ParseYearMonth("2019-01")
	fmt.Println(fmt.Sprintf("%d %d %v", year, month, err == nil))

	// Output: 2019 1 true
}

// BenchmarkParseYearMonth typically runs about 35-40ns/op on a 2014 MacBook Pro
func BenchmarkParseYearMonth(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _ = isodates.ParseYearMonth("2019-04")
	}
}
