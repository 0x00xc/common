package cache

import "encoding/json"

type Cache interface {
	Get(key interface{}) (interface{}, error)
	Put(key, val interface{}) error
}

type BCache interface {
	Get(key string) ([]byte, error)
	Put(key string, val []byte) error
	Del(key string) error
}

type JSONCache interface {
	GetJSON(key string, v interface{}) error
	PutJSON(key string, v interface{}) error
	Del(key string) error
}

type jsonCache struct {
	BCache
}

func (c *jsonCache) GetJSON(key string, v interface{}) error {
	if b, err := c.Get(key); err != nil {
		return err
	} else {
		return json.Unmarshal(b, v)
	}
}

func (c *jsonCache) PutJSON(key string, v interface{}) error {
	if b, err := json.Marshal(v); err != nil {
		return err
	} else {
		return c.Put(key, b)
	}
}

func NewJSONCache(c BCache) JSONCache {
	return &jsonCache{BCache: c}
}
