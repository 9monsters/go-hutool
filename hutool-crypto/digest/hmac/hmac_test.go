package hmac

import (
	"testing"

	"github.com/nine-monsters/go-hutool/hutool-crypto/digest"
)

func TestNew(t *testing.T) {
	hmac, err := New(digest.HmacSHA256.GetValue(), []byte{})
	if err != nil {
		t.Fatal("error", err)
	}
	base64 := hmac.DigestBase64("123", true)
	t.Log("DigestBase64", base64)
	hex := hmac.DigestHex("123")
	t.Log("DigestHex", hex)
	t.Log("Algorithm", hmac.GetAlgorithm())
}
