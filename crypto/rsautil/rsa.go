package rsautil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

// Encrypt 加密
func Encrypt(plainText []byte, public []byte) ([]byte, error) {
	block, _ := pem.Decode(public)
	if block == nil {
		return nil, errors.New("invalid public key")
	}
	pkInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk := pkInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pk, plainText)
	return cipherText, err

}

// Decrypt 解密
func Decrypt(cipherText []byte, private []byte) ([]byte, error) {
	block, _ := pem.Decode(private)
	if block == nil {
		return nil, errors.New("invalid private key")
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, pk, cipherText)
	return plainText, err
}

// EncryptLong 长数据加密
// RSA算法对原文长度有限制，长数据加密只能分段加密
func EncryptLong(in io.Reader, out io.Writer, public []byte) error {
	block, _ := pem.Decode(public)
	if block == nil {
		return errors.New("invalid public key")
	}
	pkInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pk := pkInterface.(*rsa.PublicKey)
	for {
		//RSA算法本身要求加密内容也就是明文长度m必须0<m<密钥长度n。如果小于这个长度就需要进行padding，因为如果没有padding，
		//就无法确定解密后内容的真实长度，字符串之类的内容问题还不大，以0作为结束符，但对二进制数据就很难，因为不确定后面的0是内
		//容还是内容结束符。而只要用到padding，那么就要占用实际的明文长度，于是实际明文长度需要减去padding字节长度。我们一般
		//使用的padding标准有NoPPadding、OAEPPadding、PKCS1Padding等，其中PKCS#1建议的padding就占用了11个字节。
		var buf = make([]byte, pk.Size()-11)
		n, err := in.Read(buf)
		if err != io.EOF && err != nil {
			return err
		}
		if n == 0 {
			break
		}
		part, err := rsa.EncryptPKCS1v15(rand.Reader, pk, buf[:n])
		if err != nil {
			return err
		}
		if _, err := out.Write(part); err != nil {
			return err
		}
	}
	return nil
}

// DecryptLong 长数据解密
func DecryptLong(in io.Reader, out io.Writer, private []byte) error {
	block, _ := pem.Decode(private)
	if block == nil {
		return errors.New("invalid private key")
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	for {
		var buf = make([]byte, pk.Size())
		n, err := in.Read(buf)
		if err != io.EOF && err != nil {
			return err
		}
		if n == 0 {
			break
		}
		part, err := rsa.DecryptPKCS1v15(rand.Reader, pk, buf[:n])
		if err != nil {
			return err
		}
		if _, err := out.Write(part); err != nil {
			return err
		}
	}
	return nil
}
