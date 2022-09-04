package jwt

// Signer JWT签名接口封装，通过实现此接口，完成不同算法的签名功能
type Signer interface {
	Sign(headerBase64 string, payloadBase64 string) string

	Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool

	GetAlgorithm() string
}
