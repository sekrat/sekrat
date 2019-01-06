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
