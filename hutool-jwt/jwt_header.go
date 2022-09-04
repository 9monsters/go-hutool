package jwt

const (
	ALGORITHM   = "alg"
	TYPE        = "typ"
	ContentType = "cty"
	KeyId       = "kid"
)

type Header struct {
	Claims
}
type Headerer interface {
	SetKeyId(keyId string) Header
	AddHeaders(data map[string]any) Header
}

func (h *Header) SetKeyId(keyId string) *Header {
	h.Claims.SetClaim(KeyId, keyId)
	return h
}

func (h *Header) AddHeaders(data map[string]any) *Header {
	h.Claims.PutAll(data)
	return h
}

func (h *Header) Sign() string {

}
