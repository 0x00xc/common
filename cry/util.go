package cry

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(r.Intn(256))
	}
	return b
}
