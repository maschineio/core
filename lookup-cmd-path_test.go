package core_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestLookupCmdPath(t *testing.T) {
	t.Run("ValidCommand", func(t *testing.T) {
		path, err := core.LookupCmdPath("go")
		assert.Nil(t, err)
		assert.True(t, len(path) > 0)
	})

	t.Run("InvalidCommand", func(t *testing.T) {
		path, err := core.LookupCmdPath("this-command-does-not-exist")
		assert.NotNil(t, err)
		assert.True(t, strings.Contains(err.Error(), "cant get path for command"))
		assert.Equal(t, "", path)
	})
}
