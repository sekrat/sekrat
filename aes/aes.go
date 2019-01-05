package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"github.com/ess/pylades"
)

const (
	nonceSize = 16
)

type AES struct{}

func (crypter *AES) Encrypt(key string, data []byte) ([]byte, error) {
	hashKey := crypter.normalize(key)

	block, err := aes.NewCipher(hashKey)
	if err != nil {
		return nil, errors.New("could not set up cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.New("could not set up gcm")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.New("could not set up nonce")
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

func (crypter *AES) createMac(key []byte, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (crypter *AES) Decrypt(key string, data []byte) ([]byte, error) {
	hashKey := crypter.normalize(key)

	block, err := aes.NewCipher(hashKey)
	if err != nil {
		return nil, errors.New("could not set up cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.New("could not set up gcm")
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("could not decrypt")
	}

	return decrypted, nil
}

func (crypter *AES) normalize(key string) []byte {
	sum := make([]byte, 0)

	for _, item := range sha256.Sum256([]byte(key)) {
		sum = append(sum, item)
	}

	return sum
}

//func (crypter *AES) verify(key string, mac []byte, data []byte) bool {

//}

func New() pylades.Crypter {
	return &AES{}
}
