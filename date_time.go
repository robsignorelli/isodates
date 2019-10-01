package isodates

import "time"

// ParseDateTime accepts an ISO-formatted date/time string (e.g. "2019-05-22T12:33:53.045Z") and returns the
// exact date and time that it represents.
func ParseDateTime(input string) (time.Time, error) {
	// Since there are a bunch of variants for time zone, offset, and millis/nanos, it's
	// faster to just use the standard library for this.
	return time.Parse(time.RFC3339, input)
}
