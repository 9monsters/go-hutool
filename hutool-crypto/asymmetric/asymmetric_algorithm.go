package asymmetric

type asymmetricAlgorithm[T string] struct {
	value T
}

var (
	RSA           = asymmetricAlgorithm[string]{"RSA"}
	RSA_ECB_PKCS1 = asymmetricAlgorithm[string]{"RSA/ECB/PKCS1Padding"}
	RSA_ECB       = asymmetricAlgorithm[string]{"RSA/ECB/NoPadding"}
	RSA_None      = asymmetricAlgorithm[string]{"RSA/None/NoPadding"}
)
var asymmetricAlgorithms = []asymmetricAlgorithm[string]{
	RSA, RSA_ECB_PKCS1, RSA_ECB, RSA_None,
}

func AsymmetricAlgorithmValues() []asymmetricAlgorithm[string] {
	return asymmetricAlgorithms
}

func (r asymmetricAlgorithm[string]) GetValue() (value string) {
	return r.value
}
