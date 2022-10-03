package asymmetric

import "crypto/rsa"

// Rsa rsa
type Rsa struct {
	PrivateKey    string
	PublicKey     string
	RsaPrivateKey rsa.PrivateKey
	RsaPublicKey  rsa.PublicKey
}
