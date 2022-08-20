package codec

import b64 "encoding/base64"

type Base64Encoder struct {
	StdEncoding    *b64.Encoding
	RawStdEncoding *b64.Encoding
	URLEncoding    *b64.Encoding
	RawURLEncoding *b64.Encoding
}

var Base64 = Base64Encoder{
	StdEncoding:    b64.StdEncoding,
	RawStdEncoding: b64.RawStdEncoding,
	URLEncoding:    b64.URLEncoding,
	RawURLEncoding: b64.RawURLEncoding,
}

func (ba Base64Encoder) EncodeStr(source string) string {
	return ba.StdEncoding.EncodeToString([]byte(source))
}

func (ba Base64Encoder) EncodeUrlSafeStr(source string) string {
	return ba.URLEncoding.EncodeToString([]byte(source))
}
func (ba Base64Encoder) EncodeStrWithoutPadding(source string) string {
	return ba.RawStdEncoding.EncodeToString([]byte(source))
}

func (ba Base64Encoder) EncodeUrlSafeStrWithoutPadding(source string) string {
	return ba.RawURLEncoding.EncodeToString([]byte(source))
}

func (ba Base64Encoder) DecodeStr(sEnc string) string {
	sDec, err := ba.StdEncoding.DecodeString(sEnc)
	if err != nil {
		panic("Decode error!")
	}
	return string(sDec)
}

func (ba Base64Encoder) DecodeUrlSafeStr(sEnc string) string {
	sDec, err := ba.URLEncoding.DecodeString(sEnc)
	if err != nil {
		panic("Decode error!")
	}
	return string(sDec)

}

func (ba Base64Encoder) DecodeStrWithoutPadding(sEnc string) string {
	sDec, err := ba.RawStdEncoding.DecodeString(sEnc)
	if err != nil {
		panic("Decode error!")
	}
	return string(sDec)
}

func (ba Base64Encoder) DecodeUrlSafeStrWithoutPadding(sEnc string) string {
	sDec, err := ba.RawURLEncoding.DecodeString(sEnc)
	if err != nil {
		panic("Decode error!")
	}
	return string(sDec)
}
