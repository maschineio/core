package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestCoreIsTypeFalse(t *testing.T) {
	result := core.IsType[bool]("test")
	assert.False(t, result)
}

func TestCoreIsTypeTrue(t *testing.T) {
	result := core.IsType[bool](false)
	assert.True(t, result)
}
