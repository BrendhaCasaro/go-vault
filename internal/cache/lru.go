package cache

import "errors"

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

func NewNode(key string, value string) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func NewLRUCache(capacity int) Cache {
	lru := &LRUCache{
		capacity: capacity,
		items:    make(map[string]*node),
		left:     NewNode("", ""),
		right:    NewNode("", ""),
	}

	lru.left.next, lru.right.prev = lru.right, lru.left

	return lru
}

func (lru *LRUCache) insert(node *node) {
	l, r := lru.right.prev, lru.right

	r.prev, l.next = node, node
	node.next, node.prev = r, l
}

func (lru *LRUCache) remove(node *node) {
	prev, next := node.prev, node.next

	prev.next, next.prev = next, prev
}

func (lru *LRUCache) Get(key string) (string, bool) {
	if node, ok := lru.items[key]; ok {
		lru.remove(node)
		lru.insert(node)
		return node.value, ok
	}

	return "", false
}

func (lru *LRUCache) Put(key string, value string) error {
	if node, ok := lru.items[key]; ok {
		lru.remove(node)
	}

	node := NewNode(key, value)
	lru.items[key] = node
	lru.insert(node)

	if len(lru.items) > lru.capacity {
		least := lru.left.next
		lru.remove(least)
		delete(lru.items, least.key)
	}

	return nil
}

func (lru *LRUCache) Delete(key string) error {
	if node, ok := lru.items[key]; ok {
		delete(lru.items, key)
		lru.remove(node)
		node = nil
		return nil
	}

	return errors.New("Error: provided key does not exist")
}
