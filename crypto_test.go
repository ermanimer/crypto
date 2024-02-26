package crypto

import (
	"crypto/rand"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key, err := generateRandomData(16)
	if err != nil {
		t.Error(err.Error())
	}

	data, err := generateRandomData(256)
	if err != nil {
		t.Error(err.Error())
	}

	encrypted := MustEncrypt(key, data)

	decrypted := MustDecrypt(key, encrypted)

	if !isEqual(decrypted, data) {
		t.Error("decrypted != data")
	}
}

func TestEncryptWithShortKey(t *testing.T) {
	key, err := generateRandomData(15)
	if err != nil {
		t.Error(err.Error())
	}

	data, err := generateRandomData(256)
	if err != nil {
		t.Error(err.Error())
	}

	defer func() {
		if recover() == nil {
			t.Error("recover() == nil")
		}
	}()

	_ = MustEncrypt(key, data)
}

func TestDecryptWithShortKey(t *testing.T) {
	key, err := generateRandomData(15)
	if err != nil {
		t.Error(err.Error())
	}

	data, err := generateRandomData(256)
	if err != nil {
		t.Error(err.Error())
	}

	defer func() {
		if recover() == nil {
			t.Error("recover() == nil")
		}
	}()

	_ = MustDecrypt(key, data)
}

func generateRandomData(size int) ([]byte, error) {
	data := make([]byte, size)

	if _, err := rand.Read(data); err != nil {
		return nil, err
	}

	return data, nil
}

func isEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
