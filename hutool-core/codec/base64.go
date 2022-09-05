package codec

import b64 "encoding/base64"

type base64 interface {
	EncodeStr(seg []byte) string
	EncodeUrlSafeStr(seg []byte) string
	EncodeStrWithoutPadding(seg []byte) string
	EncodeUrlSafeStrWithoutPadding(seg []byte) string
	encoding(seg []byte, safe bool, raw bool) string

	DecodeStr(sEnc string) ([]byte, error)
	DecodeUrlSafeStr(sEnc string) ([]byte, error)
	DecodeStrWithoutPadding(sEnc string) ([]byte, error)
	DecodeUrlSafeStrWithoutPadding(sEnc string) ([]byte, error)
	decoding(enc string, safe bool, raw bool) ([]byte, error)
}

type base64Encoder struct {
	stdEncoding    *b64.Encoding
	rawStdEncoding *b64.Encoding
	urlEncoding    *b64.Encoding
	rawURLEncoding *b64.Encoding
}

var Base64 base64Encoder

func init() {
	Base64 = base64Encoder{
		stdEncoding:    b64.StdEncoding,
		rawStdEncoding: b64.RawStdEncoding,
		urlEncoding:    b64.URLEncoding,
		rawURLEncoding: b64.RawURLEncoding,
	}
}

func (ba base64Encoder) EncodeStr(seg []byte) string {
	return ba.encoding(seg, false, false)
}

func (ba base64Encoder) EncodeUrlSafeStr(seg []byte) string {
	return ba.encoding(seg, true, false)
}
func (ba base64Encoder) EncodeStrWithoutPadding(seg []byte) string {
	return ba.encoding(seg, false, true)
}

func (ba base64Encoder) EncodeUrlSafeStrWithoutPadding(seg []byte) string {
	return ba.encoding(seg, true, true)
}

func (ba base64Encoder) encoding(seg []byte, safe bool, raw bool) string {
	if raw {
		if safe {
			return ba.rawURLEncoding.EncodeToString(seg)
		} else {
			return ba.rawStdEncoding.EncodeToString(seg)
		}
	} else {
		if safe {
			return ba.urlEncoding.EncodeToString(seg)
		} else {
			return ba.stdEncoding.EncodeToString(seg)
		}
	}
}

func (ba base64Encoder) DecodeStr(sEnc string) ([]byte, error) {
	return ba.decoding(sEnc, false, false)
}

func (ba base64Encoder) DecodeUrlSafeStr(sEnc string) ([]byte, error) {
	return ba.decoding(sEnc, true, false)

}

func (ba base64Encoder) DecodeStrWithoutPadding(sEnc string) ([]byte, error) {
	return ba.decoding(sEnc, false, true)

}

func (ba base64Encoder) DecodeUrlSafeStrWithoutPadding(sEnc string) ([]byte, error) {
	return ba.decoding(sEnc, true, true)
}

func (ba base64Encoder) decoding(enc string, safe bool, raw bool) ([]byte, error) {
	if raw {
		if safe {
			return ba.rawURLEncoding.DecodeString(enc)
		} else {
			return ba.rawStdEncoding.DecodeString(enc)
		}
	} else {
		if safe {
			return ba.urlEncoding.DecodeString(enc)
		} else {
			return ba.stdEncoding.DecodeString(enc)
		}
	}
}
