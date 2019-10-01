# isodates

The package `isodates` helps you convert ISO 8601 formatted date
strings into actual `time.Time{}` instances. Currently, `isodates`
supports the following formats:

* Date (e.g. "2019-05-23")
* Month-Day (e.g. "--12-25")
* Year-Month (e.g. "2019-04")
* Week (e.g. "2019-W05")
* Week-Day (e.g. "2019-W05-3")

### Basic Usage

Parsing inputs typically gives you the exact date components encoded
in the string. See "Start/End Dates" for information on how to get
exact dates in a specific range.

```go
month, day, err := isodates.ParseMonthDay("--12-25")
if err != nil {
	// do something nifty
}
christmas2019 := time.Date(2019, month, day, 0, 0, 0, 0, time.UTC)
```

### Start/End Dates

Some of the formats that `isodates` supports actually represent
ranges of date rather than a single date. For instance, the ISO
week "2019-W02" represents the entire range of Jan 7, 2019 to
Jan 13, 2019. The `isodates` package contains additional functions
that will fetch exactly midnight on the first day of the range and
others that will fetch 11:59:59pm on the last day of the range.

```go
isoWeekString := "2019-W02"
startDate, err := isodates.ParseWeekStart(isoWeekString)
if err != nil {
	// Handle error
}
endDate, err := isodates.ParseWeekEnd(isoWeekString)
if err != nil {
	// Handle error
}

// Outputs: "Start=Jan 7, 2019 End=Jan 13, 2019"
format := "Jan 02, 2006"
fmt.Printf("Start=%s, End=%s", startDate.Format(format), endDate.Format(format))
```

### Start/End Dates (Local Time)

All of the start/end helpers have a variant ending with `In` that also
takes a `*time.Location`. In these cases, the resulting date/time will
be either midnight or 11:59:59pm in the specified time zone.

```go
edt, _ := time.LoadLocation("America/New_York") 
isoWeekString := "2019-W02"
startDate, err := isodates.ParseWeekStartIn(isoWeekString, edt)
if err != nil {
	// Handle error
}
endDate, err := isodates.ParseWeekEndIn(isoWeekString, edt)
if err != nil {
	// Handle error
}

// Outputs: "Start=Jan 7, 2019 End=Jan 13, 2019"
format := "Jan 02, 2006"
fmt.Printf("Start=%s, End=%s", startDate.Format(format), endDate.Format(format))
```

### Motivation

While parsing ISO 8601 formatted dates is fairly general-purpose, my goal
was to centralize the logic associated w/ parsing Date slots when building
Alexa skills in Go/Lambda. When a user utters "what should I wear next week"
the Alexa skills API will feed an ISO week string representing the range of
next week. `isodates` tries to support all of the various date strings that
will get thrown at you. The start/end helpers will make it easier for your
to build dates ranges from the raw slot data you received.