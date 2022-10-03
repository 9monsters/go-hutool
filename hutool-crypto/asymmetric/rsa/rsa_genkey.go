package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"github.com/nine-monsters/go-hutool/hutool-crypto/errors"
)

// AsymmetricRsaKey rsaKey
type AsymmetricRsaKey struct {
	PrivateKey string
	PublicKey  string
}

// GenRsaKeyHex gen RsaKeyHex
func GenRsaKeyHex(bits int) (AsymmetricRsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return AsymmetricRsaKey{}, errors.ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return AsymmetricRsaKey{}, err
	}

	return AsymmetricRsaKey{
		PrivateKey: hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

// GenerateRsaKeyBase64 gen RsaKeyBase64
func GenerateRsaKeyBase64(bits int) (AsymmetricRsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return AsymmetricRsaKey{}, errors.ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return AsymmetricRsaKey{}, err
	}
	return AsymmetricRsaKey{
		PrivateKey: codec.Base64.EncodeStr(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  codec.Base64.EncodeStr(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}
