package main

import (
	"fmt"
	"time"

	"in-memory-cache/pkg/cache"
)

func main() {

	c := cache.New()

	key := "user1"
	value := "Bob K"
	ttl := 5 * time.Second

	c.Set(key, value, ttl)
	fmt.Printf("Set %s = %s with TTL %s\n", key, value, ttl)

	if v, found := c.Get(key); found {
		fmt.Printf("Get %s: %s\n", key, v)
	} else {
		fmt.Printf("Get %s: not found\n", key)
	}

	time.Sleep(6 * time.Second)
	if _, found := c.Get(key); found {
		fmt.Printf("Get %s after TTL: still in cache\n", key)
	} else {
		fmt.Printf("Get %s after TTL: not found\n", key)
	}

	c.Delete(key)
	fmt.Printf("Deleted %s\n", key)

	if _, found := c.Get(key); found {
		fmt.Printf("Get %s after delete: still in cache\n", key)
	} else {
		fmt.Printf("Get %s after delete: not found\n", key)
	}
}
