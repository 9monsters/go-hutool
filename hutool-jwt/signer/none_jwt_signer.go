package signer

import (
	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"github.com/nine-monsters/go-hutool/hutool-jwt"
)

var NoneJwtSignerConst *NoneJwtSign

const UnsafeAllowNoneSignatureType unsafeNoneMagicConstant = "none signing method allowed"

var NoneSignatureTypeDisallowedError error

type unsafeNoneMagicConstant string

type NoneJwtSign struct {
	JwtSign
}

func init() {
	NoneJwtSignerConst = &NoneJwtSign{}
	NoneSignatureTypeDisallowedError = jwt.NewValidationError("'none' signature type is not allowed", jwt.ValidationErrorSignatureInvalid)
	RegisterJwtSigner(NoneJwtSignerConst.GetAlgorithmId(), func() JwtSigner {
		return NoneJwtSignerConst
	})
}

func (s *NoneJwtSign) Sign(signingString string, key interface{}) (string, error) {
	if _, ok := key.(unsafeNoneMagicConstant); ok {
		return text.EMPTY, nil
	}
	return text.EMPTY, NoneSignatureTypeDisallowedError
}

func (s *NoneJwtSign) Verify(signingString, signature string, key interface{}) (res bool, err error) {
	if _, ok := key.(unsafeNoneMagicConstant); !ok {
		return false, NoneSignatureTypeDisallowedError
	}
	if signature != "" {
		return false, jwt.NewValidationError(
			"'none' signing method with non-empty signature",
			jwt.ValidationErrorSignatureInvalid,
		)
	}
	return true, nil
}

func (s *NoneJwtSign) GetAlgorithm() string {
	return "none"
}

func (s *NoneJwtSign) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
