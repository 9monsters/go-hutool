package mac

type Macer interface {
	Digest(data string) []byte
	DigestBase64(data string, isUrlSafe bool) string
	DigestHex(data string) string

	Verify(digest []byte, digestToCompare []byte) bool
}
