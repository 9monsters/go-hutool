package des

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	log "github.com/nine-monsters/go-hutool/hutool-log"
	"runtime"

	crypto "github.com/nine-monsters/go-hutool/hutool-crypto"
)

func DesCbcEncrypt(plainText, secretKey, ivDes []byte) (cipherText []byte, err error) {
	if len(secretKey) != 8 {
		return nil, crypto.ErrKeyLengthEight
	}
	block, err := des.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	paddingText := crypto.PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != block.BlockSize() {
			return nil, crypto.ErrIvDes
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(crypto.Ivdes)
	} // Initialization vector
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

func DesCbcDecrypt(cipherText, secretKey, ivDes []byte) (plainText []byte, err error) {
	if len(secretKey) != 8 {
		return nil, crypto.ErrKeyLengthEight
	}
	block, err := des.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

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

	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != block.BlockSize() {
			return nil, crypto.ErrIvDes
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(crypto.Ivdes)
	} // Initialization vector
	blockMode := cipher.NewCBCDecrypter(block, iv)

	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)

	unPaddingText, err := crypto.PKCS5UnPadding(plainText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return unPaddingText, nil
}

func DesCbcEncryptBase64(plainText, secretKey, ivAes []byte) (cipherTextBase64 string, err error) {
	encryBytes, err := DesCbcEncrypt(plainText, secretKey, ivAes)
	return base64.StdEncoding.EncodeToString(encryBytes), err
}

func DesCbcEncryptHex(plainText, secretKey, ivAes []byte) (cipherTextHex string, err error) {
	encryBytes, err := DesCbcEncrypt(plainText, secretKey, ivAes)
	return hex.EncodeToString(encryBytes), err
}

func DesCbcDecryptByBase64(cipherTextBase64 string, secretKey, ivAes []byte) (plainText []byte, err error) {
	plainTextBytes, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return DesCbcDecrypt(plainTextBytes, secretKey, ivAes)
}

func DesCbcDecryptByHex(cipherTextHex string, secretKey, ivAes []byte) (plainText []byte, err error) {
	plainTextBytes, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return DesCbcDecrypt(plainTextBytes, secretKey, ivAes)
}
