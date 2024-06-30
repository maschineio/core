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

func TestUnknownValueString(t *testing.T) {
	value := core.NewUnknownValue()
	assert.Equal(t, "unknown", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []uint8([]byte{0x6e, 0x75, 0x6c, 0x6c}), result) // represents json null
}

func TestNilValueString(t *testing.T) {
	value := core.NewNilValue()
	assert.Equal(t, "<nil>", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []uint8([]byte{0x6e, 0x75, 0x6c, 0x6c}), result) // represents json null
}

func TestStringValueString(t *testing.T) {
	value := core.NewStringValue("hello")
	assert.Equal(t, "hello", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello"), result)
}

func TestBoolValueString(t *testing.T) {
	value := core.NewBoolValue(true)
	assert.Equal(t, "true", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte("true"), result)
}

func TestBytesValueString(t *testing.T) {
	value := core.NewBytesValue([]byte("test"))
	assert.Equal(t, "test", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte("test"), result)
}

func TestFloatValueString(t *testing.T) {
	value := core.NewFloat32Value(float32(32))
	assert.Equal(t, "32", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte("32"), result)
}

func TestStringMapValueString(t *testing.T) {
	value := core.NewStringMapValue(map[string]any{"test": 42})
	assert.Equal(t, "map[test:42]", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`{"test":42}`), result)
}

func TestPointerStringMapValueString(t *testing.T) {
	value := core.NewPointerStringMapValue(&map[string]any{"test": 42})
	assert.Equal(t, "map[test:42]", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`{"test":42}`), result)
}

func TestPointerStringMapValueStringNil(t *testing.T) {
	value := core.NewPointerStringMapValue(nil)
	assert.Equal(t, "unknown", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`null`), result)
}

func TestSliceValueString(t *testing.T) {
	value := core.NewSliceValue([]any{42})
	assert.Equal(t, "[42]", value.String())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`[42]`), result)
}

func TestBoolValue(t *testing.T) {
	value := core.NewBoolValue(true)
	assert.True(t, value.BoolValue())

	result, err := value.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`true`), result)
}
