package lrucache

import (
	"sync"
)

var _ LRUCache = (*InMemoryLRUCache)(nil)

// InMemoryLRUCache - implementation of LRUCache interface
type InMemoryLRUCache struct {
	m  map[string]*cachedItem
	mu sync.Mutex

	queue            *queue
	curSize, maxSize int
}

type queue struct {
	head, tail *cachedItem
}

type cachedItem struct {
	key        string
	value      interface{}
	next, prev *cachedItem
}

// NewLRUCache construct new InMemoryLRUCache instance.
// maxSize - max size of cache. Should be >= 1
func NewLRUCache(maxSize int) *InMemoryLRUCache {
	if maxSize < 1 {
		panic("illegal size of cache provided!")
	}
	m := make(map[string]*cachedItem)
	return &InMemoryLRUCache{m: m, maxSize: maxSize, queue: &queue{}}
}

// Add - LRUCache.Add implementation
func (i *InMemoryLRUCache) Add(key, value string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	item := &cachedItem{
		key:   key,
		value: value,
	}

	oldItem, inMap := i.m[key]
	if inMap {
		i.rmItem(oldItem)
		i.putInHead(item)
	} else {
		if i.curSize == i.maxSize {
			tail := i.rmTail()
			delete(i.m, tail.key)
		} else {
			i.curSize++
		}
		i.putInHead(item)

	}

	i.m[key] = item
	return inMap
}

// Get - LRUCache.Get implementation
func (i *InMemoryLRUCache) Get(key string) (value string, ok bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	item, inMap := i.m[key]
	if inMap {
		i.rmItem(item)
		i.putInHead(item)
		return item.value.(string), inMap
	}
	return "", inMap
}

// Remove - LRUCache.Remove implementation
func (i *InMemoryLRUCache) Remove(key string) (ok bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	item, inMap := i.m[key]
	if inMap {
		i.rmItem(item)
		i.curSize--
	}
	delete(i.m, key)
	return inMap
}

func (i *InMemoryLRUCache) rmTail() *cachedItem {
	tail := i.queue.tail
	if i.queue.head == i.queue.tail {
		i.queue.head = nil
		i.queue.tail = nil
		return tail
	}
	i.queue.tail = tail.prev
	tail.prev = nil
	i.queue.tail.next = nil
	return tail
}

func (i *InMemoryLRUCache) rmItem(item *cachedItem) {
	if i.queue.head == item {
		if i.queue.tail == item {
			i.queue.head = nil
			i.queue.tail = nil
			return
		}
		i.queue.head = i.queue.head.next
		i.queue.head.prev = nil
		return
	}
	if i.queue.tail == item {
		i.queue.tail = item.prev
	} else {
		item.next.prev = item.prev
	}
	item.prev.next = item.next
	item.prev = nil
	item.next = nil
}

func (i *InMemoryLRUCache) putInHead(item *cachedItem) {
	if i.queue.head == nil {
		i.queue.head = item
		i.queue.tail = item
		return
	}
	prevHead := i.queue.head
	item.prev = nil
	item.next = prevHead
	prevHead.prev = item
	i.queue.head = item
}

// Keys - return all keys in cache
func (i *InMemoryLRUCache) Keys() []string {
	var keys []string
	cur := i.queue.head
	for cur != nil {
		keys = append(keys, cur.key)
		cur = cur.next
	}
	return keys
}
