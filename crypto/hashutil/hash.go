package hashutil

import (
	"encoding/hex"
	"hash"
)

func Sum(h hash.Hash, data []byte) string {
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
