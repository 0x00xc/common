package rsautil

import (
	"bytes"
	"strings"
	"testing"
)

func TestRSA(t *testing.T) {
	publicKey := bytes.NewBuffer([]byte{})
	privateKey := bytes.NewBuffer([]byte{})

	if err := GenerateKey(2048, publicKey, privateKey); err != nil {
		t.Error(err)
	}
	text := "hello!"

	cipherText, err := Encrypt([]byte(text), publicKey.Bytes())
	if err != nil {
		t.Error(err)
	}

	plainText, err := Decrypt(cipherText, privateKey.Bytes())
	if err != nil {
		t.Error(err)
	}
	if string(plainText) != text {
		t.Fail()
	}

}

func TestRSALong(t *testing.T) {
	publicKey := bytes.NewBuffer([]byte{})
	privateKey := bytes.NewBuffer([]byte{})

	if err := GenerateKey(1024, publicKey, privateKey); err != nil {
		t.Error(err)
	}

	text := strings.Repeat("hello ", 128)
	buf := bytes.NewBuffer([]byte{})
	if err := EncryptLong(strings.NewReader(text), buf, publicKey.Bytes()); err != nil {
		t.Error(err)
	}
	out := bytes.NewBuffer([]byte{})
	if err := DecryptLong(buf, out, privateKey.Bytes()); err != nil {
		t.Error(err)
	}
	if string(out.Bytes()) != text {
		t.Fatal()
	}
}

func BenchmarkRSAEncrypt(b *testing.B) {
	publicKey := bytes.NewBuffer([]byte{})
	privateKey := bytes.NewBuffer([]byte{})

	if err := GenerateKey(2048, publicKey, privateKey); err != nil {
		b.Error(err)
	}
	text := strings.Repeat("hello", 24)

	for i := 0; i < b.N; i++ {
		_, _ = Encrypt([]byte(text), publicKey.Bytes())
	}
}

func BenchmarkRSADecrypt(b *testing.B) {
	publicKey := bytes.NewBuffer([]byte{})
	privateKey := bytes.NewBuffer([]byte{})

	if err := GenerateKey(2048, publicKey, privateKey); err != nil {
		b.Error(err)
	}
	text := strings.Repeat("hello", 24)
	cipher, _ := Encrypt([]byte(text), publicKey.Bytes())
	for i := 0; i < b.N; i++ {
		_, _ = Decrypt(cipher, privateKey.Bytes())
	}
}
