package pokecache //Create a new internal package called pokecache in your internal directory

import (
	"sync"
	"time"
)

// I used a Cache struct to hold a map[string]cacheEntry and a mutex to protect the map across goroutines
type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct { // A cacheEntry should be a struct with two fields:
	createdAt time.Time // A time.Time that represents when the entry was created.
	val       []byte    // A []byte that represents the raw data we're caching.
}

func NewCache(interval time.Duration) Cache { //You'll probably want to expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
