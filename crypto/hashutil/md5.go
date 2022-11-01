package hashutil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	b := md5.Sum(data)
	return hex.EncodeToString(b[:])
}

func MD5s(s string) string {
	return MD5([]byte(s))
}
