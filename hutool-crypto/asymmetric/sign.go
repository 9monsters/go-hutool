package asymmetric

import "crypto/x509"

type Sign struct {
	BaseAsymmetric

	signature x509.SignatureAlgorithm
}
