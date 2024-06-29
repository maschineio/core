package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

const contextObject = "$$.context.object"

func TestContextObjectCreate(t *testing.T) {
	result := core.NewContextObject(contextObject)
	assert.NotNil(t, result)
	assert.Equal(t, "context($$.context.object).path($.context.object)", result.String())
}

func TestContextObjectKey(t *testing.T) {
	result := core.NewContextObject(contextObject)
	assert.NotNil(t, result)
	assert.Equal(t, contextObject, result.Key())
}

func TestContextObjectJSONPath(t *testing.T) {
	result := core.NewContextObject(contextObject)
	assert.NotNil(t, result)
	assert.Equal(t, contextObject, result.JSONPath())
}
