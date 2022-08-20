package jwt

import (
	"fmt"
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/text"
)

type JWT struct {
	*Claims
	*Header
	*Payload
	tokens []string
}

func (j *JWT) SetPayload(str string, data any) {
	j.SetClaim(str, data)
}

func Create() *JWT {
	return WithNone()
}

func WithNone() *JWT {
	return &JWT{
		Claims:  &Claims{},
		Header:  &Header{},
		Payload: &Payload{},
		tokens:  nil,
	}
}

func WithToken(token string) *JWT {
	jwt := WithNone()
	jwt.pares(token)
	return jwt
}

func (j *JWT) pares(token string) {
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
