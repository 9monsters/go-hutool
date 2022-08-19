package asymmetric

type SignAlgorithm struct {
	value string
}

var (
	NONEwithRSA = SignAlgorithm{"NONEwithRSA"}

	MD2withRSA = SignAlgorithm{"MD2withRSA"}
	MD5withRSA = SignAlgorithm{"MD5withRSA"}

	SHA1withRSA   = SignAlgorithm{"SHA1withRSA"}
	SHA256withRSA = SignAlgorithm{"SHA256withRSA"}
	SHA384withRSA = SignAlgorithm{"SHA384withRSA"}
	SHA512withRSA = SignAlgorithm{"SHA512withRSA"}

	NONEwithDSA = SignAlgorithm{"NONEwithDSA"}
	SHA1withDSA = SignAlgorithm{"SHA1withDSA"}

	NONEwithECDSA   = SignAlgorithm{"NONEwithECDSA"}
	SHA1withECDSA   = SignAlgorithm{"SHA1withECDSA"}
	SHA256withECDSA = SignAlgorithm{"SHA256withECDSA"}
	SHA384withECDSA = SignAlgorithm{"SHA384withECDSA"}
	SHA512withECDSA = SignAlgorithm{"SHA512withECDSA"}

	SHA256withRSA_PSS = SignAlgorithm{"SHA256WithRSA/PSS"}
	SHA384withRSA_PSS = SignAlgorithm{"SHA384WithRSA/PSS"}
	SHA512withRSA_PSS = SignAlgorithm{"SHA512WithRSA/PSS"}
)

var signAlgorithms = []SignAlgorithm{
	NONEwithRSA, MD2withRSA, MD5withRSA,
	SHA1withRSA, SHA256withRSA, SHA384withRSA, SHA512withRSA,
	NONEwithDSA, SHA1withDSA,
	NONEwithECDSA, SHA1withECDSA, SHA256withECDSA, SHA384withECDSA, SHA512withECDSA,
	SHA256withRSA_PSS, SHA384withRSA_PSS, SHA512withRSA_PSS,
}

func Values() []SignAlgorithm {
	return signAlgorithms
}

func (r SignAlgorithm) GetValue() (value string) {
	return r.value
}
