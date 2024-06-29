package core

import (
	"regexp"
	"strings"
)

// StringMatchesRegex matches a pattern on a given string.
// The value MUST be a String which MAY contain one or more "*" characters.
// The expression yields true if the data value selected by the Variable Path matches the value, where "*" in the value matches zero or more characters.
// Thus, foo*.log matches foo23.log, *.log matches zebra.log, and foo*.* matches foobar.zebra.
// No characters other than "*" have any special meaning during matching.
func StringMatchesRegex(pattern, value string) (bool, error) {
	return regexp.MatchString(processWildcard(pattern), value)
}

func processWildcard(pattern string) string {
	var result strings.Builder
	for i, literal := range strings.Split(pattern, "*") {
		if i > 0 {
			result.WriteString(".*")
		}
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return result.String()
}
