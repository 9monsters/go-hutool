package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
)

type Jwt struct {
	Raw       string // The raw token.  Populated when you Parse a token
	Claims    Claims // The second segment jwtFromToken the token
	Signature string // The third segment jwtFromToken the token.  Populated when you Parse a token
	Valid     bool   // Is the token valid?  Populated when you Parse/Verify a token

	Header  Header        // The first segment jwtFromToken the token
	Signer  signer.Signer // The signing method used or to be used
	Payload Payload       // Is the token valid?  Populated when you Parse/Verify a token
	tokens  []string
}

type jwter interface {
	setKey(key []byte) *Jwt
	setSigner(signer signer.Signer) *Jwt
	sign() string
	signWithSigner(signer signer.Signer) string
	parse(token string) *Jwt
	verify() bool
	verifyWithSigner(s signer.Signer) bool
	validate(leeway int) bool
}

func Create() *Jwt {
	return &Jwt{
		Header:  *createHeader(),
		Payload: *createPayload(),
	}
}

func (j *Jwt) setKey(key []byte) *Jwt {
	j.setSigner(signer.HS256(key))
	return j
}
func (j *Jwt) setSigner(signer signer.Signer) *Jwt {
	j.Signer = signer
	return j
}

func (j *Jwt) sign() string {
	return j.signWithSigner(j.Signer)
}

func (j *Jwt) signWithSigner(sg signer.Signer) string {
	if sg == nil {
		panic("No Signer provided!")
	}
	alg := j.Header.getClaim(Algorithm)

	s, ok := alg.(string)
	// alg is empty
	if (ok && s == "") || (!ok && alg == nil) {
		j.Header.setClaim(Algorithm, signer.GetId(sg.GetAlgorithm()))
	}
	hJson, _ := json.Marshal(j.Header)
	hSafeStr := codec.Base64.EncodeUrlSafeStr(hJson)

	pJson, _ := json.Marshal(j.Payload)
	pSafeStr := codec.Base64.EncodeUrlSafeStr(pJson)

	sign := sg.Sign(hSafeStr, pSafeStr)
	return fmt.Sprintf("%s.%s.%s", hSafeStr, pSafeStr, sign)
}

func jwtFromToken(token string) *Jwt {
	jwt := Create()
	jwt.parse(token)
	return jwt
}

func (j *Jwt) parse(token string) *Jwt {
	j.Raw = token
	j.tokens = splitToken(token)
	j.Header = *ParseBase64(j.tokens[0])
	j.Payload = *ParseBase64(j.tokens[1])

	return j
}

func splitToken(token string) []string {
	tokens := strings.Split(token, text.DOT)
	if len(tokens) != 3 {
		panic(fmt.Sprintf("The token was expected 3 parts, but got %d.", len(tokens)))
	}
	return tokens
}

func (j *Jwt) verify() bool {
	return j.verifyWithSigner(j.Signer)
}

func (j *Jwt) verifyWithSigner(s signer.Signer) bool {
	if s == nil {
		s = signer.NONE
	}

	if len(j.tokens) == 0 {
		panic("No token to verify!")
	}
	if len(j.tokens) != 3 {
		panic(fmt.Sprintf("The token was expected 3 parts, but got %d.", len(j.tokens)))
	}

	return s.Verify(j.tokens[0], j.tokens[1], j.tokens[2])
}

func (j *Jwt) validate(leeway int) bool {
	if !j.verify() {
		return false
	}
	// todo 验证时间
	return true
}

func (j *Jwt) getAlgorithm() string {
	claim := j.Header.getClaim(Algorithm)
	switch a := claim.(type) {
	case string:
		return a
	}
	return ""
}

func (j *Jwt) getSigner() signer.Signer {
	return j.Signer
}
