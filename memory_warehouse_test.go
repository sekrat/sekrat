package sekrat

import (
	"testing"
)

func TestNewMemoryWarehouse(t *testing.T) {
	warehouse := NewMemoryWarehouse()

	t.Run("it is a MemoryWarehouse", func(t *testing.T) {
		_, ok := warehouse.(*MemoryWarehouse)

		if !ok {
			t.Errorf("Expected a memory warehouse, but got something else")
		}
	})
}

func TestMemoryWarehouse_IDs(t *testing.T) {
	warehouse := &MemoryWarehouse{}

	t.Run("when there are no secrets in the warehouse", func(t *testing.T) {
		warehouse.Storage = nil

		t.Run("it is empty", func(t *testing.T) {
			numIDs := len(warehouse.IDs())

			if numIDs != 0 {
				t.Errorf("Expected 0 IDs, got %d", numIDs)

			}
		})
	})

	t.Run("when there are secrets in the warehouse", func(t *testing.T) {
		warehouse.Storage = make(map[string][]byte)
		id1 := "id1"
		id2 := "id2"

		warehouse.Storage[id1] = []byte("secret 1")
		warehouse.Storage[id2] = []byte("secret 2")

		ids := warehouse.IDs()

		t.Run("it contains all known secret IDs", func(t *testing.T) {
			numIDs := len(ids)

			if numIDs != 2 {
				t.Errorf("Expected 2 IDs, got %d", numIDs)
			}

			if !contains(id1, ids) {
				t.Errorf("Did not find '%s' in the IDs", id1)
			}

			if !contains(id2, ids) {
				t.Errorf("Did not find '%s' in the IDs", id2)
			}

		})
	})
}

func TestMemoryWarehouse_Store(t *testing.T) {
	warehouse := &MemoryWarehouse{}
	id := "id1"
	secret := []byte("omg secret")

	t.Run("when the secret does not yet exist", func(t *testing.T) {
		warehouse.Storage = nil

		err := warehouse.Store(id, secret)

		t.Run("it records the secret", func(t *testing.T) {
			if !contains(id, warehouse.IDs()) {
				t.Errorf("Expected '%s' to be in the IDs", id)
			}

			if string(warehouse.Storage[id]) != string(secret) {
				t.Errorf("Expected the secret to be '%s', got '%s'", string(secret), string(warehouse.Storage[id]))
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, got '%s'", err.Error())
			}
		})

	})

	t.Run("when the secret already exixsts", func(t *testing.T) {
		warehouse.Storage = make(map[string][]byte)
		warehouse.Storage[id] = secret

		newsecret := []byte("oh what a lovely bunch of bananas")

		err := warehouse.Store(id, newsecret)

		t.Run("it records the new secret", func(t *testing.T) {
			if !contains(id, warehouse.IDs()) {
				t.Errorf("Could not find the secret '%s'", id)
			}

			if string(warehouse.Storage[id]) != string(newsecret) {
				t.Errorf("Expected secret to be '%s', got '%s'", string(newsecret), string(warehouse.Storage[id]))
			}
		})

		t.Run("it returns a nil error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, got '%s'", err.Error())
			}
		})
	})
}

func TestMemoryWarehouse_Retrieve(t *testing.T) {
	id := "id1"
	secret := []byte("my sausages turned to gold!")

	warehouse := &MemoryWarehouse{}

	t.Run("when the secret is known", func(t *testing.T) {
		warehouse.Storage = map[string][]byte{id: secret}

		result, err := warehouse.Retrieve(id)

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, got '%s'", err.Error())
			}
		})

		t.Run("it returns the requested secret", func(t *testing.T) {
			if string(result) != string(secret) {
				t.Errorf("Expected the secret to be '%s', got '%s'", string(secret), string(result))
			}
		})
	})

	t.Run("when the secret is not known", func(t *testing.T) {
		warehouse.Storage = make(map[string][]byte)

		result, err := warehouse.Retrieve(id)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})

		t.Run("it returns no secret", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no secret, got '%s'", string(result))
			}
		})
	})
}
