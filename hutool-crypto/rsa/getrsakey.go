package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"

	crypto "github.com/nine-monsters/go-hutool/hutool-crypto"
)

type RsaKey struct {
	PrivateKey string
	PublicKey  string
}

func GenerateRsaKeyHex(bits int) (RsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return RsaKey{}, crypto.ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

func GenerateRsaKeyBase64(bits int) (RsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return RsaKey{}, crypto.ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}
