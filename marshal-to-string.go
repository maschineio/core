package core

import (
	"encoding/json"
	"fmt"
)

func MarshalToString(data any) (any, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func UnmarshalJsonBytesToAny(data any) (result any, err error) {
	if b, ok := data.([]byte); ok {
		if err := json.Unmarshal(b, &result); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
	return nil, fmt.Errorf("expected json bytes: got %T", data)
}

func TypeToAny(data any) (result any, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return
}
