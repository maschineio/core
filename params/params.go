package params

import (
	"encoding/json"
	"fmt"
	"reflect"

	"dario.cat/mergo"
	"github.com/Jeffail/gabs/v2"
	"maschine.io/core"
	"maschine.io/core/replace"
	"maschine.io/core/token"
)

var errTypeMsg = "'%v' parameter must be a '%v': detected %T"

type Parameter struct {
	p *map[string]any
}

// NewParameter creates and parameter from map[string]any
func NewParameter(p *map[string]any) *Parameter {
	return &Parameter{p: p}
}

// NewDefaultParameter creates an initialized empty Parameter type
func NewDefaultParameter() *Parameter {
	par := make(map[string]any, 0)
	return &Parameter{p: &par}
}

// Add adds a new parameter with key/value
func (s *Parameter) Add(key string, value any) {
	(*s.p)[key] = value
}

// GetParams returns the underlying map[string]any
func (s *Parameter) GetParams() *map[string]any {
	return s.p
}

func (s *Parameter) Get(key string) any {
	if s.p == nil {
		return nil
	}

	if val, found := (*s.p)[key]; found {
		return val
	}

	return nil
}

func (s *Parameter) Keys() []string {
	if s.p == nil {
		return []string{}
	}
	keys := make([]string, len(*s.p))
	i := 0
	for key := range *s.p {
		keys[i] = key
		i++
	}
	return keys
}

// String return the parameters as string
func (s *Parameter) String() string {
	if s.p != nil {
		return fmt.Sprintf("%+v", *s.p)
	}
	return "nil"
}

// Merge parameters with input
func (s *Parameter) Merge(input map[string]any) (result map[string]any, err error) {
	// we return the input, if no parameters are set
	if s.p == nil {
		return input, nil
	}

	if err = mergo.Merge(&input, *s.p, mergo.WithOverride); err != nil {
		return
	}

	return input, err
}

// MergeAsBytes merges the parameters with input
func (s *Parameter) MergeAsBytes(input map[string]any) (result []byte, err error) {
	mergedInput, err := s.Merge(input)
	if err != nil {
		return
	}
	return json.Marshal(mergedInput)
}

// GetStringSliceParam from []any slice
func GetStringSliceParam(params *Parameter, name string) ([]string, error) {
	result := make([]string, 0)

	anyValues, err := GetParam[[]any](params, name)
	if err != nil {
		return nil, err
	}

	for _, value := range anyValues {
		if strValue, ok := value.(string); ok {
			result = append(result, strValue)
		}
	}
	return result, nil
}

// GetStringSliceParamDefault from []any slice with default if parameter was not found
func GetStringSliceParamDefault(params *Parameter, name string, defaultValues []string) ([]string, error) {
	result := make([]string, 0)
	if params == nil {
		return defaultValues, nil
	}

	iface := params.Get(name)
	if iface == nil {
		return defaultValues, nil
	}

	switch t := iface.(type) {
	case []string:
		return t, nil
	case []any:
		for _, value := range t {
			if strValue, ok := value.(string); ok {
				result = append(result, strValue)
			}
		}
		return result, nil
	default:
		return result, fmt.Errorf("'%v' parameter must be a '[]string': detected %T", name, iface)
	}
}

// GetOptionalParam returns nil if parameter not exists, a pointer to T if exists
func GetOptionalParam[T any](params *Parameter, name string) (*T, error) {
	var result T

	if params == nil {
		return nil, nil
	}

	iface := params.Get(name)
	// parameter not exists => return nil
	if iface == nil {
		return nil, nil
	}

	result, ok := iface.(T)
	if !ok {
		t := reflect.TypeOf((*T)(nil)).String()
		return &result, fmt.Errorf(errTypeMsg, name, t, iface)
	}

	return &result, nil
}

// GetParam returns the param by name from the params map by name
func GetParam[T any](params *Parameter, name string) (T, error) {
	var result T

	if params == nil {
		return result, fmt.Errorf("no params exists for name: %s", name)
	}

	iface := params.Get(name)
	if iface == nil {
		return result, fmt.Errorf("'%v' parameter expected", name)
	}

	result, ok := iface.(T)
	if !ok {
		t := reflect.TypeOf((*T)(nil)).String()
		return result, fmt.Errorf(errTypeMsg, name, t, iface)
	}

	return result, nil
}

// GetParamDefault tries to get a parameter by name; if not found it returns a default value
func GetParamDefault[T any](params *Parameter, name string, defaultValue T) (T, error) {
	var result T
	if params == nil {
		return defaultValue, nil
	}

	iface := params.Get(name)
	if iface == nil {
		return defaultValue, nil
	}

	switch t := iface.(type) {
	case T:
		return t, nil
	default:
		d := reflect.TypeOf((*T)(nil)).String()
		return result, fmt.Errorf(errTypeMsg, name, d, iface)
	}
}

// ProcessParameters processes the state parameters
// if @params is nil, the functions returns nil
func ProcessParameters(params *Parameter, input any) (*Parameter, error) {
	result := NewDefaultParameter()
	if params == nil {
		return params, nil
	}

	jsonPath, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	err = processKeys(params, result, jsonPath, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func parseInput(input any) (*gabs.Container, error) {
	if input == nil {
		return nil, nil
	}

	switch i := input.(type) {
	case []byte:
		jsonPath, err := gabs.ParseJSON(i)
		if err != nil {
			return nil, err
		}
		return jsonPath, nil
	case any:
		return gabs.Wrap(input), nil
	default:
		return nil, fmt.Errorf("unexpected input type: got %T", input)
	}
}

func processKeys(params *Parameter, result *Parameter, jsonPath *gabs.Container, input any) error {
	for _, key := range params.Keys() {
		value := params.Get(key)
		if err := processValue(key, value, result, jsonPath, input); err != nil {
			return err
		}
	}
	return nil
}

func processValue(key string, value any, result *Parameter, jsonPath *gabs.Container, input any) error {
	switch vt := value.(type) {
	case token.ReplacementKV[string]:
		return processStringReplacementKV(vt, result, jsonPath)
	case token.ReplacementKV[core.Replaceable]:
		return processReplaceableReplacementKV(key, vt, result, jsonPath)
	case token.ReplacementKV[token.Token]:
		return processTokenReplacementKV(key, vt, result, input)
	case map[string]any:
		result.Add(key, replace.ReplaceMap(&vt, input, jsonPath))
	default:
		result.Add(key, vt)
	}
	return nil
}

func processStringReplacementKV(vt token.ReplacementKV[string], result *Parameter, jsonPath *gabs.Container) error {
	value := vt.Value.Value()
	switch vt.Value.Type() {
	case token.JSONPath:
		result.Add(vt.Key.Key(), jsonPath.Path(value[2:]).Data())
	case token.String:
		result.Add(vt.Key.Key(), value)
	default:
		return fmt.Errorf("check/convert vt.Value.Type() for type: %v", vt.Value.Type())
	}
	return nil
}

func processReplaceableReplacementKV(key string, vt token.ReplacementKV[core.Replaceable], result *Parameter, jsonPath *gabs.Container) error {
	p := vt.Value.Value().JSONPath()
	result.Add(key, jsonPath.Path(p[2:]).Data())
	return nil
}

func processTokenReplacementKV(key string, vt token.ReplacementKV[token.Token], result *Parameter, input any) error {
	v, err := replace.ProcessReplacementKVToken(vt.Value, input)
	if err != nil {
		return err
	}
	result.Add(key, v)
	return nil
}
