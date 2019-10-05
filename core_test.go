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

func (suite ChronoSuite) AssertTimeIn(date time.Time, err error, year int, month time.Month, day, hour, minute, second, nanos int, loc *time.Location) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(hour, date.Hour(), "incorrect hour") &&
		suite.Equal(minute, date.Minute(), "incorrect minute") &&
		suite.Equal(second, date.Second(), "incorrect second") &&
		suite.Equal(nanos, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(loc, date.Location(), "incorrect location")
}

func (suite ChronoSuite) AssertTimeUTC(date time.Time, err error, year int, month time.Month, day, hour, minute, second, nanos int) bool {
	return suite.NoError(err) &&
		suite.Equal(year, date.Year(), "incorrect year") &&
		suite.Equal(month, date.Month(), "incorrect month") &&
		suite.Equal(day, date.Day(), "incorrect day") &&
		suite.Equal(hour, date.Hour(), "incorrect hour") &&
		suite.Equal(minute, date.Minute(), "incorrect minute") &&
		suite.Equal(second, date.Second(), "incorrect second") &&
		suite.Equal(nanos, date.Nanosecond(), "incorrect nanos") &&
		suite.Equal(time.UTC, date.Location(), "incorrect location")
}

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
