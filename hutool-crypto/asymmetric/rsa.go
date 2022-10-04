package asymmetric

import "crypto/rsa"

var ALGORITHM_RSA = RSA_ECB_PKCS1

type Rsa struct {
	AsymmetricCrypto
}

func generatePrivateKey(modulus int, privateExponent int) rsa.PrivateKey {
	return rsa.PrivateKey{}
}
func generatePublicKey(modulus int, privateExponent int) rsa.PublicKey {
	return rsa.PublicKey{}
}

func (r *Rsa) decrypt(bytes []byte, keyType KeyType[int]) []byte {
	return nil
}
