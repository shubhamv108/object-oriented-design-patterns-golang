package evictionstrategies

type EvictionStrategy interface {
	Evict();
	Remove(key string);
	Access(key string);
}