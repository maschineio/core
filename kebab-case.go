// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package core

import "strings"

func KebabCase(s string) string {
	return delimiterCase(s, '-')
}

// delimiterCase converts a string into kebab-case depending on the delimiter passed
// as second argument.
func delimiterCase(s string, delimiter rune) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s)+3)

	processChar := func(buffer *[]rune, prev, curr, next rune, delimiter rune, toLower func(rune) rune) {
		if isDelimiter(curr) {
			if !isDelimiter(prev) {
				*buffer = append(*buffer, delimiter)
			}
		} else if isUpper(curr) {
			if isLower(prev) || (isUpper(prev) && isLower(next)) {
				*buffer = append(*buffer, delimiter)
			}
			*buffer = append(*buffer, toLower(curr))
		} else if curr != 0 {
			*buffer = append(*buffer, toLower(curr))
		}
	}

	var prev rune
	var curr rune
	for _, next := range s {
		processChar(&buffer, prev, curr, next, delimiter, toLower)
		prev = curr
		curr = next
	}

	if len(s) > 0 {
		processChar(&buffer, prev, curr, 0, delimiter, toLower)
	}

	return string(buffer)
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

// isDelimiter checks if a character is some kind of whitespace or '_' or '-'.
func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

// isSpace checks if a character is some kind of whitespace.
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isLower checks if a character is upper case. More precisely it evaluates if it is
// in the range of ASCII characters 'A' to 'Z'.
func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// isLower checks if a character is lower case. More precisely it evaluates if it is
// in the range of ASCII character 'a' to 'z'.
func isLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}
