package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
)

// Token data type
type Token struct {
	tokenType        TokenType
	dataType         DataType
	stateType        StateType
	strVal           string
	boolVal          bool
	intVal           int
	uint64Val        uint64
	float64Val       float64
	sliceVal         []any
	sliceTokenMapVal []Map
	mapVal           Map
	stringSliceVal   []string
	anyVal           any
	stringMapVal     map[string]any
	timestampVal     time.Time
	intrinsicFn      IntrinsicFn
}

// DataType returns the token data type
func (t Token) DataType() DataType {
	return t.dataType
}

// TokenType return the tokens type
func (t Token) TokenType() TokenType {
	return t.tokenType
}

// IsTokenType returns true if the token is from token type
func (t Token) IsTokenType(tok TokenType) bool {
	return t.tokenType == tok
}

// BoolVal returns the boolean value of the token
func (t Token) BoolVal() bool {
	return t.boolVal
}

// UInt64Val return the uint64 value of the token
func (t Token) UInt64Val() uint64 {
	return t.uint64Val
}

// Float64Val return the float64 value of the token
func (t Token) Float64Val() float64 {
	return t.float64Val
}

// StateType returns the state type of the token
func (t Token) StateType() StateType {
	return t.stateType
}

// StringVal returns the string value of the token
func (t Token) StringVal() string {
	return t.strVal
}

// SliceVal returns the slice value of the token
func (t Token) SliceVal() []any {
	return t.sliceVal
}

func (t Token) SliceTokenMapVal() []Map {
	return t.sliceTokenMapVal
}

// StringMapVal returns the map[string]any value of the token
func (t Token) StringMapVal() map[string]any {
	return t.stringMapVal
}

// AnyVal returns any typed value
func (t Token) AnyVal() any {
	return t.anyVal
}

// TimestampVal returns the timestamp value
func (t Token) TimestampVal() time.Time {
	return t.timestampVal
}

// StringSliceVal returns a slice of strings
func (t Token) StringSliceVal() []string {
	return t.stringSliceVal
}

// MapVal returns a Map
func (t Token) MapVal() Map {
	return t.mapVal
}

// IntVal return a int value
func (t Token) IntVal() int {
	return t.intVal
}

// FunctionVal return a IntrinsicFn value
func (t Token) FunctionVal() IntrinsicFn {
	return t.intrinsicFn
}

// IsComparator return true the token is a comparator token.
func (t Token) IsComparator() bool {
	return t.tokenType >= StringEquals && t.tokenType <= ArrayLenGreaterThanEqualsPath
}

// IsPathComparator return true if the variable value contains a path
func (t Token) IsPathComparator() bool {
	return (t.tokenType >= StringEqualsPath && t.tokenType <= TimestampGreaterThanEqualsPath) || (t.tokenType >= StringLenEqualsPath && t.tokenType <= ArrayLenGreaterThanEqualsPath)
}

func (t Token) IsStringComparator() bool {
	return t.tokenType >= StringEquals && t.tokenType <= StringMatches
}

func (t Token) IsStringJsonPathComparator() bool {
	return t.tokenType >= StringEqualsPath && t.tokenType <= StringGreaterThanEqualsPath
}

// String returns the string representation of the token
func (t Token) String() string {
	switch t.dataType {
	case String:
		return t.strVal
	case Any:
		return fmt.Sprintf("%+v", t.anyVal)
	case Bool:
		return fmt.Sprintf("%v", t.boolVal)
	case ContextPath:
		return t.strVal
	case Float64:
		return fmt.Sprintf("%v", t.float64Val)
	case Int:
		return fmt.Sprintf("%v", t.intVal)
	case JSONPath:
		return t.strVal
	case Slice:
		return fmt.Sprintf("%+v", t.sliceVal)
	case SliceTokenMap:
		return fmt.Sprintf("%+v", t.sliceTokenMapVal)
	case State:
		return fmt.Sprintf("%v", t.stateType)
	case StringMap:
		return fmt.Sprintf("%+v", t.stringMapVal)
	case StringSlice:
		return fmt.Sprintf("%v", t.stringSliceVal)
	case TimestampType:
		return fmt.Sprintf("%v", t.timestampVal)
	case TokenMap:
		return fmt.Sprintf("%+v", t.mapVal)
	case UInt64:
		return fmt.Sprintf("%v", t.uint64Val)
	case Function:
		return fmt.Sprintf("%+v", t.intrinsicFn)
	default:
		return t.strVal
	}
}

