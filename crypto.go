// Package crypto provides helper functions that uses standard crypto library.
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"time"
)

// Encrypt encrypts provided data with provided key using AES.
func Encrypt(key []byte, decrypted []byte) ([]byte, error) {
	gcm, err := createGCM(key)
	if err != nil {
		return nil, err

	}

	nonce, err := generateNonce(gcm.NonceSize())
	if err != nil {
		return nil, err

	}

	encrypted := gcm.Seal(nonce, nonce, decrypted, nil)

	return encrypted, nil

}

// MustEncrypt calls Encrypt function and panics on errors.
func MustEncrypt(key []byte, decrypted []byte) []byte {
	encrypted, err := Encrypt(key, decrypted)
	if err != nil {
		panic(err.Error())

	}

	return encrypted

}

// Decrypt decrypts encrypted data with provided key using AES.
func Decrypt(key []byte, encrypted []byte) ([]byte, error) {
	gcm, err := createGCM(key)
	if err != nil {
		return nil, err

	}

	nonceSize := gcm.NonceSize()
	nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err

	}

	return decrypted, nil

}

// MustDecrypt calls Decrypt function and panics on errors.
func MustDecrypt(key []byte, encrypted []byte) []byte {
	decrypted, err := Decrypt(key, encrypted)
	if err != nil {
		panic(err.Error())

	}

	return decrypted

}

func createGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err

	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err

	}

	return gcm, nil

}

func generateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)

	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err

	}

	return nonce, nil

}
