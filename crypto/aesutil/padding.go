package aesutil

import (
	"bytes"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Padding interface {
	Padding(ciphertext []byte, blockSize int) []byte
	Unpadding(data []byte) ([]byte, error)
}

type pkcs7Padding struct{}

func (p *pkcs7Padding) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == 0 {
		padding = blockSize
	}
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (p *pkcs7Padding) Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if length-unpadding >= len(data) || length-unpadding <= 0 {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

type randomPadding struct {
	seed *rand.Rand
}

func (p *randomPadding) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == 0 {
		padding = blockSize
	} else {
		padding += blockSize
	}

	temp := make([]byte, padding)
	for i := 0; i < padding-1; i++ {
		temp[i] = byte(p.seed.Intn(256))
	}
	temp[len(temp)-1] = byte(padding)
	return append(ciphertext, temp...)
}

func (p *randomPadding) Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if length-unpadding >= len(data) || length-unpadding <= 0 {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

var (
	PKCS7Padding = new(pkcs7Padding)
	//RandomPadding 使用随机数填充，可以使每次生成的密文都不一样
	RandomPadding = &randomPadding{seed: rand.New(rand.NewSource(time.Now().UnixNano()))}
)
