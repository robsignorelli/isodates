# isodates

[![Go Report Card](https://goreportcard.com/badge/github.com/robsignorelli/isodates)](https://goreportcard.com/report/github.com/robsignorelli/isodates)

The package `isodates` helps you convert ISO 8601 formatted date
strings into actual `time.Time{}` instances. Currently, `isodates`
supports the following formats:

* Date (e.g. "2019-05-23")
* Date-Time (e.g. "2019-05-23T04:44:33.999Z")
* Month-Day (e.g. "--12-25")
* Year-Month (e.g. "2019-04")
* Week (e.g. "2019-W05")
* Week-Day (e.g. "2019-W05-3")

### Basic Usage

For starters, import `isodates` into your code:

```
import "github.com/robsignorelli/isodates"
```

Parsing inputs gives you the exact date components encoded
in the string. These can usually be fed directly to `time.Date()` with
the exception of week parsing. See the section on "Start/End Dates" to
see how you can obtain a `time.Time` at the very first or last nanosecond
of the parsed date/range in one step.

```
// Simple dates
year, month, day, err := isodates.ParseDate("2019-04-01")

// Month/day values
month, day, err := isodates.ParseMonthDay("--12-25")

// Year/month values
year, month, err := isodates.ParseYearMonth("2019-12")

// ISO Week numbers
year, week, err := isodates.ParseWeek("2019-W11")

// ISO Week numbers w/ day offset
year, week, day, err := isodates.ParseWeek("2019-W11-3")

// Date/time timestamps (already a time.Time)
dateTime, err := isodates.ParseDateTime("2019-03-04T16:04:44.45678Z")
```

### Start/End Dates

Standard `isodates` parser functions just give you the raw components encoded
in the input string. Normally you need to feed those to `time.Date(...)` manually.
Most of the `ParseXyz()` functions, however, have a `ParseXyzStart()` and a `ParseXyzEnd()`
variant that will return a fully-constructed `time.Time` representing
the first and last nanosecond of the parsed range, respectively. These ranges
are in UTC. If you need local times, see the next section.

These are a convenience so that you can easily build date/time ranges that
encapsulate the entire block of time represented by the input. 

```
// Jan 6, 2019 12:00:00AM - Jan 12, 2019 11:59:59PM
weekStart, err := isodates.ParseWeekStart("2019-W02")
weekEnd, err := isodates.ParseWeekEnd("2019-W02")

// Feb 1, 2000 12:00:00AM - Feb 29, 2000 11:59:59PM
febStart, err := isodates.ParseYearMonthStart("2000-02")
febEnd, err := isodates.ParseYearMonthEnd("2000-02")
```

### Start/End Dates (Local Time)

All of the Start/End helpers have a variant ending with `In` that also
takes a `*time.Location`. In these cases, the resulting date/time will
be either midnight or 11:59:59pm in the specified time zone.

```
ny, _ := time.LoadLocation("America/New_York")

// Jan 6, 2019 12:00:00AM - Jan 12, 2019 11:59:59PM
weekStartNY, err := isodates.ParseWeekStartIn("2019-W02", ny)
weekEndNY, err := isodates.ParseWeekEndIn("2019-W02", ny)

// Feb 1, 2000 12:00:00AM - Feb 29, 2000 11:59:59PM
febStartNY, err := isodates.ParseYearMonthStartIn("2000-02", ny)
febEndNY, err := isodates.ParseYearMonthEndIn("2000-02", ny)
```

### Motivation

While parsing ISO 8601 formatted dates is fairly general-purpose, my goal
was to centralize the logic associated w/ parsing Date slots when building
Alexa skills in Go/Lambda. When a user utters "what should I wear next week"
the Alexa skills API will feed an ISO week string representing the range of
next week. `isodates` tries to support all of the various date strings that
will get thrown at you. The start/end helpers will make it easier for your
to build date ranges from the raw slot data you received.