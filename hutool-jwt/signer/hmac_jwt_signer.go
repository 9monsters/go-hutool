package signer

type HMacJWTSigner struct {
	JwtSigner
}

func (s *HMacJWTSigner) GetAlgorithmId() string {
	return GetId(s.GetAlgorithm())
}
