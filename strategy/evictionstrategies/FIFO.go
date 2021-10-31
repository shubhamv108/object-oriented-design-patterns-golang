package evictionstrategies

type FIFO struct {
}

func (L *FIFO) Evict() {
	println("Evicting FIFO")
}

func (L *FIFO) Access(key string) {
	println("Accessing FIFO")
}

func (L *FIFO) Remove(key string) {
	println("Removing FIFO")
}
