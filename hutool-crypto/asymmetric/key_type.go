package asymmetric

type KeyType[T int] struct {
	value T
}

var (
	PublicKey  = KeyType[int]{1}
	PrivateKey = KeyType[int]{2}
	SecretKey  = KeyType[int]{3}
)
