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

func TestContextGetCredentialsExist(t *testing.T) {
	ctx := context.Context{}
	ctx.Set(context.CREDENTIALSKEY, map[string]any{"user": "test"})
	result, exists := ctx.GetCredential("user")
	credExists := ctx.CredentialsExists()

	assert.NotNil(t, result)
	assert.True(t, exists)
	assert.True(t, credExists)
	assert.Equal(t, "test", result)
}

func TestContextGetCredentialsExistsFalse(t *testing.T) {
	ctx := context.Context{}
	result := ctx.CredentialsExists()

	assert.False(t, result)
}

func TestContextGetEmptyInput(t *testing.T) {
	ctx := context.Context{}
	result := ctx.GetInput()

	assert.Equal(t, []byte("{}"), result)
}
