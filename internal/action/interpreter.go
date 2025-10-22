package action

import (
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/cache"
)

func ExecuteAction(comand *Action, lru *cache.LRUCache) {
	switch comand.Type {
	case 1:
		key := comand.Args[0]
		lru.Get(key)

	case 2:
		key := comand.Args[0]
		value := strings.Join(comand.Args[1:], " ")
		// fmt.Printf("%s", value)

		lru.Put(key, value)

	case 3:
		key := comand.Args[0]
		lru.Delete(key)
	}
}
