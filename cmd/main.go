package main

import (
	"fmt"
	"log"
	"net"

	"github.com/BrendhaCasaro/go-vault/internal/action"
	"github.com/BrendhaCasaro/go-vault/internal/cache"
)

func notifyError(conn net.Conn, err error) {
	conn.Write([]byte(err.Error() + "\n"))
	log.Println(err.Error())
}

func main() {
	lru := cache.NewLRUCache(8)

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go func() {
			for {
				act, err := action.ActionFromReader(conn)
				if err != nil {
					notifyError(conn, err)
					continue
				}
				if act == nil {
					continue
				}

				message, err := action.ExecuteAction(act, lru)
				if err != nil {
					notifyError(conn, err)
					continue
				}

				conn.Write([]byte(message + "\n"))
			}
		}()
	}
}
