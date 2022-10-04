package errors

import "errors"

var (
	ErrCipherKey        = errors.New("the secret key is wrong and cannot be decrypted. Please check")
	ErrRsaBits          = errors.New("bits 1024 or 2048")
	ErrCreateHmac       = errors.New("create hmac error")
	ErrHmacAlgorithm    = errors.New("not support algorithm")
	ErrPrivateKeyNil    = errors.New("private key must not null when use it ")
	ErrUnsupportKeyType = errors.New("unsupported key type")
)
