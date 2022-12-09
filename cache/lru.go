package cache

import (
	"container/list"
	"sync"
	"time"
)

type element struct {
	e        *list.Element
	data     interface{}
	expireAt int64
}

type LRUCache struct {
	keys     *list.List
	OnRemove func(key, val interface{})
	n        int

	lock *sync.Mutex
	data map[interface{}]*element
}

func NewLRUCache(maxLength int) *LRUCache {
	return &LRUCache{
		keys: list.New(),
		n:    maxLength,
		lock: new(sync.Mutex),
		data: make(map[interface{}]*element),
	}
}

func (c *LRUCache) Len() int {
	return c.keys.Len()
}

func (c *LRUCache) Put(key, val interface{}, expire ...time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.keys.Len() >= c.n {
		oldest := c.keys.Front()
		if c.OnRemove != nil {
			if old, ok := c.data[oldest.Value]; ok {
				c.OnRemove(oldest.Value, old.data)
			}
		}
		delete(c.data, oldest.Value)
		c.keys.Remove(oldest)
	}
	e := c.keys.PushBack(key)
	ele := &element{e: e, data: val}
	if len(expire) > 0 {
		ele.expireAt = time.Now().Add(expire[0]).Unix()
	}
	c.data[key] = ele
}

func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	data, ok := c.data[key]
	if !ok {
		return nil, ok
	}

	if data.expireAt > 0 && data.expireAt < time.Now().Unix() {
		delete(c.data, key)
		c.keys.Remove(data.e)
		return nil, false
	}

	c.keys.MoveToBack(data.e)
	return data.data, true
}

func (c *LRUCache) Del(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	data, ok := c.data[key]
	if !ok {
		return
	}
	delete(c.data, key)
	c.keys.Remove(data.e)
}
