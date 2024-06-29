package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func Test_contextObjectCreate(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "context($$.context.object).path($.context.object)", result.String())
}

func Test_contextObjectKey(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$$.context.object", result.Key())
}

func Test_contextObjectJSONPath(t *testing.T) {
	result := core.NewContextObject("$$.context.object")
	assert.NotNil(t, result)
	assert.Equal(t, "$.context.object", result.JSONPath())
}
