package jwt

import (
	"encoding/json"
	"github.com/nine-monsters/go-hutool/hutool-core/codec"
)

type RegisteredClaims struct {
	Issuer    string       `json:"iss,omitempty"`
	Subject   string       `json:"sub,omitempty"`
	Audience  ClaimStrings `json:"aud,omitempty"`
	ExpiresAt *NumericDate `json:"exp,omitempty"`
	NotBefore *NumericDate `json:"nbf,omitempty"`
	IssuedAt  *NumericDate `json:"iat,omitempty"`
	ID        string       `json:"jti,omitempty"`
}

type Claims map[string]any

type claims interface {
	setClaim(name string, value any)
	putAll(headerClaims map[string]any)
	getClaim(name string) any
}

var claimsData Claims

func initClaims() {
	if &claimsData == nil {
		claimsData = make(map[string]any)
	}
}

func createClaims() *Claims {
	return &Claims{}
}

func (c Claims) setClaim(name string, value any) {
	if value == nil {
		delete(c, name)
	}
	c[name] = value
}

func (c Claims) putAll(headerClaims map[string]any) {
	if len(headerClaims) > 0 {
		for name, data := range headerClaims {
			c.setClaim(name, data)
		}
	}
}

func (c Claims) getClaim(name string) any {
	return c[name]
}

func ParseBase64(tokenPart string) *map[string]any {
	var data = make(map[string]any)
	decodeStr, _ := codec.Base64.DecodeStr(tokenPart)
	_ = json.Unmarshal(decodeStr, &data)
	return &data
}
