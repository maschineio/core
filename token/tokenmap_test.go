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
func TestTokenMapTimeoutSeconds(t *testing.T) {
	t.Run("no timeout token", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewStringToken(token.StartAt, "start"))
		assert.Nil(t, err)

		timeout := tokMap.TimeoutSeconds()
		assert.Nil(t, timeout)
	})

	t.Run("has timeout token", func(t *testing.T) {
		expected := uint64(30)
		tokMap := token.Map{}
		// TimeoutSeconds Token mit NewUint64Token erstellen
		tok := token.NewUInt64Token(token.TimeoutSeconds, expected)
		err := tokMap.Append(tok)
		assert.Nil(t, err)

		timeout := tokMap.TimeoutSeconds()
		assert.NotNil(t, timeout)
		assert.Equal(t, expected, *timeout)
	})

	t.Run("empty map", func(t *testing.T) {
		tokMap := token.Map{}
		timeout := tokMap.TimeoutSeconds()
		assert.Nil(t, timeout)
	})
}
func TestTokenMapComment(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		tokMap := token.Map{}
		comment := tokMap.Comment()
		assert.Nil(t, comment)
	})

	t.Run("no comment token", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewStringToken(token.StartAt, "start"))
		assert.Nil(t, err)

		comment := tokMap.Comment()
		assert.Nil(t, comment)
	})

	t.Run("has comment token", func(t *testing.T) {
		expected := "test comment"
		tokMap := token.Map{}
		tok := token.NewStringToken(token.Comment, expected)
		err := tokMap.Append(tok)
		assert.Nil(t, err)

		comment := tokMap.Comment()
		assert.NotNil(t, comment)
		assert.Equal(t, expected, *comment)
	})
}
func TestTokenMapVersion(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		tokMap := token.Map{}
		assert.Panics(t, func() {
			tokMap.Version()
		})
	})

	t.Run("no version token", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewStringToken(token.StartAt, "start"))
		assert.Nil(t, err)

		assert.Panics(t, func() {
			tokMap.Version()
		})
	})

	t.Run("has version token", func(t *testing.T) {
		expected := "1.0.0"
		tokMap := token.Map{}
		tok := token.NewStringToken(token.Version, expected)
		err := tokMap.Append(tok)
		assert.Nil(t, err)

		version := tokMap.Version()
		assert.Equal(t, expected, version)
	})
}

func TestTokenMapStartAt(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		tokMap := token.Map{}
		assert.Panics(t, func() {
			tokMap.StartAt()
		})
	})

	t.Run("no startat token", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewStringToken(token.Version, "1.0.0"))
		assert.Nil(t, err)

		assert.Panics(t, func() {
			tokMap.StartAt()
		})
	})

	t.Run("has startat token", func(t *testing.T) {
		expected := "FirstState"
		tokMap := token.Map{}
		tok := token.NewStringToken(token.StartAt, expected)
		err := tokMap.Append(tok)
		assert.Nil(t, err)

		startAt := tokMap.StartAt()
		assert.Equal(t, expected, startAt)
	})
}
func TestTokenMapHasComparator(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		tokMap := token.Map{}
		assert.False(t, tokMap.HasComparator())
	})

	t.Run("no comparator tokens", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewStringToken(token.Next, "next"))
		assert.Nil(t, err)
		err = tokMap.Append(token.NewStringToken(token.Version, "1.0"))
		assert.Nil(t, err)
		assert.False(t, tokMap.HasComparator())
	})

	t.Run("one comparator token", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
		assert.Nil(t, err)
		err = tokMap.Append(token.NewStringToken(token.Next, "next"))
		assert.Nil(t, err)
		assert.True(t, tokMap.HasComparator())
	})

	t.Run("multiple comparator tokens", func(t *testing.T) {
		tokMap := token.Map{}
		err := tokMap.Append(token.NewBoolToken(token.IsArray, true))
		assert.Nil(t, err)
		err = tokMap.Append(token.NewBoolToken(token.IsNull, false))
		assert.Nil(t, err)
		assert.True(t, tokMap.HasComparator())
	})
}
