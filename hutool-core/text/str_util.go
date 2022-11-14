package text

import (
	"strings"
	"unicode/utf8"

	validator "github.com/nine-monsters/go-hutool/hutool-core/validator"
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

// Creates an slice of slice values not included in the other given slice.
func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return result
}

func CamelCaseToUnderscore(str string) string {
	return validator.CamelCaseToUnderscore(str)
}

func UnderscoreToCamelCase(str string) string {
	return validator.UnderscoreToCamelCase(str)
}

func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return FindString(array, str) > -1
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

// Contains checks if the string contains the substring.
func Contains(str, substring string) bool {
	return strings.Contains(str, substring)
}
