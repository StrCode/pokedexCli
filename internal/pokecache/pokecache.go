package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	c.cache[key] = cacheEntry{
		time.Now(),
		val,
	}
	defer c.m.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	entry, exists := c.cache[key]
	defer c.m.Unlock()
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Current time:", t)
			c.m.Lock()
			keys := []string{}
			for i, cache := range c.cache {
				t1 := cache.createdAt
				fmt.Println("cache time:", t1)
				if t.Sub(t1) > interval {
					keys = append(keys, i)
				}
			}

			for _, key := range keys {
				delete(c.cache, key)
			}
			c.m.Unlock()
		}
	}
}
