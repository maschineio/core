package core

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

// MaskSensitiveData
func MaskSensitiveData(e zapcore.Entry) error {
	e.Message = strings.Replace(e.Message, "dbpass", "**masked**", -1)
	return nil
}
