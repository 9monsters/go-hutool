package util

import "regexp"

func IsMatch(reg *regexp.Regexp, content string) bool {
	if len(content) == 0 {
		// 提供null的字符串为不匹配
		return false
	}
	return reg.MatchString(content)
}
