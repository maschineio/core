package replace

import (
	"fmt"

	"maschine.io/core/token"
)

func ProcessReplacementKVToken(vt token.DataValue[token.Token], input any) (value any, err error) {
	tok := vt.Value()
	switch tok.DataType() {
	case token.Function:
		value, err = tok.FunctionVal().Execute(input)
	default:
		return nil, fmt.Errorf("unknown type on replacement kv token: got %s", tok.DataType().String())
	}
	return
}
