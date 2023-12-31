package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
    cache map[string]cacheEntry
    mux *sync.Mutex
}

type cacheEntry struct {
    val []byte
    createdAt time.Time 
}

// new cache creator 
// Also sets up cache referesh
func NewCache(interval time.Duration) Cache {
    c := Cache{
        cache: make(map[string]cacheEntry),
        mux: &sync.Mutex{},
    }
    go c.pluckLoop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mux.Lock()
    c.cache[key] = cacheEntry {
        val: val,
        createdAt: time.Now().UTC(),
    }
    defer c.mux.Unlock()

}

func (c *Cache) Get(key string) ([]byte, bool) {
    cached, ok := c.cache[key] 
    return cached.val, ok
}


// So this basically generalizes a cache referesh
/*
    So what we can do is create a ticker that will always run, and
    say every 5 minutes it will pluck the cache items needed
*/
func (c *Cache) pluckLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.pluck(interval) 
    }
}

// For some series of minutes say 
func (c *Cache) pluck(interval time.Duration) {
    c.mux.Lock()
    duration := time.Now().UTC().Add(-interval)
    for k,v := range c.cache {
        if v.createdAt.Before(duration) {
            delete(c.cache, k)
        }
    }
    defer c.mux.Unlock()
}
