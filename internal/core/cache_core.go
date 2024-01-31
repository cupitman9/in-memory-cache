package core

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value    interface{}
	ExpireAt time.Time
}

type CacheCore struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

func NewCore() *CacheCore {
	return &CacheCore{
		items: make(map[string]CacheItem),
	}
}

func (c *CacheCore) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	expireArt := time.Now().Add(ttl)
	c.items[key] = CacheItem{
		Value:    value,
		ExpireAt: expireArt,
	}
}

func (c *CacheCore) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().After(item.ExpireAt) {
		return nil, false
	}
	return item.Value, true
}

func (c *CacheCore) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}
