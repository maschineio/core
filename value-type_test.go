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
