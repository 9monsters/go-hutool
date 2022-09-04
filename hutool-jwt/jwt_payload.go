package jwt

import (
	"encoding/json"
	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	"time"
)

const (
	ISSUER    = "iss"
	SUBJECT   = "sub"
	AUDIENCE  = "aud"
	ExpiresAt = "exp"
	NotBefore = "nbf"
	IssuedAt  = "iat"
	Id        = "jti"
)

type Payloader interface {
	Claimer
	SetPayload(str string, data any)
	SetIssuer(issuer string)
	SetSubject(subject string)
	SetAudience(audience ...string)
	SetExpiresAt(expiresAt time.Time)
	SetNotBefore(norBefore time.Time)
	SetIssuedAt(issuedAt time.Time)
	SetJWTId(jwtId string)
}

type Payload struct {
}

func (p *Payload) AddPayloads(payload map[string]any) {

}
func (p *Payload) SetPayload(name string, data any) {
}

func (p *Payload) SetIssuer(issuer string) {
	p.SetPayload(ISSUER, issuer)
}
func (p *Payload) SetSubject(issuer string) {
	p.SetPayload(SUBJECT, issuer)
}
func (p *Payload) SetAudience(issuer ...string) {
	p.SetPayload(AUDIENCE, issuer)
}
func (p *Payload) SetExpiresAt(issuer int64) {
	p.SetPayload(ExpiresAt, issuer)
}
func (p *Payload) SetNotBefore(issuer int64) {
	p.SetPayload(NotBefore, issuer)
}
func (p *Payload) SetIssuedAt(issuedAt int64) {
	p.SetPayload(IssuedAt, issuedAt)
}
func (p *Payload) SetJWTId(jwtId string) {
	p.SetPayload(Id, jwtId)
}

func (p *Payload) Parse(tokenPart string) {
	str := codec.Base64.DecodeStr(tokenPart)
	mapData := make(map[string]any, 5)
	err := json.Unmarshal([]byte(str), &mapData)
	if err != nil {
		p.claims = mapData
	}
}
