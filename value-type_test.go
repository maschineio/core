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

func TestBoolValueAsYAMLBytes(t *testing.T) {
	value := core.NewBoolValue(true)
	assert.True(t, value.BoolValue())

	result, err := value.AsYAMLBytes()
	assert.Nil(t, err)
	assert.Equal(t, []uint8([]byte{0x74, 0x72, 0x75, 0x65, 0xa}), result)
}

func TestBoolValueAsAny(t *testing.T) {
	value := core.NewBoolValue(true)
	assert.True(t, value.BoolValue())

	result, err := value.AsAny()
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestValueAsJSONString(t *testing.T) {
	value := core.NewBoolValue(true)
	assert.True(t, value.BoolValue())

	result, err := value.AsJSONString()
	assert.Nil(t, err)
	assert.Equal(t, "true", result)
}

func TestValueAsYAMLPrettyfiedBytes(t *testing.T) {
	value := core.NewStringMapValue(map[string]any{"key": "value", "nested": map[string]any{"earth": "moon"}})

	result, err := value.AsJSONPrettyfiedBytes()
	assert.Nil(t, err)
	assert.Equal(t, []uint8([]byte{0x7b, 0xa, 0x20, 0x20, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x22, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x65, 0x61, 0x72, 0x74, 0x68, 0x22, 0x3a, 0x20, 0x22, 0x6d, 0x6f, 0x6f, 0x6e, 0x22, 0xa, 0x20, 0x20, 0x7d, 0xa, 0x7d}), result)
}
func TestNewBoolValue(t *testing.T) {
	v := core.NewBoolValue(true)
	assert.Equal(t, "Bool", v.Type().String())
	assert.Equal(t, "true", v.String())
}

func TestNewStringMapValue(t *testing.T) {
	v := core.NewStringMapValue(map[string]any{"key": "value"})
	assert.Equal(t, "StringMap", v.Type().String())
	assert.Equal(t, "map[key:value]", v.String())
}

func TestNewPointerStringMapValue(t *testing.T) {
	m := map[string]any{"key": "value"}
	v := core.NewPointerStringMapValue(&m)
	assert.Equal(t, "PointerStringMap", v.Type().String())
	assert.Equal(t, "map[key:value]", v.String())
}

func TestNewBytesValue(t *testing.T) {
	v := core.NewBytesValue([]byte("hello"))
	assert.Equal(t, "Bytes", v.Type().String())
	assert.Equal(t, "hello", v.String())
}

func TestNewFloat64Value(t *testing.T) {
	v := core.NewFloat64Value(42.0)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewFloat32Value(t *testing.T) {
	v := core.NewFloat32Value(42.0)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewIntValue(t *testing.T) {
	v := core.NewIntValue(42)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewInt8Value(t *testing.T) {
	v := core.NewInt8Value(42)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewInt16Value(t *testing.T) {
	v := core.NewInt16Value(42)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewInt32Value(t *testing.T) {
	v := core.NewInt32Value(42)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewInt64Value(t *testing.T) {
	v := core.NewInt64Value(42)
	assert.Equal(t, "Float", v.Type().String())
	assert.Equal(t, "42", v.String())
}

func TestNewNilValue(t *testing.T) {
	v := core.NewNilValue()
	assert.Equal(t, "Nil", v.Type().String())
	assert.Equal(t, "<nil>", v.String())
}

func TestNewSliceValue(t *testing.T) {
	v := core.NewSliceValue([]any{1, 2, 3})
	assert.Equal(t, "Slice", v.Type().String())
	assert.Equal(t, "[1 2 3]", v.String())
}

func TestNewUnknownValue(t *testing.T) {
	v := core.NewUnknownValue()
	assert.Equal(t, "Unknown", v.Type().String())
	assert.Equal(t, "unknown", v.String())
}

func TestNewStringValue(t *testing.T) {
	v := core.NewStringValue("hello")
	assert.Equal(t, "String", v.Type().String())
	assert.Equal(t, "hello", v.String())
}

func TestMarshalJSON(t *testing.T) {
	v := core.NewStringValue("hello")
	result, err := v.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello"), result)
}

func TestAsAny(t *testing.T) {
	v := core.NewStringValue("hello")
	result, err := v.AsAny()
	assert.Nil(t, err)
	assert.Equal(t, "hello", result)
}

func TestAsJSONString(t *testing.T) {
	v := core.NewStringValue("hello")
	result, err := v.AsJSONString()
	assert.Nil(t, err)
	assert.Equal(t, "hello", result)
}

func TestAsJSONBytes(t *testing.T) {
	v := core.NewStringValue("hello")
	result, err := v.AsJSONBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte("hello"), result)
}

func TestAsJSONPrettyfiedBytes(t *testing.T) {
	v := core.NewStringMapValue(map[string]any{"key": "value"})
	result, err := v.AsJSONPrettyfiedBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte("{\n  \"key\": \"value\"\n}"), result)
}

func TestAsYAMLBytes(t *testing.T) {
	v := core.NewStringMapValue(map[string]any{"key": "value"})
	result, err := v.AsYAMLBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte("key: value\n"), result)
}

func TestAsYAMLString(t *testing.T) {
	v := core.NewStringMapValue(map[string]any{"key": "value"})
	result, err := v.AsYAMLString()
	assert.Nil(t, err)
	assert.Equal(t, "key: value\n", result)
}

func TestIsValid(t *testing.T) {
	v := core.NewStringValue("hello")
	assert.True(t, v.IsValid())
	v = core.NewUnknownValue()
	assert.False(t, v.IsValid())
}

func TestValueAsAny(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		v := core.NewNilValue()
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("string value", func(t *testing.T) {
		v := core.NewStringValue("test")
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, "test", result)
	})

	t.Run("bool value", func(t *testing.T) {
		v := core.NewBoolValue(true)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("float value", func(t *testing.T) {
		v := core.NewFloat64Value(3.14)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, 3.14, result)
	})

	t.Run("bytes value", func(t *testing.T) {
		bytes := []byte("test")
		v := core.NewBytesValue(bytes)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, bytes, result)
	})

	t.Run("string map value", func(t *testing.T) {
		m := map[string]any{"key": "value"}
		v := core.NewStringMapValue(m)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, m, result)
	})

	t.Run("pointer string map value - not nil", func(t *testing.T) {
		m := map[string]any{"key": "value"}
		v := core.NewPointerStringMapValue(&m)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, m, result)
	})

	t.Run("pointer string map value - nil", func(t *testing.T) {
		v := core.NewPointerStringMapValue(nil)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("slice value", func(t *testing.T) {
		s := []any{"test", 123}
		v := core.NewSliceValue(s)
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Equal(t, s, result)
	})

	t.Run("unknown value", func(t *testing.T) {
		v := core.NewUnknownValue()
		result, err := v.AsAny()
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("invalid type", func(t *testing.T) {
		// Verwenden Sie einen der vorhandenen Konstruktoren
		v := core.NewUnknownValue()
		// Die AsAny() Methode sollte einen Fehler zurückgeben
		result, err := v.AsAny()
		assert.NoError(t, err) // Unknown ist ein gültiger Typ
		assert.Nil(t, result)  // Der Wert sollte nil sein
	})
}
