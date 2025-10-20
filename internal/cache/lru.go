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
		right:    newNode("", ""),
	}
}

func (lru *LRUCache) insert(node *node) {}
func (lru *LRUCache) remove(node *node) {
	prev, next := node.prev, node.next
}

func (lru *LRUCache) get(key string) string {
	if node, ok := lru.items[key]; ok {
		lru.insert(node)
		return node.value
	}

	return ""
}

func (lru *LRUCache) put(key string, value string) {
	if node, ok := lru.items[key]; ok {
		lru.remove(node)
	}

	node := newNode(key, value)
	lru.items[key] = node
	lru.insert(node)

	if len(lru.items) == lru.capacity {
		least := lru.left
		lru.remove(least)
		delete(lru.items, least.key)
	}
}
