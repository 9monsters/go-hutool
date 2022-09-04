package jwt

import (
	"fmt"
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"github.com/nine-monsters/go-hutool/hutool-core/text"
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
)

var DecodePaddingAllowed bool

type Jwter interface {
	Claimer
	Headerer
	Payloader
	signer.JwtSigner
}

type Jwt struct {
	*Claims
	*Header
	*Payload
	tokens []string
}

func Create() *Jwt {
	return &Jwt{
		Claims:  &Claims{},
		Header:  &Header{},
		Payload: &Payload{},
		tokens:  nil,
	}
}

func (j *Jwt) SetPayload(str string, data any) {
	j.SetClaim(str, data)
}

func CreateWithToken(token string) *Jwt {
	jwt := Create()
	jwt.Pares(token)
	return jwt
}

func (j *Jwt) Pares(token string) {
	tokens := splitToken(token)
	j.tokens = tokens
	j.Header.Parse(j.tokens[0])
	j.Payload.Parse(j.tokens[1])
}

func splitToken(token string) []string {
	tokens := strings.Split(token, text.DOT)
	if len(tokens) != 3 {
		msg, _ := fmt.Printf("The token was expected 3 parts, but got %d.", len(tokens))
		panic(msg)
	}
	return tokens
}

func encodeSegment(seg []byte) string {
	return codec.Base64.EncodeStr(seg)
}

func decodeSegment(seg string) ([]byte, error) {
	if DecodePaddingAllowed {
		if l := len(seg) % 4; l > 0 {
			seg += strings.Repeat("=", 4-l)
		}
		return codec.Base64.DecodeUrlSafeStrWithoutPadding(seg), nil
	}

	return codec.Base64.DecodeUrlSafeStr(seg), nil
}
