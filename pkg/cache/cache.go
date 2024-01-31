package cache

import (
	"time"

	"in-memory-cache/internal/core"
)

type Cache struct {
	core *core.CacheCore
}

func New() *Cache {
	return &Cache{
		core: core.NewCore(),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.core.Set(key, value, ttl)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.core.Get(key)
}

func (c *Cache) Delete(key string) {
	c.core.Delete(key)
}
