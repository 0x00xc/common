package randx

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	CAPITAL    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MINUSCULES = "abcdefghijklmnopqrstuvwxyz"
	NUMBER     = "1234567890"
	LETTER     = MINUSCULES + CAPITAL + NUMBER
	HEX        = "1234567890abcdef"
)

var _rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Intn(n int) int {
	return _rand.Intn(n)
}

func Float() float64 {
	return _rand.Float64()
}

func Bytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(_rand.Intn(256))
	}
	return b
}

func String(n int, seed ...string) string {
	s := LETTER
	if len(seed) > 0 {
		s = seed[0]
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = s[_rand.Intn(len(s))]
	}
	return string(b)
}

func Hex(n int) string {
	return String(n, HEX)
}

func BytesHex(n int) string {
	return hex.EncodeToString(Bytes(n))
}
