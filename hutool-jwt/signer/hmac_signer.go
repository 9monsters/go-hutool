package signer

import (
	"github.com/nine-monsters/go-hutool/hutool-crypto/digest/hmac"
)

type HmacSigner struct {
	hmac hmac.Hmac
}

func (h *HmacSigner) Sign(headerBase64 string, payloadBase64 string) string {
	data := headerBase64 + "." + payloadBase64
	digestBase64 := h.hmac.DigestBase64(data, true)
	return digestBase64
}

func (h *HmacSigner) Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool {
	sign := h.Sign(headerBase64, payloadBase64)
	bytes1 := []byte(sign)
	bytes2 := []byte(signBase64)
	return h.hmac.Verify(bytes1, bytes2)
}

func (h *HmacSigner) GetAlgorithm() string {
	return h.hmac.GetAlgorithm()
}
