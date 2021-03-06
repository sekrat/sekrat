package sekrat

import (
	"testing"
)

func TestPassthroughCrypter_Encrypt(t *testing.T) {
	key := "super secret"
	data := []byte("oh so secret data")
	crypter := &PassthroughCrypter{}

	result, err := crypter.Encrypt(key, data)

	t.Run("it returns no error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expected no error, got '%s'", err.Error())
		}
	})

	t.Run("it returns the data unchanged", func(t *testing.T) {
		if string(result) != string(data) {
			t.Errorf("Expected '%s', got '%s'", string(data), string(result))
		}
	})
}

func TestPassthroughCrypter_Decrypt(t *testing.T) {
	key := "super secret"
	data := []byte("oh so secret data")
	crypter := &PassthroughCrypter{}

	result, err := crypter.Decrypt(key, data)

	t.Run("it returns no error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expected no error, got '%s'", err.Error())
		}
	})

	t.Run("it returns the data unchanged", func(t *testing.T) {
		if string(result) != string(data) {
			t.Errorf("Expected '%s', got '%s'", string(data), string(result))
		}
	})

}
