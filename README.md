# crypto

crypto provides helper functions that uses standard crypto library.

# Functions

- **Encrypt**(key []byte, decrypted []byte) ([]byte, error)
- **MustEncrypt**(key []byte, decrypted []byte) []byte
- **Decrypt**(key []byte, encrypted []byte) ([]byte, error)
- **MustDecrypt**(key []byte, encrypted []byte) []byte

The key should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
