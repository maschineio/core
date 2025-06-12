package token

import (
	"fmt"
	"reflect"
)

type Func func(p []any, input any) (any, error)

type IntrinsicFn interface {
	Execute(input any) (any, error) // executes the function
	Name() string                   // function name
	String() string                 // string representation of the function
	IsFunc() bool                   // reflect the function is a function
}

type IntrinsicParam interface {
	Params() []any
	String() string
	Append(param any)
}

type iParam struct {
	params []any
}

func (i iParam) Params() []any {
	return i.params
}

func (i iParam) String() string {
	var ret string
	if i.params != nil {
		for idx, p := range i.params {
			ret += fmt.Sprintf("%v", p)
			if len(i.params)-1 != idx {
				ret += ","
			}
		}
	}
	return ret
}

func (i *iParam) Append(param any) {
	i.params = append(i.params, param)
}

func NewIntrinsicParameters(param ...any) IntrinsicParam {
	return &iParam{params: param}
}

// iFunction holds
type iFunction struct {
	name     string
	params   IntrinsicParam
	function *Func
}

func (i iFunction) Execute(input any) (any, error) {
	if i.function == nil {
		return nil, fmt.Errorf("intrinsic function has no bounded function")
	}

	f := *i.function
	return f(i.params.Params(), input)
}

func (i iFunction) Name() string {
	return i.name
}

func (i iFunction) IsFunc() bool {
	if i.function == nil {
		return false
	}
	return reflect.TypeOf(*i.function).Kind() == reflect.Func
}

func (i iFunction) String() string {
	return fmt.Sprintf("%s(%v)=%+v", i.name, i.params, i.function)
}

func NewIntrinsicFn(name string, params IntrinsicParam, fn *Func) IntrinsicFn {
	return iFunction{
		name:     name,
		params:   params,
		function: fn,
	}
}
