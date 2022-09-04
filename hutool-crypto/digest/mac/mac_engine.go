package mac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/nine-monsters/go-hutool/hutool-crypto/digest"
)

func CreateEngine(algorithm string, key string) (hash.Hash, error) {
	keys := []byte(key)
	switch algorithm {
	case digest.HmacMD5.GetValue():
		return hmac.New(md5.New, keys), nil

	case digest.HmacSHA1.GetValue():
		return hmac.New(sha1.New, keys), nil

	case digest.HmacSHA256.GetValue():
		return hmac.New(sha256.New, keys), nil

	case digest.HmacSHA384.GetValue():
		return hmac.New(sha512.New384, keys), nil

	case digest.HmacSHA512.GetValue():
		return hmac.New(sha512.New, keys), nil

	default:
		panic("not support algorithm")
	}
}