// NewStringToken returns a new string token with token type
func NewStringToken(tok TokenType, val string) Token {
	return Token{tokenType: tok, dataType: String, strVal: val}
}

// NewBoolToken returns a new boolean value token with token type
func NewBoolToken(tok TokenType, val bool) Token {
	return Token{tokenType: tok, dataType: Bool, boolVal: val}
}

// NewUInt64Token returns a new uint64 value token of token type
func NewUInt64Token(tok TokenType, val uint64) Token {
	return Token{tokenType: tok, dataType: UInt64, uint64Val: val}
}

// NewJSONPathToken returns a new JSONPath value token of token type.
// If the JSONPath is not valid it returns an error
func NewJSONPathToken(tok TokenType, val string) (Token, error) {
	_, err := jsonpath.New(val)
	if err != nil {
		return Token{tokenType: tok, dataType: JSONPath, strVal: val}, err
	}
	return Token{tokenType: tok, dataType: JSONPath, strVal: val}, nil
}

func NewContextPathToken(tok TokenType, val string) (Token, error) {
	if !strings.HasPrefix(val, "$$") {
		return Token{tokenType: tok, dataType: ContextPath, strVal: val}, fmt.Errorf("expected context object json path")
	}
	// we must remove the first char '$' from the incoming string; since its a $ we do not decode the rune size
	jsonPath := val[1:]

	// check jsonPath is a valid path
	_, err := jsonpath.New(jsonPath)
	if err != nil {
		return Token{tokenType: tok, dataType: ContextPath, strVal: val}, err
	}
	return Token{tokenType: tok, dataType: ContextPath, strVal: jsonPath}, nil
}

// NewTypeToken returns a new token token with a value of StateType
func NewTypeToken(val StateType) Token {
	return Token{tokenType: Type, dataType: State, stateType: val}
}

// NewSliceToken returns a new token with TokenType and a slice of []any
func NewSliceToken(tok TokenType, val []any) Token {
	return Token{tokenType: tok, dataType: Slice, sliceVal: val}
}

func NewSliceTokenMap(tok TokenType, val []Map) Token {
	return Token{tokenType: tok, dataType: SliceTokenMap, sliceTokenMapVal: val}
}

// NewStringMapToken returns a token with TokenType and map[string]any value
func NewStringMapToken(tok TokenType, val map[string]any) Token {
	return Token{tokenType: tok, dataType: StringMap, stringMapVal: val}
}

// NewAnyToken returns a token with TokenType and any value
func NewAnyToken(tok TokenType, val any) Token {
	return Token{tokenType: tok, dataType: Any, anyVal: val}
}

// NewTimestampToken returns a token with TokenType and a time.Time value
func NewTimestampToken(tok TokenType, val time.Time) Token {
	return Token{tokenType: tok, dataType: TimestampType, timestampVal: val}
}

// NewIllegalToken returns a ILLEGAL token with dataType any and a nil value
func NewIllegalToken() Token {
	return Token{tokenType: ILLEGAL, dataType: Any, anyVal: nil}
}

// NewStringSliceToken returns a token with TokenType and values of type string slice
func NewStringSliceToken(tok TokenType, val []string) Token {
	return Token{tokenType: tok, dataType: StringSlice, stringSliceVal: val}
}

// NewFloat64Token returns a float64 token
func NewFloat64Token(tok TokenType, val float64) Token {
	return Token{tokenType: tok, dataType: Float64, float64Val: val}
}

// NewTokenMapToken returns a Map token
func NewTokenMapToken(tok TokenType, val Map) Token {
	return Token{tokenType: tok, dataType: TokenMap, mapVal: val}
}

func NewIntToken(tok TokenType, val int) Token {
	return Token{tokenType: tok, dataType: Int, intVal: val}
}

func NewIntrinsicFnToken(tok TokenType, val IntrinsicFn) Token {
	return Token{tokenType: tok, dataType: Function, intrinsicFn: val}
}
