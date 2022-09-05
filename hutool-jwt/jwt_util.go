package jwt

import "github.com/nine-monsters/go-hutool/hutool-jwt/signer"

func CreateToken(payload map[string]any, key []byte) string {
	return Create().
		withPayload(payload).
		setKey(key).
		sign()
}

func CreateTokenWithSigner(payload map[string]any, signer signer.Signer) string {
	return Create().
		withPayload(payload).
		setSigner(signer).
		sign()
}

func CreateTokenWithHeader(headers map[string]any, payloads map[string]any, key []byte) string {
	return Create().
		withPayload(payloads).
		withHeader(headers).
		setKey(key).
		sign()
}

func CreateTokenWithHeaderAndSigner(headers map[string]any, payloads map[string]any, signer signer.Signer) string {
	return Create().
		withPayload(payloads).
		withHeader(headers).
		setSigner(signer).
		sign()
}

func (j *Jwt) withHeader(headers map[string]any) *Jwt {
	h := *createHeader()
	h.addHeaders(headers)
	j.Header = h

	return j
}

func (j *Jwt) withPayload(payloads map[string]any) *Jwt {
	p := *createPayload()
	p.addPayloads(payloads)
	j.Payload = p

	return j
}

func Parse(token string) *Jwt {
	return of(token)
}
func Verify(token string, key []byte) bool {
	return of(token).
		setKey(key).
		verify()
}
