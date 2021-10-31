package evictionstrategies

type LFU struct {
}

func (L *LFU) Evict() {
	println("Evicting LFU")
}

func (L *LFU) Access(key string) {
	println("Accessing LFU")
}

func (L *LFU) Remove(key string) {
	println("Removing LFU")
}