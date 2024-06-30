package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestStateTypeString(t *testing.T) {
	tok := token.TaskType

	assert.Equal(t, "TaskType", tok.String())
}

func TestStateTypeError(t *testing.T) {
	testCases := []token.StateType{
		-1, // less than 0
	}
	result := testCases[0].String()
	assert.Equal(t, "StateType(-1)", result)
}
