package jwt

import (
	"crypto/subtle"
	"encoding/json"
	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"time"
)

const (
	// ISSUER
	ISSUER = "iss"
	// SUBJECT
	SUBJECT = "sub"
	// AUDIENCE
	AUDIENCE = "aud"
	// EXPIRES_AT
	EXPIRES_AT = "exp"
	// NOT_BEFORE
	NOT_BEFORE = "nbf"
	// ISSUED_AT
	ISSUED_AT = "iat"
	// JWT_ID
	JWT_ID = "jti"
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

func verifyAud(aud []string, cmp string, required bool) bool {
	if len(aud) == 0 {
		return !required
	}
	// use a var here to keep constant time compare when looping over a number of claims
	result := false

	var stringClaims string
	for _, a := range aud {
		if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
			result = true
		}
		stringClaims = stringClaims + a
	}

	// case where "" is sent in one or many aud claims
	if len(stringClaims) == 0 {
		return !required
	}

	return result
}

func verifyExp(exp *time.Time, now time.Time, required bool) bool {
	if exp == nil {
		return !required
	}
	return now.Before(*exp)
}

func verifyIat(iat *time.Time, now time.Time, required bool) bool {
	if iat == nil {
		return !required
	}
	return now.After(*iat) || now.Equal(*iat)
}

func verifyNbf(nbf *time.Time, now time.Time, required bool) bool {
	if nbf == nil {
		return !required
	}
	return now.After(*nbf) || now.Equal(*nbf)
}

func verifyIss(iss string, cmp string, required bool) bool {
	if iss == "" {
		return !required
	}
	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}
