package core

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

// MaskSensitiveData
// TODO: its not only password, it shoult be a regex, and also users are also sensitive
func MaskSensitiveData(e zapcore.Entry) error {
	e.Message = strings.Replace(e.Message, "dbpass", "**masked**", -1)
	return nil
}
