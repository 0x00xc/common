package sqlutil

import (
	"testing"
	"time"
)

func TestQueryBuilder(t *testing.T) {
	q := Query("example").
		Select("id,name").
		Where("created_at >= ? AND created_at < ?", time.Now().Add(-1*time.Hour), time.Now()).
		Or(
			Where("id > ?", 1).Where("id < ?", 99),
		).
		Order("created_at DESC").
		Limit(10).
		Offset(100)
	st, val := q.Build()
	for _, v := range val {
		t.Log(v)
	}
	t.Log(st)
}

func BenchmarkQueryBuilder(b *testing.B) {
	q := Query("example").
		Select("id,name").
		Where("created_at >= ? AND created_at < ?", time.Now().Add(-1*time.Hour), time.Now()).
		Or(
			Where("id > ?", 1).
				Where("id < ?", 99),
		).
		Order("created_at DESC").
		Limit(10).
		Offset(100)
	for i := 0; i < b.N; i++ {
		q.Build()
	}
}
