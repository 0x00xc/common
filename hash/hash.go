package hash

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

func Hash(h hash.Hash, b []byte) []byte {
	h.Write(b)
	return h.Sum(nil)
}

func Hashs(h hash.Hash, b []byte) string {
	return hex.EncodeToString(Hash(h, b))
}

func MD5(b []byte) []byte {
	return Hash(md5.New(), b)
}

func MD5s(b []byte) string {
	return hex.EncodeToString(MD5(b))
}
