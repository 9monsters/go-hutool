package signer

import (
	"github.com/nine-monsters/go-hutool/hutool-core/text"
)

const IdNone string = "none"

var NONE *NoneSigner

func init() {
	NONE = &NoneSigner{}
}

type NoneSigner struct{}

func (NoneSigner) Sign(headerBase64 string, payloadBase64 string) string {
	return text.EMPTY
}

func (NoneSigner) Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool {
	return text.IsEmpty(signBase64)
}

func (NoneSigner) GetAlgorithm() string {
	return IdNone
}
