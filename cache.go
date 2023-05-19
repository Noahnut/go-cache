package gocache

type CachePolicy int

const (
	DefaultCacheSize = 1024
)

const (
	LFU_POLICY CachePolicy = iota + 1
	TINY_LFU_POLICY
)

type ICache interface {
	Get(key interface{}) (value interface{}, exist bool)
	Set(key interface{}, value interface{})
	Delete(key interface{}) (deleted bool)
	Contains(key interface{}) (exist bool)
	Resize(size int)
	Clean()
}

func NewCache(policy ...CachePolicy) ICache {
	if len(policy) == 0 {
		return &lfu{}
	}

	switch policy[0] {
	case LFU_POLICY:
		return NewLFU(DefaultCacheSize)
	case TINY_LFU_POLICY:
		return &tinyLFU{}
	}

	return &lfu{}
}
