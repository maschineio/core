package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestTokenMap_GetComparatorTokenCount_EqualsZero(t *testing.T) {
	tokMap := token.Map{}
	assert.Equal(t, 0, tokMap.GetComparatorTokenCount())
}

func TestTokenMap_GetComparatorTokenCount_EqualsOne(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)
	err = tokMap.Append(token.NewBoolToken(token.End, true))
	assert.Nil(t, err)
	assert.Equal(t, 1, tokMap.GetComparatorTokenCount())
}

func TestTokenMap_Append_TokenAlreadyExists(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)
	err = tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.NotNil(t, err)
	assert.Equal(t, "'IsArray' already exists", err.Error())
}

func TestTokenMap_GetComparator_NoComparatorFound(t *testing.T) {
	tokMap := token.Map{}
	tok, err := tokMap.GetComparator()
	assert.Nil(t, tok)
	assert.NotNil(t, "", err.Error())
}

func TestTokenMap_GetComparator_ToManyComparatorsFound(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)

	err = tokMap.Append(token.NewBoolToken(token.IsNull, false))
	assert.Nil(t, err)

	tok, err := tokMap.GetComparator()
	assert.Nil(t, tok)
	assert.Equal(t, "to many comparators found", err.Error())
}

func TestTokenMap_GetComparator_Valid(t *testing.T) {
	tokMap := token.Map{}
	err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
	assert.Nil(t, err)

	err = tokMap.Append(token.NewStringToken(token.Next, "hello world"))
	assert.Nil(t, err)

	tok, err := tokMap.GetComparator()
	assert.Nil(t, err)
	assert.NotNil(t, tok)
}

func TestTokenMap_NewTokenMap_NoTokens(t *testing.T) {
	tm, err := token.NewTokenMap()
	assert.Nil(t, err)
	assert.Equal(t, token.Map{}, tm)
}

func TestTokenMap_NewTokenMap_DuplicateEntry(t *testing.T) {
	t1 := token.NewBoolToken(token.End, true)
	_, err := token.NewTokenMap(t1, t1)
	assert.NotNil(t, err)
	assert.Equal(t, "'End' already exists", err.Error())
}

func TestTokenMap_NewTokenMap_Valid(t *testing.T) {
	t1 := token.NewBoolToken(token.End, true)
	tm, err := token.NewTokenMap(t1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tm))
}