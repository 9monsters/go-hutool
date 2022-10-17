package lang

import "regexp"

var (
	GENERAL *regexp.Regexp
	NUMBERS *regexp.Regexp
	HEX     *regexp.Regexp
)

func init() {
	GENERAL = regexp.MustCompile(GeneralReg)
	NUMBERS = regexp.MustCompile(NumbersReg)
	HEX = regexp.MustCompile(HexReg)
}
