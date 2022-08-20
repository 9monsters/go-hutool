package asymmetric

type SignAlgorithm[T string] struct {
	value T
}

var (
	NONEwithRSA = SignAlgorithm[string]{"NONEwithRSA"}

	MD2withRSA = SignAlgorithm[string]{"MD2withRSA"}
	MD5withRSA = SignAlgorithm[string]{"MD5withRSA"}

	SHA1withRSA   = SignAlgorithm[string]{"SHA1withRSA"}
	SHA256withRSA = SignAlgorithm[string]{"SHA256withRSA"}
	SHA384withRSA = SignAlgorithm[string]{"SHA384withRSA"}
	SHA512withRSA = SignAlgorithm[string]{"SHA512withRSA"}

	NONEwithDSA = SignAlgorithm[string]{"NONEwithDSA"}
	SHA1withDSA = SignAlgorithm[string]{"SHA1withDSA"}

	NONEwithECDSA   = SignAlgorithm[string]{"NONEwithECDSA"}
	SHA1withECDSA   = SignAlgorithm[string]{"SHA1withECDSA"}
	SHA256withECDSA = SignAlgorithm[string]{"SHA256withECDSA"}
	SHA384withECDSA = SignAlgorithm[string]{"SHA384withECDSA"}
	SHA512withECDSA = SignAlgorithm[string]{"SHA512withECDSA"}

	SHA256withRSA_PSS = SignAlgorithm[string]{"SHA256WithRSA/PSS"}
	SHA384withRSA_PSS = SignAlgorithm[string]{"SHA384WithRSA/PSS"}
	SHA512withRSA_PSS = SignAlgorithm[string]{"SHA512WithRSA/PSS"}
)

var signAlgorithms = []SignAlgorithm[string]{
	NONEwithRSA, MD2withRSA, MD5withRSA,
	SHA1withRSA, SHA256withRSA, SHA384withRSA, SHA512withRSA,
	NONEwithDSA, SHA1withDSA,
	NONEwithECDSA, SHA1withECDSA, SHA256withECDSA, SHA384withECDSA, SHA512withECDSA,
	SHA256withRSA_PSS, SHA384withRSA_PSS, SHA512withRSA_PSS,
}

func Values() []SignAlgorithm[string] {
	return signAlgorithms
}

func (r SignAlgorithm[string]) GetValue() (value string) {
	return r.value
}
