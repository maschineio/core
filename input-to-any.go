package core

import (
	"encoding/json"
	"fmt"
)

// InputToBytes converts a input (mostly []byte) to Interface{}
func InputToBytes(input any) (result []byte, err error) {
	switch t := input.(type) {
	case []byte:
		return input.([]byte), nil
	case map[string]any:
		if result, err := json.Marshal(t); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	case Value:
		return t.AsJSONBytes()
	default:
		return nil, fmt.Errorf("core: unknown input type: got %T", input)
	}
}
