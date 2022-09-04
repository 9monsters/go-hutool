package signer

import (
	"crypto"
	"crypto/hmac"
	"errors"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"github.com/nine-monsters/go-hutool/hutool-core/text"
	jwt "github.com/nine-monsters/go-hutool/hutool-jwt"
)

var (
	SigningMethodHS256  *HMacJwtSign
	SigningMethodHS384  *HMacJwtSign
	SigningMethodHS512  *HMacJwtSign
	ErrSignatureInvalid = errors.New("signature is invalid")
)

func init() {
	// HS256
	SigningMethodHS256 = &HMacJwtSign{"HS256", crypto.SHA256}
	RegisterJwtSigner(SigningMethodHS256.GetAlgorithmId(), func() JwtSigner {
		return SigningMethodHS256
	})

	// HS384
	SigningMethodHS384 = &HMacJwtSign{"HS384", crypto.SHA384}
	RegisterJwtSigner(SigningMethodHS384.GetAlgorithmId(), func() JwtSigner {
		return SigningMethodHS384
	})

	// HS512
	SigningMethodHS512 = &HMacJwtSign{"HS512", crypto.SHA512}
	RegisterJwtSigner(SigningMethodHS512.GetAlgorithmId(), func() JwtSigner {
		return SigningMethodHS512
	})
}

type HMacJWTSigner interface {
	JwtSigner
}

type HMacJwtSign struct {
	Name string
	Hash crypto.Hash
}

func CreateHmacSigner(key []byte) *HMacJwtSign {
	if key == nil {
		panic("Signer key must be not null!")
	}
	return &HMacJwtSign{}
}

func (s *HMacJwtSign) Sign(signingString string, key interface{}) (string, error) {
	if keyBytes, ok := key.([]byte); ok {
		if !s.Hash.Available() {
			return text.EMPTY, jwt.ErrHashUnavailable
		}

		hasher := hmac.New(s.Hash.New, keyBytes)
		hasher.Write([]byte(signingString))

		return codec.Base64.EncodeUrlSafeStrWithoutPadding(hasher.Sum(nil)), nil
	}

	return text.EMPTY, jwt.ErrInvalidKeyType
}

func (s *HMacJwtSign) Verify(signingString, signature string, key interface{}) (bool, error) {
	keyBytes, ok := key.([]byte)
	if !ok {
		return false, jwt.ErrInvalidKeyType
	}

	sig, err := codec.Base64.DecodeStr(signature)
	if err != nil {
		return false, err
	}

	if !s.Hash.Available() {
		return false, jwt.ErrHashUnavailable
	}

	hasher := hmac.New(s.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return false, ErrSignatureInvalid
	}

	return true, nil
}

func (s *HMacJwtSign) GetAlgorithm() string {
	return text.EMPTY
}

func (s *HMacJwtSign) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
