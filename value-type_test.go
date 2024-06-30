package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestValueTypeStringMapString(t *testing.T) {
	v := core.NewStringMapValue(map[string]any{})
	assert.Equal(t, "StringMap", v.Type().String())
}

//func TestDataTypeError(t *testing.T) {
//	testCases := []token.DataType{
//		-1, // less than 0
//	}
//	result := testCases[0].String()
//	assert.Equal(t, "DataType(-1)", result)
//}

func TestGetValueTypeByteSlice(t *testing.T) {
	v, err := core.GetTypedValue([]byte("hello"))
	assert.Nil(t, err)
	assert.Equal(t, "Bytes", v.Type().String())
}

func TestGetValueTypeBool(t *testing.T) {
	v, err := core.GetTypedValue(false)
	assert.Nil(t, err)
	assert.Equal(t, "Bool", v.Type().String())
}

func TestGetValueTypeStringMap(t *testing.T) {
	v, err := core.GetTypedValue(map[string]any{"key": "value"})
	assert.Nil(t, err)
	assert.Equal(t, "StringMap", v.Type().String())
}

func TestGetValueTypePtrStringMap(t *testing.T) {
	v, err := core.GetTypedValue(&map[string]any{"key": "value"})
	assert.Nil(t, err)
	assert.Equal(t, "PointerStringMap", v.Type().String())
}

func TestGetValueTypeSliceVal(t *testing.T) {
	v, err := core.GetTypedValue([]any{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, "Slice", v.Type().String())
}

func TestGetValueTypeInt(t *testing.T) {
	v, err := core.GetTypedValue(int(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}
