package digest

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512Hex Sha512Hex
func Sha512Hex(data []byte) string {
	return hex.EncodeToString(Sha512(data))
}

// Sha512 Sha512
func Sha512(data []byte) []byte {
	digest := sha512.New()
	digest.Write(data)
	return digest.Sum(nil)
}
