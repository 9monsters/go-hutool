package signer

import "github.com/nine-monsters/go-hutool/hutool-core/text"

type AsymmetricJwtSigner interface {
	JwtSigner
}
type AsymmetricJwtSign struct {
	JwtSign
}

func CreateAsymmetricJwtSigner(key []byte) *AsymmetricJwtSign {
	if key == nil {
		panic("Signer key must be not null!")
	}
	return &AsymmetricJwtSign{}
}

func (s *AsymmetricJwtSign) Sign(headerBase64 string, payloadBase64 string) string {
	return text.EMPTY
}

func (s *AsymmetricJwtSign) Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool {
	return true
}

func (s *AsymmetricJwtSign) GetAlgorithm() string {
	return text.EMPTY
}

func (s *AsymmetricJwtSign) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}

func SetSigner(jwtSigner *AsymmetricJwtSign) {
	return SetSigner()
}
