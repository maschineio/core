package replace

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"maschine.io/core"
	"maschine.io/core/token"
)

func GetValue(key string, value any, input any, jsonPath *gabs.Container) (string, any, error) {
	switch vt := value.(type) {
	case token.ReplacementKV[string]:
		// vt.Key.Key() contains the new key to be used
		// we know that the result value of the DataValue is a string
		value := vt.Value.Value()
		switch vt.Value.Type() {
		case token.JSONPath:
			return vt.Key.Key(), jsonPath.Path(value[2:]).Data(), nil
		case token.String:
			return vt.Key.Key(), value, nil
		default:
			return key, nil, fmt.Errorf("check/convert vt.Value.Type() for type: %v", vt.Value.Type())
		}

	case token.ReplacementKV[core.Replaceable]:
		p := vt.Value.Value().JSONPath()
		return key, jsonPath.Path(p[2:]).Data(), nil

	case token.ReplacementKV[token.Token]:
		v, err := ProcessReplacementKVToken(vt.Value, input)
		if err != nil {
			return key, nil, err
		}
		return key, v, err

	default:
		// everything that comes here is inserted 1:1 into the parameters
		return key, vt, nil
	}
}
