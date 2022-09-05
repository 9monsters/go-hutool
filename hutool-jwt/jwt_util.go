package jwt

func CreateToken(payload map[string]any, key []byte) string {
	return CreateTokenWithHeader(nil, payload, key)
}

func CreateTokenWithHeader(headers map[string]any, payloads map[string]any, key []byte) string {
	jwt := Create()
	h := *createHeader()
	h.addHeaders(headers)
	jwt.Header = h

	p := *createPayload()
	p.addPayloads(payloads)
	jwt.Payload = p

	jwt.setKey(key)
	return jwt.sign()
}
