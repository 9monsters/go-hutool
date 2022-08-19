package jwt

type Jwt struct {
	Claims
	Header
	Payload
}

func (j *Jwt) SetPayload(str string, data any) {
	j.SetClaim(str, data)
}

func New() *Jwt {
	return &Jwt{}
}
