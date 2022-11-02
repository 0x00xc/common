package sqlutil

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	b := Insert("example").
		Value("id", 1).
		Value("created_at", time.Now()).
		Value("updated_at", time.Now())

	st, val := b.Build()
	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)
}
