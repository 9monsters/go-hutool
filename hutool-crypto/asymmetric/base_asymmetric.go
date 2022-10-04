package asymmetric

import (
	"crypto"
	"github.com/nine-monsters/go-hutool/hutool-crypto/errors"
	"sync"

	"github.com/nine-monsters/go-hutool/hutool-core/text"
)

// BaseAsymmetric base Asymmetric
type BaseAsymmetric struct {
	algorithm  string
	privateKey crypto.PrivateKey
	publicKey  crypto.PublicKey

	lock sync.RWMutex
}

func (as *BaseAsymmetric) getPrivateKeyBase64() string {
	if as.privateKey == nil {
		return text.EMPTY
	}
	return ""
}

func (as *BaseAsymmetric) getKeyByType(keyType KeyType[int]) (crypto.PrivateKey, error) {
	switch keyType {
	case PrivateKey:
		if as.privateKey == nil {
			return nil, errors.ErrPrivateKeyNil
		}
		return as.privateKey, nil
	case PublicKey:
		if as.publicKey == nil {
			return nil, errors.ErrPrivateKeyNil
		}
		return as.publicKey, nil
	default:
		return nil, errors.ErrUnsupportKeyType
	}
}
