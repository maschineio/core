package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestContextObjectCreate(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "context($$.context.object).path($.context.object)", result.String())
}

func TestContextObjectKey(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$$.context.object", result.Key())
}

func TestContextObjectJSONPath(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$.context.object", result.JSONPath())
}
