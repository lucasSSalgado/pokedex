package pokecache

import (
	"time"
)

type Cache struct {
	cacheList map[string] CacheEntry
}

type CacheEntry struct {
	data    []byte
	creatAt time.Time
}

func NewCache() Cache {
	cache := Cache{
		cacheList: make(map[string]CacheEntry),
	}
	interval := time.Minute * 5
	go cache.removeLoop(interval)
	return cache
}

func (c *Cache) Add(key string, data []byte) {

	entry := CacheEntry {
		data: data,
		creatAt: time.Now().UTC(),
	}
	
	c.cacheList[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	
	val, ok := c.cacheList[key]
	return val.data, ok
}

func (c *Cache) remove(interval time.Duration) {
	
	cacheValues := c.cacheList	
	timeAgo := time.Now().UTC().Add(-interval)

	for k, val := range cacheValues {
		if val.creatAt.Before(timeAgo) {
			delete(cacheValues, k)
		}		 
	}
}

func (c *Cache) removeLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.remove(interval)
	}
}