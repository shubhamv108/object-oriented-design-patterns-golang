package evictionstrategies

type LRU struct {
}

func (L *LRU) Evict() {
	println("Evicting LRU")
}

func (L *LRU) Access(key string) {
	println("Accessing LRU")
}

func (L *LRU) Remove(key string) {
	println("Removing LRU")
}