package core

import (
	"fmt"
	"time"
)

func GetDurationFromTimestamp(ts string) (duration *time.Duration, err error) {
	timeValue, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	diff := timeValue.Sub(now)
	if diff <= 0 {
		return nil, fmt.Errorf("TimestampError: the timestamp must be a date in the future")
	}
	duration = &diff
	return
}
