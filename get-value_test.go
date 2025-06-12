package core_test

import (
	"testing"

	"maschine.io/core"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestGetValue(t *testing.T) {
	t.Run("CorrectType", func(t *testing.T) {
		val, ok := core.GetValue[int](42)
		assert.True(t, ok)
		assert.Equal(t, 42, val)
	})

	t.Run("IncorrectType", func(t *testing.T) {
		val, ok := core.GetValue[string](42)
		assert.False(t, ok)
		assert.Equal(t, "", val)
	})

	t.Run("StringType", func(t *testing.T) {
		val, ok := core.GetValue[string]("hello")
		assert.True(t, ok)
		assert.Equal(t, "hello", val)
	})
}

func TestGetValuesWithLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	input := map[string]interface{}{
		"strVal": "hello",
		"intVal": 123,
	}
	t.Run("ValidPathsCorrectType", func(t *testing.T) {
		val1, val2, ok := core.GetValuesWithLogger[string](
			"token",
			"$.strVal",
			"$.strVal",
			input,
			logger,
		)
		assert.True(t, ok)
		assert.Equal(t, "hello", val1)
		assert.Equal(t, "hello", val2)
	})

	t.Run("InvalidPath", func(t *testing.T) {
		val1, val2, ok := core.GetValuesWithLogger[string](
			"token",
			"$.doesNotExist",
			"$.strVal",
			input,
			logger,
		)
		assert.False(t, ok)
		assert.Equal(t, "", val1)
		assert.Equal(t, "", val2)
	})

	t.Run("TypeMismatch", func(t *testing.T) {
		val1, val2, ok := core.GetValuesWithLogger[string](
			"token",
			"$.intVal",
			"$.intVal",
			input,
			logger,
		)
		assert.False(t, ok)
		assert.Equal(t, "", val1)
		assert.Equal(t, "", val2)
	})
}

func TestGetTimestampValueWithLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	input := map[string]interface{}{
		"validTime":   "2023-01-02T15:04:05Z",
		"invalidTime": "not-a-time",
	}

	t.Run("ValidTimestamps", func(t *testing.T) {
		val1, val2, ok := core.GetTimestampValueWithLogger(
			"token",
			"$.validTime",
			"$.validTime",
			input,
			logger,
		)
		assert.True(t, ok)
		assert.NotNil(t, val1)
		assert.NotNil(t, val2)
	})

	t.Run("OneInvalidTimestamp", func(t *testing.T) {
		val1, val2, ok := core.GetTimestampValueWithLogger(
			"token",
			"$.validTime",
			"$.invalidTime",
			input,
			logger,
		)
		assert.False(t, ok)
		assert.Nil(t, val1)
		assert.Nil(t, val2)
	})

	t.Run("InvalidPath", func(t *testing.T) {
		val1, val2, ok := core.GetTimestampValueWithLogger(
			"token",
			"$.doesNotExist",
			"$.invalidTime",
			input,
			logger,
		)
		assert.False(t, ok)
		assert.Nil(t, val1)
		assert.Nil(t, val2)
	})
}
