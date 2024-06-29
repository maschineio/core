package token

import (
	"fmt"
)

// Map represents the parsed state machine.
type Map map[TokenType]Token

// TimeoutSeconds returns the timeout in seconds if exist;
// nil otherwise
func (m Map) TimeoutSeconds() *uint64 {
	if m.Has(TimeoutSeconds) {
		timeout := m.GetToken(TimeoutSeconds).UInt64Val()
		return &timeout
	}
	return nil
}

// Comment get the comment from the Map if exists
func (m Map) Comment() *string {
	if m.Has(Comment) {
		comment := m.GetToken(Comment).StringVal()
		return &comment
	}
	return nil
}

// Version returns the Version of this state machine
func (m Map) Version() string {
	return m.GetToken(Version).StringVal()
}

// StartAt returns the start state id
func (m Map) StartAt() string {
	return m.GetToken(StartAt).StringVal()
}

// Append a token to the Map. If the token already exists in the
// Map a error will be raised.
func (m Map) Append(tok Token) error {
	if m.Has(tok.tokenType) {
		return fmt.Errorf("'%v' already exists", tok.tokenType.String())
	}
	m[tok.tokenType] = tok
	return nil
}

// Has checks the Map for the token type.
// It returns true if the token exists.
func (m Map) Has(tokType TokenType) bool {
	_, ok := m[tokType]
	return ok
}

// GetToken returns the token type if exists;
// if its not exists the method returns nil.
func (m Map) GetToken(tokType TokenType) *Token {
	if val, ok := m[tokType]; ok {
		return &val
	}
	return nil
}

// GetComparatorTokenCount calculates the number of comparators
func (m Map) GetComparatorTokenCount() int {
	count := 0
	for _, tok := range m {
		if tok.IsComparator() {
			count++
		}
	}
	return count
}

func (m Map) HasComparator() bool {
	return m.GetComparatorTokenCount() > 0
}

// GetComparator return the token of a comparator.
// If comparator count is zero or greater than one
// a error will be returned.
func (m Map) GetComparator() (tokResult *Token, err error) {
	if m.GetComparatorTokenCount() == 0 {
		return nil, fmt.Errorf("no comparator found")
	}

	if m.GetComparatorTokenCount() > 1 {
		return nil, fmt.Errorf("to many comparators found")
	}

	for _, tok := range m {
		if tok.IsComparator() {
			return &tok, nil
		}
	}

	return
}

// NewTokenMap creates a new Map and puts the tokens into the map
func NewTokenMap(tokens ...Token) (Map, error) {
	result := Map{}

	for _, tok := range tokens {
		if err := result.Append(tok); err != nil {
			return nil, err
		}
	}

	return result, nil
}
