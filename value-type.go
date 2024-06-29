package core

import (
	"bytes"
	"encoding/json"
	"fmt"

	"sigs.k8s.io/yaml"
)

type Value struct {
	dataType         Type
	boolValue        bool
	stringMapValue   map[string]any
	ptStringMapValue *map[string]any
	bytesValue       []byte
	stringValue      string
	// stringSliceValue []string
	floatValue float64
	anyValue   any
	sliceValue []any
}

func (v *Value) Type() Type {
	return v.dataType
}

func (v *Value) String() string {
	switch v.Type() {
	case Unknown:
		return "unknown"
	case Nil:
		return fmt.Sprintf("%v", nil)
	case String:
		return v.stringValue
	case Bool:
		return fmt.Sprintf("%v", v.boolValue)
	case Bytes: // hopefully this is already a json
		return fmt.Sprintf("%v", string(v.bytesValue))
	case Float:
		return fmt.Sprintf("%v", v.floatValue)
	case StringMap:
		return fmt.Sprintf("%+v", v.stringMapValue)
	case PointerStringMap:
		if v.ptStringMapValue != nil {
			return fmt.Sprintf("%+v", *v.ptStringMapValue)
		}
		return "unknown"
	case Slice:
		return fmt.Sprintf("%+v", v.sliceValue)

	default:
		return "core.Value: cant marshal unknown type"
	}
}

func (v *Value) BoolValue() bool {
	return v.boolValue
}

func (v *Value) MarshalJSON() (r []byte, err error) {
	switch v.Type() {
	case Nil:
		return json.Marshal(nil)
	case String:
		return []byte(v.stringValue), nil
	case Bool:
		return json.Marshal(v.boolValue)
	case Float:
		return json.Marshal(v.floatValue)
	case Bytes: // hopefully this is already a json
		return v.bytesValue, nil
	case StringMap:
		return json.Marshal(v.stringMapValue)
	case PointerStringMap:
		if v.ptStringMapValue != nil {
			return json.Marshal(*v.ptStringMapValue)
		}
		return json.Marshal(nil)
	case Slice:
		return json.Marshal(v.sliceValue)
	// case StringSlice:
	// 	fmt.Printf("#### stringslice value: %+v", v.stringSliceValue)
	// 	return json.Marshal(v.stringSliceValue)
	case Unknown:
		return json.Marshal(nil)
	default:
		return r, fmt.Errorf("core.Value: cant marshal unknown type: %v", v.Type())
	}
}

func (v *Value) AsAny() (result any, err error) {
	var b []byte
	b, err = v.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return
	}
	return
}

func (v *Value) AsJSONString() (s string, err error) {
	var b []byte
	b, err = v.AsJSONBytes()
	if err != nil {
		return
	}
	return string(b), nil
}

func (v *Value) AsJSONBytes() (s []byte, err error) {
	var b []byte
	b, err = v.MarshalJSON()
	if err != nil {
		return
	}
	return b, nil
}

func (v *Value) AsJSONPrettyfiedBytes() (s []byte, err error) {
	var prettyJSON bytes.Buffer
	src, err := v.AsJSONBytes()
	if err != nil {
		return
	}
	err = json.Indent(&prettyJSON, src, "", "  ")
	if err != nil {
		return
	}
	return prettyJSON.Bytes(), nil
}

func (v *Value) AsYAMLBytes() (s []byte, err error) {
	var b []byte
	b, err = v.AsJSONBytes()
	if err != nil {
		return
	}

	yml, err := yaml.JSONToYAML(b)
	if err != nil {
		return
	}
	return yml, nil
}

func (v *Value) AsYAMLString() (s string, err error) {
	var b []byte
	b, err = v.AsYAMLBytes()
	if err != nil {
		return
	}

	return string(b), nil
}

func (v *Value) IsValid() bool {
	return v.dataType != Unknown
}

func NewBoolValue(value bool) *Value {
	return &Value{boolValue: value, dataType: Bool}
}

func NewStringMapValue(m map[string]any) *Value {
	return &Value{stringMapValue: m, dataType: StringMap}
}

func NewPointerStringMapValue(m *map[string]any) *Value {
	return &Value{ptStringMapValue: m, dataType: PointerStringMap}
}

func NewBytesValue(b []byte) *Value {
	return &Value{bytesValue: b, dataType: Bytes}
}

func NewFloat64Value(f float64) *Value {
	return &Value{floatValue: f, dataType: Float}
}

func NewFloat32Value(f float32) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewIntValue(f int) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewInt8Value(f int8) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewInt16Value(f int16) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewInt32Value(f int32) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewInt64Value(f int64) *Value {
	return &Value{floatValue: float64(f), dataType: Float}
}

func NewNilValue() *Value {
	return &Value{anyValue: nil, dataType: Nil}
}

func NewSliceValue(s []any) *Value {
	return &Value{sliceValue: s, dataType: Slice}
}

// func NewStringSliceValue(s []string) Value {
// 	return Value{stringSliceValue: s, dataType: Slice}
// }

func NewUnknownValue() *Value {
	return &Value{dataType: Unknown}
}

func NewStringValue(s string) *Value {
	return &Value{stringValue: s, dataType: String}
}

func GetTypedValue(val any) (result *Value, err error) {
	switch t := val.(type) {
	case []byte:
		return NewBytesValue(t), nil
	case bool:
		return NewBoolValue(t), nil
	case map[string]any:
		return NewStringMapValue(t), nil
	case *map[string]any:
		return NewPointerStringMapValue(t), nil
	// case []string:
	// 	return NewStringSliceValue(t), nil
	case []any:
		return NewSliceValue(t), nil
	case int:
		return NewIntValue(t), nil
	case int8:
		return NewInt8Value(t), nil
	case int16:
		return NewInt16Value(t), nil
	case int32:
		return NewInt32Value(t), nil
	case int64:
		return NewInt64Value(t), nil
	case float32:
		return NewFloat32Value(t), nil
	case float64:
		return NewFloat64Value(t), nil
	case nil:
		return NewNilValue(), nil
	case string:
		return NewStringValue(t), nil
	case Value:
		return &t, nil
	case *Value:
		return t, nil
	default:
		return result, fmt.Errorf("GetTypedValue: unknown type %T", val)
	}
}
