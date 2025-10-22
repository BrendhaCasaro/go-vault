package action

import (
	"errors"
	"strings"

	"github.com/BrendhaCasaro/go-vault/internal/cache"
)

func ExecuteAction(comand *Action, c cache.Cache) (string, error) {
	switch comand.Type {
	case GET:
		if len(comand.Args) < 1 {
			return "", errors.New("Not enough arguments")
		}
		key := comand.Args[0]

		value, ok := c.Get(key)
		if ok {
			return value, nil
		}

		return "", errors.New("Value searched not exist")

	case PUT:
		if len(comand.Args) < 2 {
			return "", errors.New("Not enough arguments")
		}
		key := comand.Args[0]
		value := strings.Join(comand.Args[1:], " ")

		err := c.Put(key, value)
		if err != nil {
			return "", err
		}

		return "Object inserted", nil

	case DELETE:
		if len(comand.Args) < 1 {
			return "", errors.New("Not enough arguments")
		}
		key := comand.Args[0]

		err := c.Delete(key)
		if err != nil {
			return "", err
		}

		return "Object deleted", nil

	default:
		return "", errors.New("Command is not valid")
	}
}
