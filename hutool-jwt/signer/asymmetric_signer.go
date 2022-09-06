package signer

import "github.com/nine-monsters/go-hutool/hutool-crypto/digest/hmac"

type AsymmetricSigner struct {
	hmac hmac.Hmac
}
