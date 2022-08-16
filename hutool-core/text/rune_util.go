package text

import "unicode"

// IsBlankRunes return rune is all blank
func IsBlankRunes(runes []rune) bool {
	if len(runes) == 0 {
		return true
	}
	for _, c := range runes {
		if !IsBlankRune(c) {
			return false
		}
	}
	return true
}

// IsBlankRune return rune is blank
func IsBlankRune(c rune) bool {
	return unicode.IsSpace(c)
}
