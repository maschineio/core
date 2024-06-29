package token_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestToken_AppendBoolToken(t *testing.T) {
	tmap := token.Map{}
	boolToken := token.NewBoolToken(token.End, true)
	err := tmap.Append(boolToken)
	assert.Nil(t, err)
	assert.True(t, tmap.Has(token.End))
	assert.Equal(t, token.Bool, boolToken.DataType())
	assert.Equal(t, token.End, boolToken.TokenType())
	assert.True(t, boolToken.IsTokenType(token.End))
	assert.True(t, boolToken.BoolVal())
	assert.Equal(t, "true", boolToken.String())
	// get token
	tok := tmap.GetToken(token.End)
	assert.NotNil(t, tok)

	tokNotFound := tmap.GetToken(token.StartAt)
	assert.Nil(t, tokNotFound)
}

func TestToken_StringToken(t *testing.T) {
	tok := token.NewStringToken(token.Comment, "comment")
	assert.Equal(t, "comment", tok.String())
	assert.Equal(t, "comment", tok.StringVal())
	assert.Equal(t, token.Comment, tok.TokenType())
	assert.Equal(t, token.String, tok.DataType())
}

func TestToken_Uint64Token(t *testing.T) {
	tok := token.NewUInt64Token(token.TimeoutSeconds, uint64(100))
	assert.Equal(t, "100", tok.String())
	assert.Equal(t, uint64(100), tok.UInt64Val())
}

func TestToken_TypeToken(t *testing.T) {
	tok := token.NewTypeToken(token.TaskType)
	assert.Equal(t, "TaskType", tok.String())
	assert.Equal(t, token.TaskType, tok.StateType())
}

func TestToken_SliceToken(t *testing.T) {
	tok := token.NewSliceToken(token.Parameters, []any{})
	assert.Equal(t, "[]", tok.String())
	assert.Equal(t, []any{}, tok.SliceVal())
}

func TestToken_StringMapToken(t *testing.T) {
	tok := token.NewStringMapToken(token.Parameters, map[string]any{})
	assert.Equal(t, "map[]", tok.String())
	assert.Equal(t, map[string]any{}, tok.StringMapVal())
}

func TestToken_JSONPathToken(t *testing.T) {
	tok, err := token.NewJSONPathToken(token.HeartbeatSecondsPath, "$.heartbeat")
	assert.Nil(t, err)
	assert.Equal(t, "$.heartbeat", tok.String())
}

func TestToken_JSONPathToken_Failing(t *testing.T) {
	_, err := token.NewJSONPathToken(token.HeartbeatSecondsPath, "$heartbeat")
	assert.NotNil(t, err)
	assert.Equal(t, "parsing error: $heartbeat\t:1:2 - 1:11 unexpected Ident while scanning operator", err.Error())
}

func TestToken_AnyToken(t *testing.T) {
	tok := token.NewAnyToken(token.Result, nil)
	assert.Nil(t, tok.AnyVal())
	assert.Equal(t, "<nil>", tok.String())
}

func TestToken_TimestampToken(t *testing.T) {
	myTime := time.Date(2009, time.November, 10, 23, 45, 0, 0, time.UTC)
	tok := token.NewTimestampToken(token.Timestamp, myTime)
	assert.Equal(t, myTime, tok.TimestampVal())
	assert.Equal(t, "TimestampType", tok.DataType().String())
	assert.Equal(t, "Timestamp", tok.TokenType().String())
}

func TestToken_IllegalToken(t *testing.T) {
	tok := token.NewIllegalToken()
	assert.Equal(t, nil, tok.AnyVal())
	assert.Equal(t, "Any", tok.DataType().String())
	assert.Equal(t, "ILLEGAL", tok.TokenType().String())
}

func TestToken_Float64(t *testing.T) {
	tok := token.NewFloat64Token(token.BackoffRate, 2.1)
	assert.Equal(t, "Float64", tok.DataType().String())
	assert.Equal(t, 2.1, tok.Float64Val())
	assert.Equal(t, "2.1", tok.String())
	assert.Equal(t, "BackoffRate", tok.TokenType().String())
}

func TestToken_StringSliceToken(t *testing.T) {
	tok := token.NewStringSliceToken(token.ErrorEquals, []string{"alpha", "beta", "gamma"})
	assert.Equal(t, "StringSlice", tok.DataType().String())
	assert.Equal(t, "ErrorEquals", tok.TokenType().String())
	assert.Equal(t, "[alpha beta gamma]", tok.String())
	assert.Equal(t, []string{"alpha", "beta", "gamma"}, tok.StringSliceVal())
}

func TestToken_SliceTokenMap(t *testing.T) {
	tok := token.NewSliceTokenMap(token.Retry, []token.Map{})
	assert.Equal(t, "Retry", tok.TokenType().String())
	assert.Equal(t, "SliceTokenMap", tok.DataType().String())
	assert.Equal(t, "[]", tok.String())
	assert.Equal(t, []token.Map{}, tok.SliceTokenMapVal())
}

func TestToken_IntToken(t *testing.T) {
	tok := token.NewIntToken(token.MaxAttempts, 1)
	assert.Equal(t, "MaxAttempts", tok.TokenType().String())
	assert.Equal(t, "Int", tok.DataType().String())
	assert.Equal(t, 1, tok.IntVal())
	assert.Equal(t, "1", tok.String())
}

func TestToken_TokenMapToken(t *testing.T) {
	tok := token.NewTokenMapToken(token.Parameters, token.Map{})
	assert.Equal(t, "Parameters", tok.TokenType().String())
	assert.Equal(t, "TokenMap", tok.DataType().String())
	assert.Equal(t, token.Map{}, tok.MapVal())
	assert.Equal(t, "map[]", tok.String())
}

func TestToken_ContextPathToken_InvalidJSONPath(t *testing.T) {
	_, err := token.NewContextPathToken(token.ContextObject, "$.test")
	assert.NotNil(t, err)
	assert.Equal(t, "expected context object json path", err.Error())
}

func TestToken_ContextPathToken_InvalidContextPath(t *testing.T) {
	_, err := token.NewContextPathToken(token.ContextObject, "$$.[]test")
	assert.NotNil(t, err)
	assert.Equal(t, "parsing error: $.[]test\t:1:3 - 1:4 unexpected \"[\" while scanning JSON select expected Ident, \".\" or \"*\"", err.Error())
}

func TestToken_ContextPathToken_ValidContextPath(t *testing.T) {
	tok, err := token.NewContextPathToken(token.ContextObject, "$$.test")
	assert.Nil(t, err)
	assert.Equal(t, "ContextPath", tok.DataType().String())
	assert.Equal(t, "$.test", tok.String())
	assert.Equal(t, "ContextObject", tok.TokenType().String())
}
