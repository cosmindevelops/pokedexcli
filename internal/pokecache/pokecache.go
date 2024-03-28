package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux	*sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}

	// we need to run this method in a goroutine. we can't run reapLoop in the main goroutine, 
	//we need to run it in a goroutine bechause we will be sitting here running the reapLoop on the main thread for ever and newCache will never returned
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// We use a mutex here because we don't want to have multiple goroutines at the same time adding to the cache
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache{
		if v.createdAt.Before(timeAgo){
			delete(c.cache, k)
		}
	}
}