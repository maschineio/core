package core

import (
	"time"

	"github.com/PaesslerAG/jsonpath"
	"go.uber.org/zap"
)

func GetValue[T any](val any) (result T, ok bool) {
	if result, ok = val.(T); ok {
		return
	}
	return result, false
}

func GetTimestampValueWithLogger(tokenName, op1, op2 string, input any, logger *zap.Logger) (val1, val2 *time.Time, ok bool) {
	var strVal1, strVal2 string
	var val1Ok, val2Ok bool
	if strVal1, strVal2, ok = GetValuesWithLogger[string](tokenName, op1, op2, input, logger); !ok {
		return nil, nil, ok
	}
	if val1, val1Ok = GetTimestamp(strVal1); !val1Ok {
		if logger != nil {
			logger.Debug(tokenName, zap.Any("variable error", zap.Any("unexpected timestamp", strVal1)))
		}
		return nil, nil, false
	}
	if val2, val2Ok = GetTimestamp(strVal2); !val2Ok {
		if logger != nil {
			logger.Debug(tokenName, zap.Any("comparator error", zap.Any("unexpected timestamp", strVal2)))
		}
		return nil, nil, false
	}

	return val1, val2, true
}

func GetValuesWithLogger[T any](tokenName, op1, op2 string, input any, logger *zap.Logger) (val1 T, val2 T, ok bool) {
	// we have to get the value from input for var
	variable, err := jsonpath.Get(op1, input)
	if err != nil {
		if logger != nil {
			logger.Debug(tokenName, zap.Any("variable error", err))
		}
		ok = false
		return
	}
	compare, err := jsonpath.Get(op2, input)
	if err != nil {
		if logger != nil {
			logger.Debug(tokenName, zap.Any("comparator error", err))
		}
		ok = false
		return
	}

	// both (variable and comparator) must be from the same type
	val1, ok = GetValue[T](variable)
	if !ok {
		if logger != nil {
			logger.Debug(tokenName + ": unexpected variable type")
		}
		return
	}
	val2, ok = GetValue[T](compare)
	if !ok {
		if logger != nil {
			logger.Debug(tokenName + ": unexpected comparator value type")
		}
		return
	}

	return
}
