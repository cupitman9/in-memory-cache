# In-Memory Cache

Simple in-memory cache library in Go, supporting time-to-live (TTL) for cache entries.

## Features

- Set cache entries with a specified TTL (time-to-live).
- Retrieve cache entries.
- Delete cache entries.

## Usage

To use the cache, import the package and create a cache instance:

```go
import (
    "in-memory-cache/pkg/cache"
    "time"
)

func main() {
    c := cache.New()

    // Set a cache entry with TTL
    c.Set("key", "value", 10*time.Second)

    // Get a cache entry
    if value, found := c.Get("key"); found {
        fmt.Println("Found:", value)
    }

    // Delete a cache entry
    c.Delete("key")
}
