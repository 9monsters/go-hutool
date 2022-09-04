package signer

import (
	"github.com/nine-monsters/go-hutool/hutool-crypto/digest/hmac"
	jwt "github.com/nine-monsters/go-hutool/hutool-jwt"
)

func None() jwt.Signer {
	return NONE
}

func HS256(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS256"), key)
}

func HS384(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS384"), key)
}

func HS512(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS512"), key)
}

func RS256(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS256"), key)
}

func RS384(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS384"), key)
}

func RS512(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS512"), key)
}

func ES256(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS256"), key)
}

func ES384(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS384"), key)
}

func ES512(key string) jwt.Signer {
	return CreateSigner(GetAlgorithmById("HS512"), key)
}

// CreateSigner 创建签名器
// algorithmId 算法ID
func CreateSigner(algorithmId string, key string) jwt.Signer {
	if algorithmId == "" || IdNone == algorithmId {
		return None()
	}
	algorithm := GetAlgorithm(algorithmId)

	h, err := hmac.New(algorithm, key)
	if err != nil {
		return &HmacSigner{
			hmac: h,
		}
	}
	return nil
}
