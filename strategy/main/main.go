package main

import (
	"object-oriented-design-patterns-golang/strategy/cache"
	"object-oriented-design-patterns-golang/strategy/evictionstrategies"
)

func main() {
	lRUCache := cache.NewCache(&evictionstrategies.LRU{}, 2);
	lRUCache.Put("A", "1");
	lRUCache.Put("B", "1");
	lRUCache.Put("C", "1");

	lFUCache := cache.NewCache(&evictionstrategies.LFU{}, 2);
	lFUCache.Put("A", "1");
	lFUCache.Put("B", "1");
	lFUCache.Put("C", "1");

	fIFOCache := cache.NewCache(&evictionstrategies.FIFO{}, 2);
	fIFOCache.Put("A", "1");
	fIFOCache.Put("B", "1");
	fIFOCache.Put("C", "1");

}
