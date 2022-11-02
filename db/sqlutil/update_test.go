package sqlutil

import (
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	b := Update("example").
		Set("`name`", "hello").
		Set("created_at = ?", time.Now()).
		Where("id", 1)

	st, val := b.Build()

	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)

}
