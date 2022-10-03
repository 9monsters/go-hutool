package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"runtime"

	codec "github.com/nine-monsters/go-hutool/hutool-core/codec"
	digest "github.com/nine-monsters/go-hutool/hutool-crypto/digest"
	log "github.com/nine-monsters/go-hutool/hutool-log"
)

func rsaSign(msg, priKey []byte) (sign []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				log.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(priKey)
	hashed := digest.Sha256(msg)
	sign, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func rsaVerifySign(msg []byte, sign []byte, pubKey []byte) bool {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				log.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	publicKey, err := x509.ParsePKCS1PublicKey(pubKey)
	if err != nil {
		return false
	}

	hashed := digest.Sha256(msg)
	result := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, sign)
	return result == nil
}

// SignBase64 rsa SignBase64
func SignBase64(msg []byte, base64PriKey string) (base64Sign string, err error) {
	priBytes, err := codec.Base64.DecodeStr(base64PriKey)
	if err != nil {
		return "", err
	}
	sign, err := rsaSign(msg, priBytes)
	if err != nil {
		return "", err
	}
	return codec.Base64.EncodeStr(sign), nil
}

// VerifySignBase64 rsa VerifySignBase64
func VerifySignBase64(msg []byte, base64Sign, base64PubKey string) bool {
	signBytes, err := codec.Base64.DecodeStr(base64Sign)
	if err != nil {
		return false
	}
	pubBytes, err := codec.Base64.DecodeStr(base64PubKey)
	if err != nil {
		return false
	}
	return rsaVerifySign(msg, signBytes, pubBytes)
}

// SignHex rsa SignHex
func SignHex(msg []byte, hexPriKey string) (hexSign string, err error) {
	priBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return "", err
	}
	sign, err := rsaSign(msg, priBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sign), nil
}

// VerifySignHex rsa VerifySignHex
func VerifySignHex(msg []byte, hexSign, hexPubKey string) bool {
	signBytes, err := hex.DecodeString(hexSign)
	if err != nil {
		return false
	}
	pubBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return false
	}
	return rsaVerifySign(msg, signBytes, pubBytes)
}
