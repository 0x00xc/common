package cry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid cipher text")
	}
	unPadding := int(data[length-1])
	if unPadding > length {
		return nil, errors.New("invalid cipher text")
	}
	return data[:(length - unPadding)], nil
}

func padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	var padText []byte
	if padding > 1 {
		padText = randBytes(padding - 1)
	}
	padText = append(padText, byte(padding))
	return append(data, padText...)
}

func unpadding(data []byte) ([]byte, error) {
	return pkcs7UnPadding(data)
}

//AesEncrypt 加密
func AESEncrypt(data []byte, key, iv []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))

	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//AesDecrypt 解密
func AESDecrypt(data []byte, key, iv []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	//blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = unpadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func AESEncryptIV(data []byte, key []byte) ([]byte, error) {
	iv := randBytes(16)
	res, err := AESEncrypt(data, key, iv)
	return append(res, iv...), err
}

func AESDecryptIV(data []byte, key []byte) ([]byte, error) {
	if len(data) < 16 {
		return nil, errors.New("invalid cipher text")
	}
	iv := data[len(data)-16:]
	return AESDecrypt(data[:len(data)-16], key, iv)
}
