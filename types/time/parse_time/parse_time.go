package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var tzRe = regexp.MustCompile(`\[.+\]`)

// extractLocation willl extract location information from query (e.g
// / "todat [US/Pacific]". It returns the location and trimed query
// without the location.
func extractLocation(query string) (*time.Location, string, error) {
	loc := tzRe.FindStringIndex(query)
	if loc == nil { // No time zone in query, return UTC
		return time.UTC, query, nil
	}

	// +1 to remove []
	locName := query[loc[0]+1 : loc[1]-1]
	tz, err := time.LoadLocation(locName)
	if err != nil {
		return nil, query, err
	}

	// Remove time zone from query
	query = strings.TrimSpace(query[:loc[0]])
	return tz, query, nil
}

func today(loc *time.Location) time.Time {
	t := time.Now().In(loc)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t
}

var unitNames = map[string]time.Duration{
	"minute": time.Minute,
	"hour":   time.Hour,
	"day":    24 * time.Hour,
	"week":   7 * 24 * time.Hour,
}

// parseDelta parses a string like "3 days" and returns the duration.
// (3 * 24 * time.Hour) and the rounding duration. (24 * time.Hour)
func parseDelta(q string) (time.Duration, time.Duration, error) {
	var amount time.Duration
	var unit string
	_, err := fmt.Sscanf(q, "%d %s", &amount, &unit)
	if err != nil {
		return time.Duration(0), time.Duration(0), err
	}

	// weeks -> week
	unit = strings.TrimSuffix(unit, "s")

	d, ok := unitNames[unit]
	if !ok {
		err := fmt.Errorf("unknown duration %q", unit)
		return time.Duration(0), time.Duration(0), err
	}

	// Negative duration since times are offset from now
	return -(amount * d), d, nil
}

// roundTime rounds time to delta frequency.
func roundTime(t time.Time, delta time.Duration) time.Time {
	year, month, day := t.Year(), t.Month(), t.Day()
	hour, minute := t.Hour(), t.Minute()

	switch {
	case delta >= time.Hour:
		minute = 0
		fallthrough
	case delta >= 24*time.Hour:
		minute, hour = 0, 0
	}

	return time.Date(year, month, day, hour, minute, 0, 0, t.Location())
}

// parseTime parses a time query (such as "2 weeks ago", "today GMT" ... )
// and will return the corresponding time.
func parseTime(query string) (time.Time, error) {
	loc, query, err := extractLocation(query)
	if err != nil {
		return time.Time{}, err
	}

	if query == "today" {
		return today(loc), nil
	}

	// Try YYYY-mm-ddTHH:MM first
	t, err := time.ParseInLocation("2006-01-02T15:04", query, loc)
	if err == nil {
		return t, nil
	}

	delta, round, err := parseDelta(query)
	if err != nil {
		return time.Time{}, err
	}

	t = time.Now().In(loc).Add(delta)
	return roundTime(t, round), nil
}

func main() {

}
