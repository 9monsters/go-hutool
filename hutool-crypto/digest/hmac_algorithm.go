package digest

type HmacAlgorithm[T string] struct {
	value T
}

var (
	HmacMD5    = HmacAlgorithm[string]{"HmacMD5"}
	HmacSHA1   = HmacAlgorithm[string]{"HmacSHA1"}
	HmacSHA256 = HmacAlgorithm[string]{"HmacSHA256"}
	HmacSHA384 = HmacAlgorithm[string]{"HmacSHA384"}
	HmacSHA512 = HmacAlgorithm[string]{"HmacSHA512"}

	HmacSM3 = HmacAlgorithm[string]{"HmacSM3"}
	SM4CMAC = HmacAlgorithm[string]{"SM4CMAC"}
)

var hmacAlgorithms = []HmacAlgorithm[string]{
	HmacMD5, HmacSHA1, HmacSHA256, HmacSHA384, HmacSHA512, HmacSM3, SM4CMAC,
}

func Values() []HmacAlgorithm[string] {
	return hmacAlgorithms
}

func (r HmacAlgorithm[string]) GetValue() string {
	return r.value
}
