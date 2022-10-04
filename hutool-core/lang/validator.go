package lang

import (
	"regexp"

	"github.com/nine-monsters/go-hutool/hutool-core/util"
)

func isHex(value string) bool {
	return isMatchRegex(HEX, value)
}

func isMatchRegex(reg *regexp.Regexp, value string) bool {
	return util.IsMatch(reg, value)
}
