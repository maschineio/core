package core

import (
	"time"
)

// convertTimestamp converts a timestamp string to go time
func GetTimestamp(timestamp string) (tstamp *time.Time, ok bool) {
	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return tstamp, false
	}
	return &ts, true
}

// IsTimestamp validates a timestap string (conforms to RFC3339)
func IsTimestamp(timestamp string) bool {
	_, err := time.Parse(time.RFC3339, timestamp)
	return err == nil
}
