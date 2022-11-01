package aesutil

import (
	"encoding/hex"
	"testing"
)

func TestAES(t *testing.T) {
	key := []byte("sixteenletterkey")
	cipherText, err := WithPadding(RandomPadding).Encrypt([]byte("hello world!"), key)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(cipherText))

	data, err := WithPadding(RandomPadding).Decrypt(cipherText, key)
	if err != nil {
		t.Error(err)
	}
	if string(data) != "hello world!" {
		t.Fail()
	}

}
