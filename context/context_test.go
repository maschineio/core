package context_test

import (
	"testing"
	"time"

	stdcontext "context"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"maschine.io/core/context"
)

func TestBackground(t *testing.T) {
	ctx := context.Background()

	assert.NotNil(t, ctx)
	assert.Empty(t, ctx.Keys)
}

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

func TestContextGetValue(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", 42)
	result := ctx.Value("key")
	result2 := ctx.Value("not-exists")

	assert.Equal(t, 42, result)
	assert.Nil(t, result2)
}

func TestContextMustGet(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", "value")
	result := ctx.MustGet("key")

	assert.Equal(t, "value", result)

	assert.Panics(t, func() {
		ctx.MustGet("not-exists")
	})
}

func TestContextGetString(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", "value")
	result := ctx.GetString("key")

	assert.Equal(t, "value", result)
}

func TestContextGetStringWithDefault(t *testing.T) {
	ctx := context.Context{}
	result := ctx.GetStringWithDefault("key", "default")

	assert.Equal(t, "default", result)
}

func TestContextGetInt(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", 42)
	result := ctx.GetInt("key")

	assert.Equal(t, 42, result)
}

func TestContextGetBool(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", true)
	result := ctx.GetBool("key")

	assert.True(t, result)
}

func TestContextGetFloat64(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", 42.42)
	result := ctx.GetFloat64("key")

	assert.Equal(t, 42.42, result)
}

func TestContextGetTime(t *testing.T) {
	ctx := context.Context{}
	now := time.Now()
	ctx.Set("key", now)
	result := ctx.GetTime("key")

	assert.Equal(t, now, result)
}

func TestContextGetDuration(t *testing.T) {
	ctx := context.Context{}
	duration := time.Duration(42)
	ctx.Set("key", duration)
	result := ctx.GetDuration("key")

	assert.Equal(t, duration, result)
}

func TestContextGetStringSlice(t *testing.T) {
	ctx := context.Context{}
	slice := []string{"a", "b", "c"}
	ctx.Set("key", slice)
	result := ctx.GetStringSlice("key")

	assert.Equal(t, slice, result)
}

func TestContextGetStringMap(t *testing.T) {
	ctx := context.Context{}
	m := map[string]any{"a": 1, "b": 2}
	ctx.Set("key", m)
	result := ctx.GetStringMap("key")

	assert.Equal(t, m, result)
}

func TestContextGetStringMapString(t *testing.T) {
	ctx := context.Context{}
	m := map[string]string{"a": "1", "b": "2"}
	ctx.Set("key", m)
	result := ctx.GetStringMapString("key")

	assert.Equal(t, m, result)
}

func TestContextGetStringMapStringSlice(t *testing.T) {
	ctx := context.Context{}
	m := map[string][]string{"a": {"1", "2"}, "b": {"3", "4"}}
	ctx.Set("key", m)
	result := ctx.GetStringMapStringSlice("key")

	assert.Equal(t, m, result)
}
func TestContextGetBytes(t *testing.T) {
	ctx := context.Context{}
	data := []byte("test")
	ctx.Set("key", data)
	result := ctx.GetBytes("key")

	assert.Equal(t, data, result)
}

func TestContextGetInt64(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", int64(42))
	result := ctx.GetInt64("key")

	assert.Equal(t, int64(42), result)
}

func TestContextGetUint(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", uint(42))
	result := ctx.GetUint("key")

	assert.Equal(t, uint(42), result)
}

func TestContextGetUint64(t *testing.T) {
	ctx := context.Context{}
	ctx.Set("key", uint64(42))
	result := ctx.GetUint64("key")

	assert.Equal(t, uint64(42), result)
}

func TestContextGetInputAsInterface(t *testing.T) {
	ctx := context.Context{}
	ctx.SetInput([]byte(`{"key":"value"}`))
	result, err := ctx.GetInputAsInterface()

	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"key":"value"}`), result)
}

func TestContextGetInputAsMap(t *testing.T) {
	ctx := context.Context{}
	ctx.SetInput([]byte(`{"key":"value"}`))
	result, err := ctx.GetInputAsMap()

	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"key": "value"}, result)
}

func TestContextSetLogger(t *testing.T) {
	ctx := context.Context{}
	logger, _ := zap.NewProduction()
	ctx.SetLogger(logger)
	result := ctx.DefaultLogger()

	assert.Equal(t, logger, result)
}

func TestContextGetLogger(t *testing.T) {
	ctx := context.Context{}
	logger, _ := zap.NewProduction()
	ctx.SetLogger(logger)
	result := ctx.GetLogger(context.LOGGERKEY)

	assert.Equal(t, logger, result)
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	assert.NotNil(t, ctx)
	assert.NotNil(t, cancel)

	cancel() // Abbruch auslösen

	select {
	case <-ctx.Done():
		assert.Equal(t, stdcontext.Canceled, ctx.Err())
	case <-time.After(1 * time.Second):
		t.Fatal("context was not canceled")
	}
}

func TestContextWithDeadline(t *testing.T) {
	parent := context.Background()
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(parent, deadline)

	assert.NotNil(t, ctx)
	assert.NotNil(t, cancel)

	time.Sleep(2 * time.Second)
	assert.Equal(t, stdcontext.DeadlineExceeded, ctx.Err())
}

func TestContextWithTimeout(t *testing.T) {
	parent := context.Background()
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(parent, timeout)

	assert.NotNil(t, ctx)
	assert.NotNil(t, cancel)

	time.Sleep(2 * time.Second)
	assert.Equal(t, stdcontext.DeadlineExceeded, ctx.Err())
}

func TestContextWithValue(t *testing.T) {
	parent := context.Background()
	ctx := context.WithValue(parent, "key", "value")

	assert.NotNil(t, ctx)
	assert.Equal(t, "value", ctx.Value("key"))
}
