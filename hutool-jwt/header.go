package jwt

const (
	Algorithm   = "alg"
	Type        = "typ"
	ContentType = "cty"
	KeyId       = "kid"
)

type RegisteredHeaders struct {
	Algorithm   string `json:"alg,omitempty"`
	Type        string `json:"typ,omitempty"`
	ContentType string `json:"cty,omitempty"`
	KeyId       string `json:"kid,omitempty"`
}

type Header map[string]any

type headers interface {
	claims
	addHeaders(headerClaims map[string]any)
}

var headerData Header

func initHeaders() {
	if &headerData == nil {
		headerData = make(map[string]any)
	}
}

func createHeader() *Header {
	return &Header{}
}

func (h Header) addHeaders(headers map[string]any) {
	h.putAll(headers)
}

func (h Header) setClaim(name string, value any) {
	if value == nil {
		delete(h, name)
	}
	h[name] = value
}

func (h Header) putAll(headers map[string]any) {
	if len(headers) > 0 {
		for name, data := range headers {
			h.setClaim(name, data)
		}
	}
}

func (h Header) getClaim(name string) any {
	return h[name]
}
