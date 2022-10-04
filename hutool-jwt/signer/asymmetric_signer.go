package signer

// AsymmetricSigner AsymmetricSigner
type AsymmetricSigner struct {
}

func (as *AsymmetricSigner) Sign(headerBase64 string, payloadBase64 string) string {

}

func (as *AsymmetricSigner) Verify(headerBase64 string, payloadBase64 string, signBase64 string) bool {

}

func (as *AsymmetricSigner) GetAlgorithm() string {

}
