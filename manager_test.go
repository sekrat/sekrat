package sekrat

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	warehouse := NewMemoryWarehouse()
	crypter := NewPassthroughCrypter()

	manager := New(warehouse, crypter)

	t.Run("it uses the provided warehouse", func(t *testing.T) {
		if manager.Warehouse != warehouse {
			t.Errorf("Did not get the expected warehouse")
		}
	})

	t.Run("it uses the provided crypter", func(t *testing.T) {
		if manager.Crypter != crypter {
			t.Errorf("Did not get the expected crypter")
		}
	})
}

func TestManager_IDs(t *testing.T) {
	id := "id1"
	data := []byte("i am some data")

	warehouse := NewMemoryWarehouse()
	crypter := NewPassthroughCrypter()
	manager := New(warehouse, crypter)

	t.Run("when there are no secrets", func(t *testing.T) {
		warehouse.Storage = nil

		ids := manager.IDs()

		t.Run("it is empty", func(t *testing.T) {
			if len(ids) != 0 {
				t.Errorf("Expected 0 IDs, got %d", len(ids))
			}
		})

	})

	t.Run("when there are secrets", func(t *testing.T) {
		warehouse.Storage = nil
		warehouse.Store(id, data)
		ids := manager.IDs()

		t.Run("it contains the ID for all known secrets", func(t *testing.T) {
			if !contains(id, ids) {
				t.Errorf("Expected '%s' to be present", id)
			}
		})
	})
}

func TestManager_Put(t *testing.T) {
	id := "id1"
	key := "put"
	data := []byte("put test")
	warehouse := &TestWarehouse{}
	crypter := &TestCrypter{}
	manager := New(warehouse, crypter)

	t.Run("when everything goes well", func(t *testing.T) {
		t.Run("and the secret is not already stored", func(t *testing.T) {
			warehouse.Reset()
			crypter.Reset()

			err := manager.Put(id, key, data)

			t.Run("it encrypts the data", func(t *testing.T) {
				if crypter.EncryptKey != key || string(data) != string(crypter.Data) {
					t.Errorf("Expected the crypter to be called")
				}
			})

			t.Run("it stores the data", func(t *testing.T) {
				if string(warehouse.Wrapped.Storage[id]) != string(data) {
					t.Errorf("Expected the secret to be '%s', got '%s'", string(data), string(warehouse.Wrapped.Storage[id]))
				}
			})

			t.Run("it returns no error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected no error, got '%s'", err.Error())
				}
			})
		})

		t.Run("and the secret is already stored", func(t *testing.T) {
			warehouse.Reset()
			crypter.Reset()
			warehouse.Store(id, data)

			newdata := []byte("overwrite test")

			err := manager.Put(id, key, newdata)

			t.Run("it encrypts the new data", func(t *testing.T) {
				if crypter.EncryptKey != key || string(newdata) != string(crypter.Data) {
					t.Errorf("Expected the crypter to be called")
				}

			})

			t.Run("it overwrites the old data", func(t *testing.T) {
				if string(warehouse.Wrapped.Storage[id]) != string(newdata) {
					t.Errorf("Expected the secret to be '%s', got '%s'", string(newdata), string(warehouse.Wrapped.Storage[id]))
				}
			})

			t.Run("it returns no error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected no error, got '%s'", err.Error())
				}
			})
		})
	})

	t.Run("when the crypter returns an error", func(t *testing.T) {
		warehouse.Reset()
		crypter.Reset()
		crypter.Buggy = true

		err := manager.Put(id, key, data)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("when the warehouse returns an error", func(t *testing.T) {
		warehouse.Reset()
		crypter.Reset()
		warehouse.Buggy = true

		err := manager.Put(id, key, data)

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func TestManager_Get(t *testing.T) {
	id := "id1"
	key := "get"
	data := []byte("get test")
	warehouse := &TestWarehouse{}
	crypter := &TestCrypter{}
	manager := New(warehouse, crypter)

	t.Run("when everything goes well", func(t *testing.T) {
		warehouse.Reset()
		crypter.Reset()
		warehouse.Store(id, data)

		result, err := manager.Get(id, key)

		t.Run("it decrypts the crypted secret", func(t *testing.T) {
			if crypter.DecryptKey != key || string(crypter.Data) != string(data) {
				t.Errorf("Expected the crypter to be called")
			}
		})

		t.Run("it returns the decrypted secret", func(t *testing.T) {
			if string(result) != string(data) {
				t.Errorf("Expected the result to be '%s', got '%s'", string(data), string(result))
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, got '%s'", err.Error())
			}
		})
	})

	t.Run("when the crypter returns an error", func(t *testing.T) {
		warehouse.Reset()
		crypter.Reset()
		warehouse.Store(id, data)
		crypter.Buggy = true

		result, err := manager.Get(id, key)

		t.Run("it returns no secret", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no secret, got '%s'", string(result))
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("when the warehouse returns an error", func(t *testing.T) {
		warehouse.Reset()
		crypter.Reset()
		warehouse.Store(id, data)
		warehouse.Buggy = true

		result, err := manager.Get(id, key)

		t.Run("it returns no secret", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no secret, got '%s'", string(result))
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

type TestWarehouse struct {
	Buggy   bool
	Wrapped *MemoryWarehouse
}

func (warehouse *TestWarehouse) IDs() []string {
	return warehouse.Wrapped.IDs()
}

func (warehouse *TestWarehouse) Store(id string, data []byte) error {
	if warehouse.Buggy {
		return errors.New("StorageFailure")
	}

	return warehouse.Wrapped.Store(id, data)
}

func (warehouse *TestWarehouse) Retrieve(id string) ([]byte, error) {
	if warehouse.Buggy {
		return nil, errors.New("RetrieveError")
	}

	return warehouse.Wrapped.Retrieve(id)
}

func (warehouse *TestWarehouse) Reset() {
	warehouse.Buggy = false
	warehouse.Wrapped = NewMemoryWarehouse()
}

type TestCrypter struct {
	Buggy      bool
	EncryptKey string
	DecryptKey string
	Data       []byte
}

func (crypter *TestCrypter) Encrypt(key string, data []byte) ([]byte, error) {
	crypter.EncryptKey = key
	crypter.Data = data

	var err error

	if crypter.Buggy {
		err = errors.New("EncryptError")
	}

	return data, err
}

func (crypter *TestCrypter) Decrypt(key string, data []byte) ([]byte, error) {
	crypter.DecryptKey = key
	crypter.Data = data

	var err error

	if crypter.Buggy {
		err = errors.New("DecryptError")
	}

	return data, err
}

func (crypter *TestCrypter) Reset() {
	crypter.Buggy = false
	crypter.Data = nil
	crypter.EncryptKey = ""
	crypter.DecryptKey = ""
}
