package sekrat

// PassthroughCrypter is a Crypter implementation that doesn't actually encrypt
// or decrypt data. More than anything, this is just a reference implementation
// for a minimally funcational Crypter.
type PassthroughCrypter struct{}

// Encrypt takes an encryption key and a secret and returns an encrypted secret
// and an error. The resulting secret is identical to that which is passed in,
// and the error is always nil.
func (crypter *PassthroughCrypter) Encrypt(key string, data []byte) ([]byte, error) {
	return data, nil
}

// Decrypt takes an encryption key and a secret and returns a decrypted secret
// and an error. The resulting secret is identical to that which is passed in,
// and the error is always nil.
func (crypter *PassthroughCrypter) Decrypt(key string, data []byte) ([]byte, error) {
	return data, nil
}

// NewPassthroughCrypter creates and returns a PassthroughCrypter for use with
// a Manager.
func NewPassthroughCrypter() *PassthroughCrypter {
	return &PassthroughCrypter{}
}
