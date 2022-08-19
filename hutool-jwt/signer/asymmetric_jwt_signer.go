package signer

type AsymmetricJWTSigner struct {
	JwtSigner
}

func (s *AsymmetricJWTSigner) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
