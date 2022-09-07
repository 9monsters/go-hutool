package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
)

type Validate struct {
	Jwt
}
type Validator interface {
	validateAlgorithm() (*Validate, error)
	validateAlgorithmWithSigner(signer signer.Signer) (*Validate, error)
	validateAlgorithmWithJwtAndSigner(jwt Jwt, sg signer.Signer) error
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
	err := v.validateAlgorithmWithJwtAndSigner(v.Jwt, signer)
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

func (v Validate) validateDate() (Validate, error) {
	truncate := time.Now().Truncate(time.Second)
	return v.validateDateWithDate(truncate.Unix())
}
func (v Validate) validateDateWithDate(date int64) (Validate, error) {

	err := v.validateDateWithDateAndLeeway(date, 0)
	if err != nil {
		return v, err
	}
	return v, nil
}
func (v Validate) validateDateWithDateAndLeeway(date int64, leeway int64) error {
	now := time.Now().Unix()

	// 检查生效时间（生效时间不能晚于当前时间）
	notBefore := v.Payload.getClaim(NOT_BEFORE)
	switch t := notBefore.(type) {
	case string:
		i, _ := strconv.ParseInt(t, 10, 64)
		validateNotAfter(NOT_BEFORE, i, now, leeway)
	case int, int8, int16, int32, int64:

	}

	// 检查失效时间（失效时间不能早于当前时间）
	expiresAt := v.Payload.getClaim(EXPIRES_AT)
	switch e := expiresAt.(type) {
	case string:
		i, _ := strconv.ParseInt(e, 10, 64)
		err := validateNotBefore(EXPIRES_AT, i, now, leeway)
		if err != nil {
			return err
		}
	case int, int8, int16, int32, int64:

	}

	// 检查签发时间（签发时间不能晚于当前时间）
	issueAt := v.Payload.getClaim(ISSUED_AT)
	switch ia := issueAt.(type) {
	case string:
		i, _ := strconv.ParseInt(ia, 10, 64)
		err := validateNotAfter(ISSUED_AT, i, now, leeway)
		if err != nil {
			return err
		}
	case int, int8, int16, int32, int64:

	}
	return nil
}

func validateNotAfter(fieldName string, dateToCheck int64, now int64, leeway int64) error {
	if dateToCheck == 0 {
		return nil
	}
	if leeway > 0 {
		dateToCheck += leeway
	}
	if dateToCheck > now {
		return errors.New(fmt.Sprintf("'%s':[%d] is after now:[%d]", fieldName, dateToCheck, now))
	}
	return nil
}
func validateNotBefore(fieldName string, dateToCheck int64, now int64, leeway int64) error {
	if dateToCheck == 0 {
		return nil
	}
	if leeway > 0 {
		dateToCheck -= leeway
	}
	if dateToCheck < now {
		return errors.New(fmt.Sprintf("'%s':[%d] is before now:[%d]", fieldName, dateToCheck, now))
	}
	return nil
}
