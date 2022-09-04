package signer

import "sync"

var signingMethods = map[string]func() JwtSigner{}
var signingMethodLock = new(sync.RWMutex)

// JwtSigner JWT签名接口封装，通过实现此接口，完成不同算法的签名功能
type JwtSigner interface {
	Sign(signingString string, key interface{}) (string, error)

	Verify(signingString, signature string, key interface{}) (bool, error)

	GetAlgorithm() string

	GetAlgorithmId() string
}

type JwtSign struct {
}

//func (s *JwtSign) Sign(headerBase64 string, payloadBase64 string) string {
//	return text.EMPTY
//}
//
//func (s *JwtSign) Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool {
//	return true
//}
//
//func (s *JwtSign) GetAlgorithm() string {
//	return text.EMPTY
//}

//func (s *JwtSign) GetAlgorithmId() string {
//	return GetId(s.GetAlgorithm())
//}

func RegisterJwtSigner(alg string, f func() JwtSigner) {
	signingMethodLock.Lock()
	defer signingMethodLock.Unlock()

	signingMethods[alg] = f
}

func GetSigningMethod(alg string) (signer JwtSigner) {
	signingMethodLock.RLock()
	defer signingMethodLock.RUnlock()

	if methodF, ok := signingMethods[alg]; ok {
		signer = methodF()
	}
	return
}

func GetAlgorithms() (algs []string) {
	signingMethodLock.RLock()
	defer signingMethodLock.RUnlock()

	for alg := range signingMethods {
		algs = append(algs, alg)
	}
	return
}
