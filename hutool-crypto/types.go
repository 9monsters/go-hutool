package crypto

import "github.com/nine-monsters/go-hutool/hutool-core/constraints"

type EnumFunc[T constraints.Ordered] interface {
	GetValue() (value T)
}
