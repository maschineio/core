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
