package rsautil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

func GenerateKey(bits int, publicKeyWriter, privateKeyWriter io.Writer) error {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(private)

	privateBlock := &pem.Block{Type: "RSA Private Key", Bytes: x509PrivateKey}
	if err := pem.Encode(privateKeyWriter, privateBlock); err != nil {
		return err
	}

	public := private.Public()
	x509PublicKey, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{Type: "RSA Public Key", Bytes: x509PublicKey}
	if err := pem.Encode(publicKeyWriter, publicBlock); err != nil {
		return err
	}
	return nil
}
