package main

import (
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/action"
)

func main() {
	r := strings.NewReader("Get my_key blabla\r\n")

	action.ActionFromReader(r)
}
