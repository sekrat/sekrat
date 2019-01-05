package plain

import (
	"github.com/ess/pylades"
)

type Plain struct{}

func (crypter *Plain) Encrypt(key string, data []byte) ([]byte, error) {
	return data, nil
}

func (crypter *Plain) Decrypt(key string, data []byte) ([]byte, error) {
	return data, nil
}

func New() pylades.Crypter {
	return &Plain{}
}
