package pylades

import (
	"errors"
)

func New(warehouse Warehouse, crypter Crypter) *Manager {
	return &Manager{Warehouse: warehouse, Crypter: crypter}
}

type Manager struct {
	Warehouse Warehouse
	Crypter   Crypter
}

func (manager *Manager) IDs() []string {
	return manager.Warehouse.IDs()
}

func (manager *Manager) Put(id string, key string, data []byte) error {
	crypted, err := manager.Crypter.Encrypt(key, data)
	if err != nil {
		return errors.New("encryption failed")
	}

	err = manager.Warehouse.Store(id, crypted)
	if err != nil {
		return errors.New("storage failed")
	}

	return nil
}

func (manager *Manager) Get(id string, key string) ([]byte, error) {
	crypted, err := manager.Warehouse.Retrieve(id)
	if err != nil {
		return nil, errors.New("retrieval failed")
	}

	data, err := manager.Crypter.Decrypt(key, crypted)
	if err != nil {
		return nil, errors.New("decryption failed")
	}

	return data, nil
}

type Warehouse interface {
	IDs() []string
	Store(string, []byte) error
	Retrieve(string) ([]byte, error)
}

type Crypter interface {
	Encrypt(string, []byte) ([]byte, error)
	Decrypt(string, []byte) ([]byte, error)
}
