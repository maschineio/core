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

func TestContextSetInput(t *testing.T) {
	ctx := context.Context{}
	ctx.SetInput([]byte("test"))
	result := ctx.GetInput()

	assert.Equal(t, []byte("test"), result)
}

func TestContextSetInputNotBytes(t *testing.T) {
	ctx := context.Context{}
	ctx.Set(context.INPUTKEY, 42)
	result := ctx.GetInput()

	assert.Equal(t, []uint8([]byte{0x7b, 0x7d}), result)
}

//func TestContextGetValue(t *testing.T) {
//	ctx := context.Context{}
//	ctx.Set("key", 42)
//	result := ctx.Value("key")
//	result2 := ctx.Value("not-exists")
//
//	assert.Equal(t, 42, result)
//	assert.Nil(t, result2)
//}
