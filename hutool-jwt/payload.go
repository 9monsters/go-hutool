package jwt

type Payload map[string]any

type payloads interface {
	claims
	addPayloads(payloadClaims map[string]any)
	setPayload(name string, value any)
}

var payloadData map[string]any

func initPayload() {
	if &payloadData == nil {
		payloadData = make(map[string]any)
	}
}

func createPayload() *Payload {
	return &Payload{}
}

func (p Payload) addPayloads(payloads map[string]any) {
	p.putAll(payloads)
}

func (p Payload) setPayload(name string, value any) {
	p.setClaim(name, value)
}

func (p Payload) setClaim(name string, value any) {
	if value == nil {
		delete(p, name)
	}
	p[name] = value
}

func (p Payload) putAll(payloads map[string]any) {
	if len(payloads) > 0 {
		for name, data := range payloads {
			p.setClaim(name, data)
		}
	}
}

func (p Payload) getClaim(name string) any {
	return p[name]
}
