package text

import "github.com/nine-monsters/go-hutool/hutool-core/text/finder"

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
