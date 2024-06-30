package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestDataTypeString(t *testing.T) {
	tok := token.String
	assert.Equal(t, "String", tok.String())
}

func TestDataTypeError(t *testing.T) {
	testCases := []token.DataType{
		-1, // less than 0
	}
	result := testCases[0].String()
	assert.Equal(t, "DataType(-1)", result)
}
