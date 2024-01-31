package test

import (
	"testing"
	"time"

	"in-memory-cache/pkg/cache"
)

func TestCacheSetGet(t *testing.T) {
	c := cache.New()
	key := "key1"
	value := "value1"
	ttl := 5 * time.Second

	c.Set(key, value, ttl)

	got, found := c.Get(key)
	if !found || got != value {
		t.Errorf("expected %v, got %v", value, got)
	}

	time.Sleep(6 * time.Second)
	_, found = c.Get(key)
	if found {
		t.Errorf("expected value to be expired")
	}
}

func TestCacheDelete(t *testing.T) {
	c := cache.New()
	key := "key2"
	value := "value2"

	c.Set(key, value, 1*time.Minute)
	c.Delete(key)

	_, found := c.Get(key)
	if found {
		t.Errorf("expected value to be deleted")
	}
}
