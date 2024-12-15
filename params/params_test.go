package params_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/params"
)

func TestGetParam(t *testing.T) {
	// Test case: Parameter exists and is of correct type
	t.Run("Parameter exists and is of correct type", func(t *testing.T) {
		p := map[string]any{"key": 42}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter does not exist
	t.Run("Parameter does not exist", func(t *testing.T) {
		p := map[string]any{}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter exists but is of incorrect type
	t.Run("Parameter exists but is of incorrect type", func(t *testing.T) {
		p := map[string]any{"key": "value"}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter map is nil
	t.Run("Parameter map is nil", func(t *testing.T) {
		var param *params.Parameter
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestGetParamDefault(t *testing.T) {
	// Test case: Parameter exists and is of correct type
	t.Run("Parameter exists and is of correct type", func(t *testing.T) {
		p := map[string]any{"key": 42}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 0)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter does not exist
	t.Run("Parameter does not exist", func(t *testing.T) {
		p := map[string]any{}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 42)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter exists but is of incorrect type
	t.Run("Parameter exists but is of incorrect type", func(t *testing.T) {
		p := map[string]any{"key": "value"}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 42)

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter map is nil
	t.Run("Parameter map is nil", func(t *testing.T) {
		var param *params.Parameter
		result, err := params.GetParamDefault(param, "key", 42)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})
}
