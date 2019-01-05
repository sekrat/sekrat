package memory

import (
	"errors"
)

type Memory struct {
	Storage map[string][]byte
}

func (warehouse *Memory) IDs() []string {
	warehouse.setup()

	keys := make([]string, 0)

	for id, _ := range warehouse.Storage {
		keys = append(keys, id)
	}

	return keys
}

func (warehouse *Memory) Store(id string, data []byte) error {
	warehouse.setup()

	warehouse.Storage[id] = data

	return nil
}

func (warehouse *Memory) Retrieve(id string) ([]byte, error) {
	warehouse.setup()

	found := false
	for _, key := range warehouse.IDs() {
		if key == id {
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("not found")
	}

	return warehouse.Storage[id], nil
}

func (warehouse *Memory) setup() {
	if warehouse.Storage == nil {
		warehouse.Storage = make(map[string][]byte)
	}
}

func New() *Memory {
	return &Memory{}
}
