package token_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestNewIntrinsicParameters(t *testing.T) {
	p := token.NewIntrinsicParameters()
	assert.NotNil(t, p)
	assert.Empty(t, p.Params())

	p = token.NewIntrinsicParameters("a", 123)
	assert.Len(t, p.Params(), 2)
	assert.Equal(t, "a", p.Params()[0])
	assert.Equal(t, 123, p.Params()[1])
}

func TestIParam(t *testing.T) {
	p := token.NewIntrinsicParameters("val1")
	assert.Equal(t, 1, len(p.Params()))
	p.Append("val2")
	assert.Equal(t, 2, len(p.Params()))
	assert.Equal(t, "val1,val2", p.String())
}

func TestNewIntrinsicFn(t *testing.T) {
	fn := token.NewIntrinsicFn("testFn", token.NewIntrinsicParameters("p1"), nil)
	assert.Equal(t, "testFn", fn.Name())
	assert.False(t, fn.IsFunc())
	assert.Contains(t, fn.String(), "testFn")
}

func TestIFunctionExecute(t *testing.T) {
	successFunc := token.Func(func(p []any, input any) (any, error) {
		return fmt.Sprintf("executed with %v and %v", p, input), nil
	})
	failFunc := token.Func(func(p []any, input any) (any, error) {
		return nil, errors.New("failed")
	})

	t.Run("NilFunction", func(t *testing.T) {
		f := token.NewIntrinsicFn("nilFn", token.NewIntrinsicParameters("p1"), nil)
		res, err := f.Execute(nil)
		assert.Nil(t, res)
		assert.ErrorContains(t, err, "has no bounded function")
	})

	t.Run("SuccessFunction", func(t *testing.T) {
		f := token.NewIntrinsicFn("successFn", token.NewIntrinsicParameters("p1", "p2"), &successFunc)
		res, err := f.Execute("inputVal")
		assert.NoError(t, err)
		assert.Equal(t, "executed with [p1 p2] and inputVal", res)
	})

	t.Run("FailFunction", func(t *testing.T) {
		f := token.NewIntrinsicFn("failFn", token.NewIntrinsicParameters("p1"), &failFunc)
		res, err := f.Execute(nil)
		assert.Nil(t, res)
		assert.ErrorContains(t, err, "failed")
	})
}

func TestIFunctionIsFunc(t *testing.T) {
	validFn := token.Func(func(p []any, input any) (any, error) { return nil, nil })
	f := token.NewIntrinsicFn("validFn", token.NewIntrinsicParameters(), &validFn)
	assert.True(t, f.IsFunc())

	var ptr *token.Func
	f2 := token.NewIntrinsicFn("nilFn", token.NewIntrinsicParameters(), ptr)
	assert.False(t, f2.IsFunc())
}

func TestIFunctionString(t *testing.T) {
	testFn := token.Func(func(p []any, input any) (any, error) { return nil, nil })
	f := token.NewIntrinsicFn("strFn", token.NewIntrinsicParameters("p1"), &testFn)
	s := f.String()
	assert.Contains(t, s, "strFn")
	assert.Contains(t, s, "p1")
	assert.True(t, reflect.TypeOf(f).Kind() == reflect.Struct)
}

func TestIParamEmptyString(t *testing.T) {
	// Leerer Parameter für volle Abdeckung von iParam.String()
	p := token.NewIntrinsicParameters()
	assert.Equal(t, 0, len(p.Params()))
	assert.Equal(t, "", p.String(), "leerer Parameter sollte einen leeren String zurückgeben")
}

func TestIParamMixedTypes(t *testing.T) {
	// Verschiedene Typen in den Parametern testen (String, Zahl, Struktur)
	type dummy struct{ Name string }
	p := token.NewIntrinsicParameters("strVal", 42, dummy{Name: "structVal"})
	assert.Equal(t, 3, len(p.Params()))
	s := p.String()
	// Mindestcheck, dass alle Werte im String auftauchen
	assert.Contains(t, s, "strVal")
	assert.Contains(t, s, "42")
	assert.Contains(t, s, "structVal")
}

func TestIFunctionStringWithoutParams(t *testing.T) {
	// iFunction mit leeren Parametern testen
	testFn := token.Func(func(p []any, input any) (any, error) { return nil, nil })
	f := token.NewIntrinsicFn("emptyParamFn", token.NewIntrinsicParameters(), &testFn)
	s := f.String()
	assert.Contains(t, s, "emptyParamFn")
	// erwarten, dass weder p1 noch p2 enthalten ist
	assert.NotContains(t, s, "p1")
	assert.NotContains(t, s, "p2")
}
