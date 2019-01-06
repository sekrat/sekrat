package sekrat

import (
	"errors"
)

// MemoryWarehouse is a minimal reference implementation for the Warehouse
// interface. It stores all secrets in an in-memory map.
type MemoryWarehouse struct {
	Storage map[string][]byte
}

// IDs returns the IDs for all secrets stored in the warehouse.
func (warehouse *MemoryWarehouse) IDs() []string {
	warehouse.setup()

	keys := make([]string, 0)

	for id := range warehouse.Storage {
		keys = append(keys, id)
	}

	return keys
}

// Store takes a secret ID and a secret, stores the secret indexed by the ID,
// and returns an error. In this case, the error is always nil.
func (warehouse *MemoryWarehouse) Store(id string, data []byte) error {
	warehouse.setup()

	warehouse.Storage[id] = data

	return nil
}

// Retrieve takes a secret ID and returns the secret and an error. If the given
// ID is not known to the warehouse, the secret is nil and the error is not.
// Otherwise, the secret is populated and the error is nil.
func (warehouse *MemoryWarehouse) Retrieve(id string) ([]byte, error) {
	warehouse.setup()

	if !contains(id, warehouse.IDs()) {
		return nil, errors.New("not found")
	}

	return warehouse.Storage[id], nil
}

func (warehouse *MemoryWarehouse) setup() {
	if warehouse.Storage == nil {
		warehouse.Storage = make(map[string][]byte)
	}
}

// NewMemoryWarehouse creates a new MemoryWarehouse for use with a Manager
func NewMemoryWarehouse() *MemoryWarehouse {
	return &MemoryWarehouse{}
}

func contains(item string, ary []string) bool {
	for _, k := range ary {
		if k == item {
			return true
		}
	}

	return false
}
