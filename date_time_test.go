package isodates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/isodates"
	"github.com/stretchr/testify/suite"
)

func TestDateTimeSuite(t *testing.T) {
	suite.Run(t, new(DateTimeSuite))
}

type DateTimeSuite struct {
	suite.Suite
}

func (suite *DateTimeSuite) TestParseDateTime() {
	zoneOffset := func(date time.Time) int {
		_, offset := date.Zone()
		return offset
	}
	succeeds := func(input string, year int, month time.Month, day, hour, minute, second, nanos, offset int) {
		date, err := isodates.ParseDateTime(input)
		_ = suite.NoError(err) &&
			suite.Equal(year, date.Year()) &&
			suite.Equal(month, date.Month()) &&
			suite.Equal(day, date.Day()) &&
			suite.Equal(hour, date.Hour()) &&
			suite.Equal(minute, date.Minute()) &&
			suite.Equal(second, date.Second()) &&
			suite.Equal(nanos, date.Nanosecond()) &&
			suite.Equal(offset, zoneOffset(date))
	}
	fails := func(input string) {
		_, err := isodates.ParseDateTime(input)
		suite.Error(err)
	}
	fails("")
	fails("not valid")
	fails("------")
	fails("01-2019-21")
	fails("2019/01/02T12-33-44Z")

	// Invalid year
	fails("$G33-04-03T06:44:33Z")
	fails("-04-03T06:44:33Z")
	fails("999-04-03T06:44:33Z")
	fails("99-04-03T06:44:33Z")
	fails("9-04-03T06:44:33Z")

	// Invalid month
	fails("2019-$4-03T06:44:33Z")
	fails("2019-4-03T06:44:33Z")
	fails("2019-0-03T06:44:33Z")
	fails("2019-XX-03T06:44:33Z")

	// Invalid day
	fails("2019-03-$4T06:44:33Z")
	fails("2019-03-4T06:44:33Z")
	fails("2019-03-0T06:44:33Z")
	fails("2019-03-XXT06:44:33Z")

	// Invalid hour
	fails("2019-03-04T-1:44:33Z")
	fails("2019-03-04T44:44:33Z")

	// Invalid minute
	fails("2019-03-04T06:-4:33Z")
	fails("2019-03-04T06:X4:33Z")
	fails("2019-03-04T06:77:33Z")

	// Invalid seconds
	fails("2019-03-04T06:04:-3Z")
	fails("2019-03-04T06:04:3Z")
	fails("2019-03-04T06:04:77Z")

	// Invalid nanos
	fails("2019-03-04T06:04:33.Z")
	fails("2019-03-04T06:04:33.-4Z")
	fails("2019-03-04T06:04:33.9999999999Z") // one too many digits

	// Ensure year padding is handled
	succeeds("2019-03-04T06:04:44Z", 2019, time.March, 4, 6, 4, 44, 0, 0)
	succeeds("0119-03-04T16:04:44Z", 119, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("0019-03-04T16:04:44Z", 19, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("0009-03-04T16:04:44Z", 9, time.March, 4, 16, 4, 44, 0, 0)

	// Hour after 12/noon
	succeeds("2019-03-04T16:04:44Z", 2019, time.March, 4, 16, 4, 44, 0, 0)

	// Millis/nanos
	succeeds("2019-03-04T16:04:44.0Z", 2019, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("2019-03-04T16:04:44.000Z", 2019, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("2019-03-04T16:04:44.1Z", 2019, time.March, 4, 16, 4, 44, 100000000, 0)
	succeeds("2019-03-04T16:04:44.001Z", 2019, time.March, 4, 16, 4, 44, 1000000, 0)
	succeeds("2019-03-04T16:04:44.0002Z", 2019, time.March, 4, 16, 4, 44, 200000, 0)
	succeeds("2019-03-04T16:04:44.999999999Z", 2019, time.March, 4, 16, 4, 44, 999999999, 0)

	// Offsets
	succeeds("2019-03-04T16:04:44.000+00:00", 2019, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("2019-03-04T16:04:44.000-00:00", 2019, time.March, 4, 16, 4, 44, 0, 0)
	succeeds("2019-03-04T16:04:44.000+07:00", 2019, time.March, 4, 16, 4, 44, 0, 25200)
	succeeds("2019-03-04T16:04:44.000+07:30", 2019, time.March, 4, 16, 4, 44, 0, 27000)
	succeeds("2019-03-04T16:04:44.000-07:00", 2019, time.March, 4, 16, 4, 44, 0, -25200)
	succeeds("2019-03-04T16:04:44.000-07:30", 2019, time.March, 4, 16, 4, 44, 0, -27000)
}

func ExampleParseDateTime() {
	date, err := isodates.ParseDateTime("2019-02-24T06:44:33Z")
	if err != nil {
		fmt.Printf("oops: %v\n", err)
	}
	fmt.Println(date.Format("Jan 2, 2006 3:04PM"))

	// Output: Feb 24, 2019 6:44AM
}

// BenchmarkParseDateTime typically runs about 230-250ns/op on a 2014 MacBook Pro
func BenchmarkParseDateTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = isodates.ParseDateTime("2019-02-27T06:44:33Z")
	}
}
