package digest

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1Hex sha1 Hex
func Sha1Hex(data []byte) string {
	return hex.EncodeToString(Sha1(data))
}

// Sha1 sha1
func Sha1(data []byte) []byte {
	digest := sha1.New()
	digest.Write(data)
	return digest.Sum(nil)
}
