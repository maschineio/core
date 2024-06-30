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

func TestGetValueTypeInt8(t *testing.T) {
	v, err := core.GetTypedValue(int8(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeInt16(t *testing.T) {
	v, err := core.GetTypedValue(int16(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeInt32(t *testing.T) {
	v, err := core.GetTypedValue(int32(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeInt64(t *testing.T) {
	v, err := core.GetTypedValue(int64(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeFloat32(t *testing.T) {
	v, err := core.GetTypedValue(float32(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeFloat64(t *testing.T) {
	v, err := core.GetTypedValue(float64(42))
	assert.Nil(t, err)
	assert.Equal(t, "Float", v.Type().String()) // we are dealing with JSON float all the time
}

func TestGetValueTypeNil(t *testing.T) {
	v, err := core.GetTypedValue(nil)
	assert.Nil(t, err)
	assert.Equal(t, "Nil", v.Type().String())
}

func TestGetValueTypeString(t *testing.T) {
	v, err := core.GetTypedValue("Zaphod Beeblebrox")
	assert.Nil(t, err)
	assert.Equal(t, "String", v.Type().String())
}

func TestGetValueTypeValuePtr(t *testing.T) {
	v, err := core.GetTypedValue(core.NewBoolValue(true))
	assert.Nil(t, err)
	assert.Equal(t, "Bool", v.Type().String())
}

func TestGetValueTypeValue(t *testing.T) {
	v, err := core.GetTypedValue(*core.NewBoolValue(true))
	assert.Nil(t, err)
	assert.Equal(t, "Bool", v.Type().String())
}

func TestGetValueTypeUnknown(t *testing.T) {
	v, err := core.GetTypedValue(*core.NewUnknownValue())
	assert.Nil(t, err)
	assert.Equal(t, "Unknown", v.Type().String())
}

func TestGetValueTypeError(t *testing.T) {
	_, err := core.GetTypedValue(make(chan int))
	assert.NotNil(t, err)
	assert.EqualError(t, err, "GetTypedValue: unknown type chan int")
}

func TestValueType(t *testing.T) {
	value := core.NewBoolValue(true)
	valid := value.IsValid()
	assert.Equal(t, core.Type(3), value.Type())
	assert.True(t, valid)
}
