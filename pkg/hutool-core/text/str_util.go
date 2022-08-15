package text

import (
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/text/finder"
)

const (
	INDEX_NOT_FOUND int8   = finder.INDEX_NOT_FOUND
	NULL            string = "NULL"
	EMPTY           string = ""
	SPACE           string = " "
)

// IsBlank return str is all blank
func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	return IsBlankRunes([]rune(str))
}

// IsEmpty return str is empty
func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

// Trim returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func Trim(str string) string {
	if len(str) == 0 {
		return EMPTY
	}
	return strings.TrimSpace(str)
}

// Trims returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func Trims(str []string) []string {
	if len(str) == 0 {
		return []string{}
	}
	// slice size = 0, cap = str length
	res := make([]string, 0, len(str))
	for _, s := range str {
		res = append(res, Trim(s))
	}
	return res
}
