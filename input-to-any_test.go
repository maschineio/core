package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestCoreInputToBytesWithMaptInput(t *testing.T) {
	testCase := map[string]any{"test": 42}
	result, err := core.InputToBytes(testCase)
	assert.Nil(t, err)
	assert.Equal(t, []uint8([]byte{0x7b, 0x22, 0x74, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x34, 0x32, 0x7d}), result)
}

func TestCoreInputToBytesWithPointerToBytesError(t *testing.T) {
	testCase := []byte(`{"test": 42}`)
	_, err := core.InputToBytes(&testCase)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "core: unknown input type: got *[]uint8")
}
