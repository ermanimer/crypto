# crypto

[![Go Test](https://github.com/ermanimer/crypto/actions/workflows/go.yml/badge.svg)](https://github.com/ermanimer/crypto/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/crypto)](https://goreportcard.com/report/github.com/ermanimer/crypto)

crypto provides helper functions that uses standard crypto library.

# Installation

# Usage

The keys should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

## Encrypt

Encrypt encrypts provided data with provided key using AES.

```go
encrypted, err := crypto.Encrypt(key, decrypted)
```

## MustEncrypt

MustEncrypt calls Encrypt function and panics on errors.

```go
encrypted := crypto.MustEncrypt(key, decrypted)
```

## Decrypt

Decrypt decrypts encrypted data with provided key using AES.

```go
decrypted, err := crypto.Decrypt(key, encrypted)
```

## MustDecrypt

MustDecrypt calls Decrypt function and panics on errors.

```go
decrypted := crypto.MustDecrypt(key, encrypted)
```
