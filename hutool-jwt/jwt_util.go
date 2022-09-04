package jwt

func CreateToken(payload map[string]any, key []rune) string {
	return CreateTokenWithHeader(nil, payload, key)
}

func CreateTokenWithHeader(headers map[string]any, payload map[string]any, key []rune) string {
	jwt := Create()
	jwt.AddHeaders(headers)
	jwt.AddPayloads(payload)
	jwt.SetKeyId()
	sign := jwt.Sign()

	return sign
}

func CreateTokenWithSinger(payload map[string]any, key []rune) string {
	return CreateTokenWithSingerHeader(nil, payload, key)
}

func CreateTokenWithSingerHeader(headers map[string]any, payload map[string]any, key []rune) string {

}

func ParseToken(token string) *JWT {
	return WithToken(token)
}
