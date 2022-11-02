package smutil

import (
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
)

func SM3s(s string) string {
	return hex.EncodeToString(sm3.Sm3Sum([]byte(s)))
}

func SM3(data []byte) string {
	return hex.EncodeToString(sm3.Sm3Sum(data))
}
