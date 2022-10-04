package asymmetric

// AsymmetricDecryptor asymmetric decryptor
type AsymmetricDecrypt struct {
}
type AsymmetricDecryptor interface {
	decrypt(bytes []byte, keyType KeyType[int]) []byte
}

func (ad AsymmetricDecrypt) decrypt(bytes []byte, keyType KeyType[int]) []byte {

}
