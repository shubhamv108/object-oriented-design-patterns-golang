package cache

import "object-oriented-design-patterns-golang/strategy/evictionstrategies"

type Cache struct {
	storage          map[string]string
	evictionStrategy evictionstrategies.EvictionStrategy
	maxCapacity      int
}

func NewCache(evictionStrategy evictionstrategies.EvictionStrategy, maxCapacity int) *Cache {
	return &Cache{
		storage: make(map[string]string),
		evictionStrategy: evictionStrategy,
		maxCapacity: maxCapacity}
}

func (c *Cache) Put(key, value string) {
	if len(c.storage) == c.maxCapacity {
		c.evict()
	}
	c.storage[key] = value;
}

func (c *Cache) evict() {
	c.evictionStrategy.Evict();
}

func (c *Cache) Get(key string) string {
	return c.storage[key];
}
