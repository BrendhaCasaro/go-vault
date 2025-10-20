package cache

type LRUCache struct {
	capacity int
	items    map[string]*node
	left     *node
	right    *node
}

type node struct {
	key   string
	value string
	next  *node
	prev  *node
}

func newNode(key string, value string) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*node),
		left:     newNode("", ""),
	}
}

func (lru *LRUCache) get(key string) string {
	if value, ok := lru.items[key]; ok {
		return value.value
	}
}
