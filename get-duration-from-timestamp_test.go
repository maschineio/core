package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func Test_GetDurationFromTimestampEmptyTimestamp(t *testing.T) {
	ts := ""
	result, err := core.GetDurationFromTimestamp(ts)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	// assert.Equal(t, "context($$.context.object).path($.context.object)", result.String())
}

//

func Test_GetDurationFromTimestampValidTimestamp(t *testing.T) {
	ts := "2050-01-02T15:04:05Z"
	result, err := core.GetDurationFromTimestamp(ts)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDurationFromTimestampTimeIsExpired(t *testing.T) {
	ts := "2023-08-03T15:04:05Z"
	result, err := core.GetDurationFromTimestamp(ts)
	assert.NotNil(t, err)
	assert.Equal(t, "TimestampError: the timestamp must be a date in the future", err.Error())
	assert.Nil(t, result)
}
