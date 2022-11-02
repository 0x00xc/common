package smutil

import (
	"crypto/rand"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

func GenerateKey() ([]byte, []byte, error) {
	key, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	priv, err := EncodePrivateKey(key)
	pub := sm2.Compress(&key.PublicKey)
	return priv, pub, err
}

func SM2Encrypt(text []byte, publicKey []byte) ([]byte, error) {
	pk := sm2.Decompress(publicKey)
	return sm2.EncryptAsn1(pk, text, rand.Reader)
}

func SM2Decrypt(cipher []byte, privateKey []byte) ([]byte, error) {
	pk, err := DecodePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return sm2.DecryptAsn1(pk, cipher)
}

func DecodePrivateKey(priByte []byte) (*sm2.PrivateKey, error) {
	c := sm2.P256Sm2()
	k := new(big.Int).SetBytes(priByte)
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}

func EncodePrivateKey(key *sm2.PrivateKey) ([]byte, error) {
	return key.D.Bytes(), nil
}
