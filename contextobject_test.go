package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func Test_contextObject_create(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "context($$.context.object).path($.context.object)", result.String())
}

func Test_contextObject_Key(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$$.context.object", result.Key())
}

func Test_contextObject_JSONPath(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$.context.object", result.JSONPath())
}
