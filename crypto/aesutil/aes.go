package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
)

// Encrypt
//
//	 params
//		data 原文数据
//		key  密钥
//		ivs  自定义向量，默认取密钥
//	 return
//		密文
//		错误信息
func Encrypt(data, key []byte, ivs ...[]byte) ([]byte, error) {
	return WithPadding(PKCS7Padding).Encrypt(data, key, ivs...)
}

// Decrypt
//
//	 params
//		cipherText 密文
//		key        密钥
//		ivs        自定义向量，默认取密钥
//	 return
//		原文
//		错误信息
func Decrypt(cipherText, key []byte, ivs ...[]byte) ([]byte, error) {
	return WithPadding(PKCS7Padding).Decrypt(cipherText, key, ivs...)
}

func WithPadding(padding Padding) Crypto {
	return &crypto{padding: padding}
}

type crypto struct {
	padding Padding
}

type Crypto interface {
	Encrypt(data, key []byte, ivs ...[]byte) ([]byte, error)
	Decrypt(cipherText, key []byte, ivs ...[]byte) ([]byte, error)
}

func (c *crypto) Encrypt(data, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data = c.padding.Padding(data, blockSize)
	var iv []byte
	if len(ivs) > 0 && len(ivs[0]) >= blockSize {
		iv = ivs[0][:blockSize]
	} else {
		iv = key[:blockSize]
	}
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	cipherText := make([]byte, len(data))
	blockMode.CryptBlocks(cipherText, data)
	return cipherText, nil
}

func (c *crypto) Decrypt(cipherText, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	var iv []byte
	if len(ivs) > 0 && len(ivs[0]) >= blockSize {
		iv = ivs[0][:blockSize]
	} else {
		iv = key[:blockSize]
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	return c.padding.Unpadding(origData)
}
