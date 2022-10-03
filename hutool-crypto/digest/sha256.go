package digest

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256Hex Sha256Hex
func Sha256Hex(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}

// Sha256 Sha256
func Sha256(data []byte) []byte {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}
