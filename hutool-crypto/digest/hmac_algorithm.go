package digest

type HmacAlgorithm struct {
	value string
}

var (
	HmacMD5    = HmacAlgorithm{"HmacMD5"}
	HmacSHA1   = HmacAlgorithm{"HmacSHA1"}
	HmacSHA256 = HmacAlgorithm{"HmacSHA256"}
	HmacSHA384 = HmacAlgorithm{"HmacSHA384"}
	HmacSHA512 = HmacAlgorithm{"HmacSHA512"}

	HmacSM3 = HmacAlgorithm{"HmacSM3"}
	SM4CMAC = HmacAlgorithm{"SM4CMAC"}
)

var hmacAlgorithms = []HmacAlgorithm{
	HmacMD5, HmacSHA1, HmacSHA256, HmacSHA384, HmacSHA512, HmacSM3, SM4CMAC,
}

func Values() []HmacAlgorithm {
	return hmacAlgorithms
}

func (r *HmacAlgorithm) GetValue() string {
	return r.value
}
