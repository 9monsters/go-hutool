package signer

type NoneJWTSigner struct {
	JwtSigner
}

func (s *NoneJWTSigner) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
