package isodates_test

import (
	"time"

	"github.com/stretchr/testify/suite"
)

var locationEDT, _ = time.LoadLocation("America/New_York")
var locationPDT, _ = time.LoadLocation("America/Los_Angeles")

type ChronoSuite struct {
	suite.Suite
}

// AssertTime ensures that the given date/error result has no error and that the date/time matches all
// of "date's" individual time components. No location is assumed for this test.
func (suite ChronoSuite) AssertTime(date time.Time, err error, year int, month time.Month, day, hour, minute, second, nanos int) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(hour, date.Hour(), "incorrect hour") &&
		suite.Equal(minute, date.Minute(), "incorrect minute") &&
		suite.Equal(second, date.Second(), "incorrect second") &&
		suite.Equal(nanos, date.Nanosecond(), "incorrect nanos")
}

// AssertMidnightUTC ensures that there's no error and that the resulting 'date' is at exactly midnight on
// the given year/month/day specified. The time should be in UTC.
func (suite ChronoSuite) AssertMidnightUTC(date time.Time, err error, year int, month time.Month, day int) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(0, date.Hour(), "incorrect hour") &&
		suite.Equal(0, date.Minute(), "incorrect minute") &&
		suite.Equal(0, date.Second(), "incorrect second") &&
		suite.Equal(0, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(time.UTC, date.Location(), "incorrect location")
}

// AssertMidnightUTC ensures that there's no error and that the resulting 'date' is at exactly midnight on
// the given year/month/day specified. The time should be in the specified location/zone.
func (suite ChronoSuite) AssertMidnightIn(date time.Time, err error, year int, month time.Month, day int, loc *time.Location) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(0, date.Hour(), "incorrect hour") &&
		suite.Equal(0, date.Minute(), "incorrect minute") &&
		suite.Equal(0, date.Second(), "incorrect second") &&
		suite.Equal(0, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(loc, date.Location(), "incorrect location")
}

// AssertAlmostMidnightUTC ensures that there's no error and that the resulting 'date' is at 11:59:59pm on
// the given year/month/day specified. The time should be in UTC.
func (suite ChronoSuite) AssertAlmostMidnightUTC(date time.Time, err error, year int, month time.Month, day int) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(23, date.Hour(), "incorrect hour") &&
		suite.Equal(59, date.Minute(), "incorrect minute") &&
		suite.Equal(59, date.Second(), "incorrect second") &&
		suite.Equal(999999999, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(time.UTC, date.Location(), "incorrect location")
}

// AssertAlmostMidnightUTC ensures that there's no error and that the resulting 'date' is at 11:59:59pm on
// the given year/month/day specified. The time should be in the specified location/zone.
func (suite ChronoSuite) AssertAlmostMidnightIn(date time.Time, err error, year int, month time.Month, day int, loc *time.Location) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(23, date.Hour(), "incorrect hour") &&
		suite.Equal(59, date.Minute(), "incorrect minute") &&
		suite.Equal(59, date.Second(), "incorrect second") &&
		suite.Equal(999999999, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(loc, date.Location(), "incorrect location")
}
