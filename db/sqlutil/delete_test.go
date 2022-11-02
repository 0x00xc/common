package sqlutil

import "testing"

func TestDelete(t *testing.T) {
	b := Delete("example").Where("id = ?", 1)
	st, val := b.Build()
	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)

}
