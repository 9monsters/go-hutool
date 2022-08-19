package crypto

type EnumFunc[T string | ~int] interface {
	string | ~int
	GetValue() (value T)
}
