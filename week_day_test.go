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
	suite.Suite
}

func (suite *WeekDaySuite) TestParseWeekDay() {
	succeeds := func(input string, year int, month time.Month, day int) {
		date, err := isodates.ParseWeekDay(input)
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
		_, err := isodates.ParseWeekDay(input)
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

func ExampleParseWeekDay() {
	date, err := isodates.ParseWeekDay("2019-W02-2")
	if err != nil {
		fmt.Printf("oops: %v\n", err)
	}
	fmt.Println(date.Format("Jan 2, 2006"))

	// Output: Jan 8, 2019
}

// BenchmarkParseWeekDay typically runs about 35-40ns/op on a 2014 MacBook Pro
func BenchmarkParseWeekDay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = isodates.ParseWeekDay("2019-W04-3")
	}
}
