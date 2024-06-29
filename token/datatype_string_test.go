package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestDataType_String(t *testing.T) {
	tok := token.String

	assert.Equal(t, "String", tok.String())
}
