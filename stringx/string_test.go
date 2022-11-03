package stringx

import (
	"strconv"
	"testing"
)

func TestReplaceFunc(t *testing.T) {
	var s = "UPDATE example SET id = ?, name = ?, count = count + ?"
	s = ReplaceFunc(s, "?", func(i int) string {
		return ":" + strconv.Itoa(i+1)
	})
	t.Log(s)

}
