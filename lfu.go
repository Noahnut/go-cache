package gocache

import "sync"

type frequencyNode struct {
	frequency int
	next      *frequencyNode
	prev      *frequencyNode
	items     map[interface{}]*lfuCacheItem
}

type lfuCacheItem struct {
	key        interface{}
	value      interface{}
	parentFreq *frequencyNode
}

type lfu struct {
	mu        sync.RWMutex
	keyMap    map[interface{}]*lfuCacheItem
	header    *frequencyNode
	limitSize int
}

func NewLFU(defaultSize int) ICache {
	return &lfu{
		keyMap:    make(map[interface{}]*lfuCacheItem, defaultSize),
		header:    &frequencyNode{},
		limitSize: defaultSize,
	}
}

func (t *lfu) Get(key interface{}) (interface{}, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	item, exist := t.get(key)

	if !exist {
		return nil, false
	}

	return item.value, true
}

func (t *lfu) Set(key interface{}, value interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	item, exist := t.keyMap[key]

	if exist {
		item.value = value
		return
	}

	if len(t.keyMap) >= t.limitSize {
		// evite
		eviteItems := t.header.next.items

		for k := range eviteItems {
			delete(eviteItems, k)
			delete(t.keyMap, k)
			break
		}
	}

	if t.header.next == nil || t.header.next.frequency != 1 {
		newNext := &frequencyNode{
			frequency: 1,
			items:     make(map[interface{}]*lfuCacheItem),
		}

		newNext.next = t.header.next
		newNext.prev = t.header
		t.header.next = newNext
	}

	item = &lfuCacheItem{
		key:        key,
		value:      value,
		parentFreq: t.header.next,
	}

	t.header.next.items[key] = item

	t.keyMap[key] = item
}

func (t *lfu) Delete(key interface{}) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	item, exist := t.keyMap[key]

	if !exist {
		return false
	}

	delete(item.parentFreq.items, key)
	delete(t.keyMap, key)

	if len(item.parentFreq.items) == 0 {
		prevFreqNode := item.parentFreq.prev
		nextFreqNode := item.parentFreq.next

		if prevFreqNode != nil {
			prevFreqNode.next = nextFreqNode
		}

		if nextFreqNode != nil {
			nextFreqNode.prev = prevFreqNode
		}
	}

	return true
}

func (t *lfu) Contains(key interface{}) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	_, exist := t.get(key)
	return exist
}

func (t *lfu) Resize(size int) {
	t.limitSize = size
}

func (t *lfu) Clean() {
	t.keyMap = make(map[interface{}]*lfuCacheItem, t.limitSize)
	t.header = &frequencyNode{}
}

func (t *lfu) get(key interface{}) (*lfuCacheItem, bool) {
	item, exist := t.keyMap[key]

	if !exist {
		return nil, false
	}

	parentFreqNode := item.parentFreq
	nextFreqNode := parentFreqNode.next

	if nextFreqNode == nil || nextFreqNode.frequency != parentFreqNode.frequency+1 {
		nextFreqNode = t.createNewFreqNode(parentFreqNode, nextFreqNode)
		nextFreqNode = parentFreqNode.next
		parentFreqNode.next = nextFreqNode
		nextFreqNode.prev = parentFreqNode
	}

	delete(parentFreqNode.items, key)
	item.parentFreq = nextFreqNode
	nextFreqNode.items[key] = item

	if len(parentFreqNode.items) == 0 {
		parentFreqNode.prev.next = parentFreqNode.next
		parentFreqNode.next.prev = parentFreqNode.prev
	}

	return item, true
}

func (t *lfu) createNewFreqNode(prevFreqNode, nextFreqNode *frequencyNode) *frequencyNode {
	newFreqNode := &frequencyNode{
		frequency: prevFreqNode.frequency + 1,
		next:      nextFreqNode,
		prev:      prevFreqNode,
		items:     make(map[interface{}]*lfuCacheItem),
	}

	if nextFreqNode != nil {
		prevFreqNode.next = newFreqNode
		nextFreqNode.prev = newFreqNode
	} else {
		prevFreqNode.next = newFreqNode
	}

	return newFreqNode
}
