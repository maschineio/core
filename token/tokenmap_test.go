package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestTokenMapGetComparatorTokenCountEqualsZero(t *testing.T) {
	tokMap := token.Map{}
	assert.Equal(t, 0, tokMap.GetComparatorTokenCount())
}

func TestTokenMapGetComparatorTokenCountEqualsOne(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)
	err = tokMap.Append(token.NewBoolToken(token.End, true))
	assert.Nil(t, err)
	assert.Equal(t, 1, tokMap.GetComparatorTokenCount())
}

func TestTokenMapAppendTokenAlreadyExists(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)
	err = tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.NotNil(t, err)
	assert.Equal(t, "'IsArray' already exists", err.Error())
}

func TestTokenMapGetComparatorNoComparatorFound(t *testing.T) {
	tokMap := token.Map{}
	tok, err := tokMap.GetComparator()
	assert.Nil(t, tok)
	assert.NotNil(t, "", err.Error())
}

func TestTokenMapGetComparatorToManyComparatorsFound(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)

	err = tokMap.Append(token.NewBoolToken(token.IsNull, false))
	assert.Nil(t, err)

	tok, err := tokMap.GetComparator()
	assert.Nil(t, tok)
	assert.Equal(t, "to many comparators found", err.Error())
}

func TestTokenMapGetComparatorValid(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)

	err = tokMap.Append(token.NewStringToken(token.Next, "hello world"))
	assert.Nil(t, err)

	tok, err := tokMap.GetComparator()
	assert.Nil(t, err)
	assert.NotNil(t, tok)
}

func TestTokenMapNewTokenMapNoTokens(t *testing.T) {
	tm, err := token.NewTokenMap()
	assert.Nil(t, err)
	assert.Equal(t, token.Map{}, tm)
}

func TestTokenMapNewTokenMapDuplicateEntry(t *testing.T) {
	t1 := token.NewBoolToken(token.End, true)
	_, err := token.NewTokenMap(t1, t1)
	assert.NotNil(t, err)
	assert.Equal(t, "'End' already exists", err.Error())
}

func TestTokenMapNewTokenMapValid(t *testing.T) {
	t1 := token.NewBoolToken(token.End, true)
	tm, err := token.NewTokenMap(t1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tm))
}
