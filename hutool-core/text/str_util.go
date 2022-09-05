package text

import (
	"strings"
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

// Equals return seq1 equal to seq2
func Equals(seq1 string, seq2 string) bool {
	return EqualsIgnoreCase(seq1, seq2, false)
}

func EqualsIgnoreCase(seq1 string, seq2 string, ignoreCase bool) bool {
	if seq1 == "" {
		return seq2 == ""
	}
	if seq2 == "" {
		return false
	}
	if ignoreCase {
		return seq1 == seq2
	} else {
		return strings.EqualFold(seq1, seq2)
	}
}
