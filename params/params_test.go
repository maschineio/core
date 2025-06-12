package params_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/params"
)

func TestGetParam(t *testing.T) {
	// Test case: Parameter exists and is of correct type
	t.Run("Parameter exists and is of correct type", func(t *testing.T) {
		p := map[string]any{"key": 42}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter does not exist
	t.Run("Parameter does not exist", func(t *testing.T) {
		p := map[string]any{}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter exists but is of incorrect type
	t.Run("Parameter exists but is of incorrect type", func(t *testing.T) {
		p := map[string]any{"key": "value"}
		param := params.NewParameter(&p)
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter map is nil
	t.Run("Parameter map is nil", func(t *testing.T) {
		var param *params.Parameter
		result, err := params.GetParam[int](param, "key")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestGetParamDefault(t *testing.T) {
	// Test case: Parameter exists and is of correct type
	t.Run("Parameter exists and is of correct type", func(t *testing.T) {
		p := map[string]any{"key": 42}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 0)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter does not exist
	t.Run("Parameter does not exist", func(t *testing.T) {
		p := map[string]any{}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 42)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	// Test case: Parameter exists but is of incorrect type
	t.Run("Parameter exists but is of incorrect type", func(t *testing.T) {
		p := map[string]any{"key": "value"}
		param := params.NewParameter(&p)
		result, err := params.GetParamDefault(param, "key", 42)

		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})

	// Test case: Parameter map is nil
	t.Run("Parameter map is nil", func(t *testing.T) {
		var param *params.Parameter
		result, err := params.GetParamDefault(param, "key", 42)

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})
}
func TestNewDefaultParameter(t *testing.T) {
	param := params.NewDefaultParameter()
	assert.NotNil(t, param)
	assert.NotNil(t, param.GetParams())
	assert.Empty(t, *param.GetParams())
}

func TestParameterAdd(t *testing.T) {
	param := params.NewDefaultParameter()
	param.Add("key1", "value1")
	param.Add("key2", 42)

	assert.Equal(t, "value1", param.Get("key1"))
	assert.Equal(t, 42, param.Get("key2"))
}

func TestParameterGet(t *testing.T) {
	// t.Run("nil parameter", func(t *testing.T) {
	// 	var param *params.Parameter
	// 	assert.Nil(t, param.Get("key"))
	// })

	t.Run("existing key", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key", "value")
		assert.Equal(t, "value", param.Get("key"))
	})

	t.Run("non-existing key", func(t *testing.T) {
		param := params.NewDefaultParameter()
		assert.Nil(t, param.Get("nonexistent"))
	})
}

// func TestParameterKeys(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		params   map[string]any
// 		expected []string
// 	}{
// 		{
// 			name:     "leere Parameter",
// 			params:   map[string]any{},
// 			expected: []string{},
// 		},
// 		{
// 			name: "einfache Parameter",
// 			params: map[string]any{
// 				"key1": "value1",
// 				"key2": "value2",
// 			},
// 			expected: []string{"key1", "key2"},
// 		},
// 		{
// 			name: "verschachtelte Parameter",
// 			params: map[string]any{
// 				"key1": map[string]any{
// 					"subkey1": "value1",
// 				},
// 				"key2": "value2",
// 			},
// 			expected: []string{"key1", "key2"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := params.NewParameter(&tt.params)
// 			got := p.String()

// 			// JSON Validierung
// 			var gotJSON, wantJSON map[string]any
// 			err := json.Unmarshal([]byte(got), &gotJSON)
// 			require.NoError(t, err, "invalid character 'm' looking for beginning of value")

// 			expectedJSON, err := json.Marshal(tt.expected)
// 			require.NoError(t, err)
// 			err = json.Unmarshal(expectedJSON, &wantJSON)
// 			require.NoError(t, err, "invalid character 'm' looking for beginning of value")

// 			assert.Equal(t, wantJSON, gotJSON)
// 		})
// 	}
// }

func TestParameterString(t *testing.T) {
	// t.Run("nil parameter", func(t *testing.T) {
	// 	var param *params.Parameter
	// 	assert.Equal(t, "nil", param.String())
	// })

	t.Run("empty parameter", func(t *testing.T) {
		param := params.NewDefaultParameter()
		assert.Equal(t, "map[]", param.String())
	})

	t.Run("single value", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key", "value")
		assert.Equal(t, "map[key:value]", param.String())
	})

	t.Run("multiple values", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key1", "value1")
		param.Add("key2", 42)
		param.Add("key3", true)
		assert.Contains(t, param.String(), "key1:value1")
		assert.Contains(t, param.String(), "key2:42")
		assert.Contains(t, param.String(), "key3:true")
	})

	t.Run("nested map", func(t *testing.T) {
		param := params.NewDefaultParameter()
		nested := map[string]any{"inner": "value"}
		param.Add("outer", nested)
		assert.Contains(t, param.String(), "outer:map[inner:value]")
	})

	t.Run("array value", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("array", []string{"a", "b", "c"})
		assert.Contains(t, param.String(), "array:[a b c]")
	})
}

func TestParameterMerge(t *testing.T) {
	tests := []struct {
		name    string
		base    map[string]any
		merge   map[string]any
		want    map[string]any
		wantErr bool
	}{
		// {
		// 	name:    "nil parameter",
		// 	base:    nil,
		// 	merge:   nil,
		// 	want:    make(map[string]any),
		// 	wantErr: false,
		// },
		{
			name:    "empty maps",
			base:    map[string]any{},
			merge:   map[string]any{},
			want:    map[string]any{},
			wantErr: false,
		},
		{
			name:    "simple maps",
			base:    map[string]any{"a": 1},
			merge:   map[string]any{"b": 2},
			want:    map[string]any{"a": 1, "b": 2},
			wantErr: false,
		},
		// {
		// 	name:    "overwrite existing values",
		// 	base:    map[string]any{"a": 1},
		// 	merge:   map[string]any{"a": 2},
		// 	want:    map[string]any{"a": 2},
		// 	wantErr: false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := params.NewParameter(&tt.base)
			merge := params.NewParameter(&tt.merge)

			// Direktes Mergen ohne Zwischenkopien
			result, err := base.Merge(*merge.GetParams())

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			// Expliziter Vergleich der Maps
			assert.Equal(t, tt.want, result, "merge result is not identical to expected result")
		})
	}
}

func TestParameterMergeAsBytes(t *testing.T) {
	param := params.NewDefaultParameter()
	param.Add("key1", "value1")

	input := map[string]any{"key2": "value2"}

	result, err := param.MergeAsBytes(input)
	assert.NoError(t, err)
	assert.Contains(t, string(result), "key1")
	assert.Contains(t, string(result), "value1")
	assert.Contains(t, string(result), "key2")
	assert.Contains(t, string(result), "value2")
}

func TestGetStringSliceParam(t *testing.T) {
	t.Run("valid string slice", func(t *testing.T) {
		p := map[string]any{"key": []any{"value1", "value2"}}
		param := params.NewParameter(&p)
		result, err := params.GetStringSliceParam(param, "key")
		assert.NoError(t, err)
		assert.Equal(t, []string{"value1", "value2"}, result)
	})

	t.Run("non-existent key", func(t *testing.T) {
		param := params.NewDefaultParameter()
		result, err := params.GetStringSliceParam(param, "nonexistent")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetOptionalParam(t *testing.T) {
	t.Run("nil parameter", func(t *testing.T) {
		var param *params.Parameter
		result, err := params.GetOptionalParam[string](param, "key")
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("existing value", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key", "value")
		result, err := params.GetOptionalParam[string](param, "key")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "value", *result)
	})

	t.Run("wrong type", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key", 42)
		result, err := params.GetOptionalParam[string](param, "key")
		assert.Error(t, err)
		assert.NotNil(t, result)
	})
}

func TestProcessParameters(t *testing.T) {
	t.Run("nil parameters", func(t *testing.T) {
		result, err := params.ProcessParameters(nil, nil)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("with JSON input", func(t *testing.T) {
		param := params.NewDefaultParameter()
		param.Add("key", "value")

		input := []byte(`{"data": "test"}`)
		result, err := params.ProcessParameters(param, input)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
func TestGetStringSliceParamDefault(t *testing.T) {
	defaultValues := []string{"default1", "default2"}

	tests := []struct {
		name         string
		params       map[string]any
		key          string
		defaultVals  []string
		expected     []string
		expectError  bool
		errorMessage string
	}{
		{
			name:        "nil parameter",
			params:      nil,
			key:         "test",
			defaultVals: defaultValues,
			expected:    defaultValues,
		},
		{
			name:        "key not found",
			params:      map[string]any{"otherkey": []string{"val1"}},
			key:         "test",
			defaultVals: defaultValues,
			expected:    defaultValues,
		},
		{
			name:        "valid []string parameter",
			params:      map[string]any{"test": []string{"val1", "val2"}},
			key:         "test",
			defaultVals: defaultValues,
			expected:    []string{"val1", "val2"},
		},
		{
			name:        "valid []any parameter with strings",
			params:      map[string]any{"test": []any{"val1", "val2"}},
			key:         "test",
			defaultVals: defaultValues,
			expected:    []string{"val1", "val2"},
		},
		{
			name:        "empty slice parameter",
			params:      map[string]any{"test": []string{}},
			key:         "test",
			defaultVals: defaultValues,
			expected:    []string{},
		},
		{
			name:         "wrong parameter type",
			params:       map[string]any{"test": 123},
			key:          "test",
			defaultVals:  defaultValues,
			expectError:  true,
			errorMessage: "'test' parameter must be a '[]string': detected int",
		},
		{
			name:        "[]any with mixed types - extracts only strings",
			params:      map[string]any{"test": []any{"str1", 123, "str2", true}},
			key:         "test",
			defaultVals: defaultValues,
			expected:    []string{"str1", "str2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var param *params.Parameter
			if tt.params != nil {
				param = params.NewParameter(&tt.params)
			}

			result, err := params.GetStringSliceParamDefault(param, tt.key, tt.defaultVals)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMessage != "" {
					assert.Equal(t, tt.errorMessage, err.Error())
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
