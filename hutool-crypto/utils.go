package crypto

import "errors"

var (
	ErrCipherKey           = errors.New("the secret key is wrong and cannot be decrypted. Please check")
	ErrKeyLengthSixteen    = errors.New("a sixteen or twenty-four or thirty-two length secret key is required")
	ErrKeyLengthEight      = errors.New("a eight-length secret key is required")
	ErrKeyLengthTwentyFour = errors.New("a twenty-four-length secret key is required")
	ErrPaddingSize         = errors.New("padding size error please check the secret key or iv")
	ErrIvAes               = errors.New("a sixteen-length ivaes is required")
	ErrIvDes               = errors.New("a eight-length ivdes key is required")
	ErrRsaBits             = errors.New("bits 1024 or 2048")
)

const (
	Ivaes = "1crypto112345678"
	Ivdes = "1crypto1"

	privateFileName = "private.pem"
	publicFileName  = "public.pem"

	eccPrivateFileName = "eccprivate.pem"
	eccPublishFileName = "eccpublic.pem"

	privateKeyPrefix = " WUMAN RSA PRIVATE KEY "
	publicKeyPrefix  = " WUMAN  RSA PUBLIC KEY "

	eccPrivateKeyPrefix = " WUMAN ECC PRIVATE KEY "
	eccPublicKeyPrefix  = " WUMAN ECC PUBLIC KEY "
)
