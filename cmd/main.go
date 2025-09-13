package main

import (
	"fmt"
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/action"
)

func main() {
	r := strings.NewReader("DELETE my_key blabla blabla2\r\n")

	action, err := action.ActionFromReader(r)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", action)
}
