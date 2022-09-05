package jwt

import (
	"errors"
	"fmt"

	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
)

type Validate struct {
	Jwt
}
type Validator interface {
}

func validatorFromToken(token string) *Validate {
	return &Validate{
		*jwtFromToken(token),
	}
}

func validatorFromJwt(jwt *Jwt) *Validate {
	return &Validate{
		*jwt,
	}
}

func CreateByJwt(jwt *Jwt) *Validate {
	return &Validate{
		*jwt,
	}
}

func (v Validate) validateAlgorithm() (*Validate, error) {
	return v.validateAlgorithmWithSigner(nil)
}

func (v Validate) validateAlgorithmWithSigner(signer signer.Signer) (*Validate, error) {
	err := v.validateAlgorithmWithJwtAndSigner(v.Jwt, v.Signer)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (v Validate) validateAlgorithmWithJwtAndSigner(jwt Jwt, sg signer.Signer) error {
	if sg == nil {
		sg = jwt.getSigner()
	}
	algorithmId := jwt.getAlgorithm()
	if algorithmId == "" {
		if _, ok := sg.(signer.NoneSigner); sg == nil || ok {
			return nil
		}
		return errors.New("no algorithm defined in header")
	}
	if sg == nil {
		return errors.New("no Signer for validate algorithm")
	}

	algorithmIdInSigner := sg.GetAlgorithm()
	if false == text.Equals(algorithmId, algorithmIdInSigner) {
		return errors.New(fmt.Sprintf("algorithm [%s] defined in header doesn't match to [%s]!", algorithmId, algorithmIdInSigner))
	}

	// 通过算法验证签名是否正确
	if false == jwt.verifyWithSigner(sg) {
		return errors.New("signature verification failed")
	}
	return nil
}
