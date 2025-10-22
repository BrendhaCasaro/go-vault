package main

import (
	"fmt"
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/action"
	"github.com/BrendhaCasaro/go-vault/internal/cache"
)

func main() {
	lru := cache.NewLRUCache(5)

	r := strings.NewReader("PUT my_key blabla blabla2\r\n")

	ac, err := action.ActionFromReader(r)
	if err != nil {
		fmt.Printf("%s", err)
	}

	action.ExecuteAction(ac, lru)
}
