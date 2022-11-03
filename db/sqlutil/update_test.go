package sqlutil

import (
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	b := Update("example").
		SetColumn("`name`", "hello").
		SetColumn("created_at = ?", time.Now()).
		Set(
			KV("count = count + ?", 1),
		).
		Where("id", 1)

	st, val := b.Build()

	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)

	w := Wrapper{Placeholder: DMPlaceholder, Builder: b}
	st, val = w.Build()
	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)

}
