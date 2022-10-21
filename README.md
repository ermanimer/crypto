# crypto

[![Go Test](https://github.com/ermanimer/crypto/actions/workflows/go.yml/badge.svg)](https://github.com/ermanimer/crypto/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/crypto)](https://goreportcard.com/report/github.com/ermanimer/crypto)

crypto provides helper functions that uses standard crypto library.

# Functions

- **Encrypt**(key []byte, decrypted []byte) ([]byte, error)
- **MustEncrypt**(key []byte, decrypted []byte) []byte
- **Decrypt**(key []byte, encrypted []byte) ([]byte, error)
- **MustDecrypt**(key []byte, encrypted []byte) []byte

The key should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
