package hmac

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"hash"

	"github.com/nine-monsters/go-hutool/hutool-crypto/digest/mac"
)

type Hmac struct {
	engine    hash.Hash
	algorithm string
}

type Hmacer interface {
	mac.Macer
}

func New(algorithm string, key string) (Hmac, error) {
	engine, err := mac.CreateEngine(algorithm, key)
	if err != nil {
		return Hmac{}, errors.New("create hmac error")
	}
	hmac := Hmac{
		engine:    engine,
		algorithm: algorithm,
	}
	return hmac, nil
}

func (h Hmac) Digest(data string) []byte {
	h.engine.Write([]byte(data))
	sum := h.engine.Sum([]byte(""))
	return sum
}

func (h Hmac) DigestBase64(data string, isUrlSafe bool) string {
	h.engine.Write([]byte(data))
	sum := h.engine.Sum([]byte(""))
	var urlSafeStr string
	if isUrlSafe {
		urlSafeStr = codec.Base64.EncodeUrlSafeStr(sum)
	} else {
		urlSafeStr = codec.Base64.EncodeStr(sum)
	}
	return urlSafeStr
}

func (h Hmac) DigestHex(data string) string {
	h.engine.Write([]byte(data))
	sum := h.engine.Sum([]byte(""))
	hexStr := hex.EncodeToString(sum)
	return hexStr
}

func (h Hmac) Verify(digest []byte, digestToCompare []byte) bool {

	return bytes.Equal(digest, digestToCompare)
}

func (h Hmac) GetAlgorithm() string {
	return h.algorithm
}
