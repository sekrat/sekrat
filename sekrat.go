// Package sekrat provides an embedded key/value store for an application that
// supports pluggable storage and encryption schemes.
package sekrat

import (
	"errors"
)

// New instantiats a Manager for managing secrets.
func New(warehouse Warehouse, crypter Crypter) *Manager {
	return &Manager{Warehouse: warehouse, Crypter: crypter}
}

// A Manager orchestrates operations between a Warehouse and a Crypter to
// securely store secrets, retrieve secrets, and list the IDs of all known
// secrets.
type Manager struct {
	Warehouse Warehouse
	Crypter   Crypter
}

// IDs returns the IDs for all secrets that the manager manages.
func (manager *Manager) IDs() []string {
	return manager.Warehouse.IDs()
}

// Put takes a secret ID, an encryption key, and a chunk of data, encrypts
// that data, and stores it indexed by ID. If all goes according to plan, a nil
// error is returned. Otherwise, an actual error is returned.
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

// Get takes a secret ID and an encryption key, and it returns a decrypted
// secret and an error. If all goes according to plan, the decrypted secret is
// populated and the error is nil. Otherwise, the secret is nil and the error
// is not.
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

// Warehouse is an interface that must be implemented by a storage plugin
type Warehouse interface {
	IDs() []string
	Store(string, []byte) error
	Retrieve(string) ([]byte, error)
}

// Crypter is an interface that must be implemented by an encryption plugin
type Crypter interface {
	Encrypt(string, []byte) ([]byte, error)
	Decrypt(string, []byte) ([]byte, error)
}

/*
Copyright 2019 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
