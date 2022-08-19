package jwt

const (
	ISSUER    = "iss"
	SUBJECT   = "sub"
	AUDIENCE  = "aud"
	ExpiresAt = "exp"
	NotBefore = "nbf"
	IssuedAt  = "iat"
	JwtId     = "jti"
)

type RegisteredPayload interface {
	SetPayload(str string, data any)
}

type registeredPayloadStruct struct {
	RegisteredPayload
}

func (jwt *registeredPayloadStruct) SetIssuer(issuer string) RegisteredPayload {
	jwt.SetPayload(ISSUER, issuer)
	return jwt
}
