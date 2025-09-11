package main

import (
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/action"
)

func main() {
	r := strings.NewReader("blabla \r\n")

	action.ActionFromReader(r)
}
