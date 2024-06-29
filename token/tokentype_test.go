package token_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestTokenTypeString(t *testing.T) {
	tok := token.End
	assert.Equal(t, "End", tok.String())
}

func TestTokenTypeIsComparator(t *testing.T) {
	for _, tt := range []struct {
		tok      token.TokenType
		expected any
	}{
		// comparators
		{token.StringEquals, true},
		{token.StringLessThan, true},
		{token.StringGreaterThan, true},
		{token.StringLessThanEquals, true},
		{token.StringGreaterThanEquals, true},
		{token.StringMatches, true},
		{token.NumericEquals, true},
		{token.NumericLessThan, true},
		{token.NumericGreaterThan, true},
		{token.NumericLessThanEquals, true},
		{token.NumericGreaterThanEquals, true},
		{token.BooleanEquals, true},
		{token.TimestampEquals, true},
		{token.TimestampLessThan, true},
		{token.TimestampGreaterThan, true},
		{token.TimestampLessThanEquals, true},
		{token.TimestampGreaterThanEquals, true},
		{token.IsNull, true},
		{token.IsPresent, true},
		{token.IsNumeric, true},
		{token.IsString, true},
		{token.IsBoolean, true},
		{token.IsTimestamp, true},
		{token.StringEqualsPath, true},
		{token.StringLessThanPath, true},
		{token.StringGreaterThanPath, true},
		{token.StringLessThanEqualsPath, true},
		{token.StringGreaterThanEqualsPath, true},
		{token.NumericEqualsPath, true},
		{token.NumericLessThanPath, true},
		{token.NumericGreaterThanPath, true},
		{token.NumericLessThanEqualsPath, true},
		{token.NumericGreaterThanEqualsPath, true},
		{token.BooleanEqualsPath, true},
		{token.TimestampEqualsPath, true},
		{token.TimestampLessThanPath, true},
		{token.TimestampGreaterThanPath, true},
		{token.TimestampLessThanEqualsPath, true},
		{token.TimestampGreaterThanEqualsPath, true},
		{token.StringLenEquals, true},
		{token.StringLenLessThan, true},
		{token.StringLenGreaterThan, true},
		{token.StringLenLessThanEquals, true},
		{token.StringLenGreaterThanEquals, true},
		{token.IsStringEmpty, true},
		{token.ArrayLenEquals, true},
		{token.ArrayLenLessThan, true},
		{token.ArrayLenGreaterThan, true},
		{token.ArrayLenLessThanEquals, true},
		{token.ArrayLenGreaterThanEquals, true},
		{token.IsArrayEmpty, true},
		{token.IsArray, true},
		{token.StringLenEqualsPath, true},
		{token.StringLenLessThanPath, true},
		{token.StringLenGreaterThanPath, true},
		{token.StringLenLessThanEqualsPath, true},
		{token.StringLenGreaterThanEqualsPath, true},
		{token.ArrayLenEqualsPath, true},
		{token.ArrayLenLessThenPath, true},
		{token.ArrayLenGreaterThanPath, true},
		{token.ArrayLenLessThanEqualsPath, true},
		{token.ArrayLenGreaterThanEqualsPath, true},
		// none comparators
		{token.ILLEGAL, false},
		{token.Comment, false},
		{token.End, false},
		{token.False, false},
		{token.Next, false},
		{token.Null, false},
		{token.Resource, false},
		{token.StartAt, false},
		{token.TimeoutSeconds, false},
		{token.TimeoutSecondsPath, false},
		{token.True, false},
		{token.Version, false},
		{token.InputPath, false},
		{token.OutputPath, false},
		{token.ResultPath, false},
		{token.ResultSelector, false},
		{token.Type, false},
		{token.States, false},
		{token.Id, false},
		{token.Parameters, false},
		{token.HeartbeatSecondsPath, false},
		{token.Result, false},
		{token.Credentials, false},
		{token.Seconds, false},
		{token.SecondsPath, false},
		{token.Timestamp, false},
		{token.TimestampPath, false},
		{token.Error, false},
		{token.Cause, false},
		{token.ErrorEquals, false},
		{token.IntervalSeconds, false},
		{token.MaxAttempts, false},
		{token.BackoffRate, false},
		{token.Retry, false},
		{token.Catch, false},
		{token.Choices, false},
		{token.Default, false},
		{token.Variable, false},
		{token.SimpleBoolRule, false},
		{token.AndBoolRule, false},
		{token.OrBoolRule, false},
		{token.NotBoolRule, false},
	} {
		t.Run(fmt.Sprintf("%v", tt.tok.String()), func(t *testing.T) {
			mTok := token.NewTokenMapToken(tt.tok, token.Map{})
			assert.Equal(t, tt.expected, mTok.IsComparator())
		})
	}
}
