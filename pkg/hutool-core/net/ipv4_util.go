package net

import (
	"bytes"
	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"strconv"
)

const (
	LOCAL_IP           string = "127.0.0.1"
	IP_SPLIT_MARK      string = text.DASHED
	IP_MASK_SPLIT_MARK string = text.SLASH
	IP_MASK_MAX        int    = 32
)

func LongToIPv4(longIP int64) string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(longIP>>24&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP>>16&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP>>8&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP&0xFF, 10))
	return buffer.String()
}
