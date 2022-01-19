package cache

import "testing"

func TestLruCache(t *testing.T) {
	cache := NewLRUCache(5)
	for i := 0; i < 6; i++ {
		cache.Put(i, i)
	}

	cache.Get(1)
	cache.Put(6, 6)

	for i := 0; i < 7; i++ {
		t.Log(cache.Get(i))
	}

}
