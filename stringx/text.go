package stringx

import "math"

func PixelLen(s string) int {
	var n float64
	for _, v := range []rune(s) {
		if v < 256 {
			n += 0.5
		} else {
			n += 1
		}
	}
	return int(math.Ceil(n))
}
