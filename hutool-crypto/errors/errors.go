package errors

import "errors"

var (
	ErrCipherKey = errors.New("The secret key is wrong and cannot be decrypted. Please check")
	ErrRsaBits   = errors.New("bits 1024 or 2048")
)
