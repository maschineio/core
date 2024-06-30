package context_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/context"
)

func TestContextSetGet(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("test", 42)
	result, exists := ctx.Get("test")

	assert.NotNil(t, result)
	assert.True(t, exists)
	assert.Equal(t, 42, result)
}

func TestContextGetCredentialsNonExist(t *testing.T) {
	ctx := context.Context{}

	result, exists := ctx.GetCredential("test")

	assert.Nil(t, result)
	assert.False(t, exists)
}
