package cry

import (
	"testing"
)

func TestAES(t *testing.T) {
	key := randBytes(16)
	iv := randBytes(16)
	b1, err := AESEncrypt([]byte("hello"), key, iv)
	if err != nil {
		t.Log(err)
	}
	t.Log(len(b1))
	b2, err := AESDecrypt(b1, key, iv)
	t.Log(len(b2), string(b2))
}

func TestAESIv(t *testing.T) {
	key := randBytes(16)
	b1, err := AESEncryptIV([]byte("hello"), key)
	if err != nil {
		t.Log(err)
	}
	t.Log(len(b1))
	b2, err := AESDecryptIV(b1, key)
	t.Log(err)
	t.Log(len(b2), string(b2))
}
