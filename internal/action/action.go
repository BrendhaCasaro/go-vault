package action

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

type ActionType int

const (
	NULL ActionType = iota
	GET
	PUT
	DELETE
)

type Action struct {
	Type ActionType
	Args []string
}

func ActionFromReader(reader io.Reader) (*Action, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanCRLF)
	scanner.Scan()

	return parseAction(scanner.Bytes())
}

func parseAction(content []byte) (*Action, error) {
	comandFields := bytes.Split(content, []byte(" "))

	actionType := parseActionType(string(comandFields[0]))
	if actionType == NULL {
		return nil, errors.New("Invalid action")
	}

	args := make([]string, len(comandFields[1:]))
	for i, cf := range comandFields[1:] {
		args[i] = string(cf)
	}
	if len(args) == 0 {
		return nil, errors.New("Empty arguments")
	}

	return &Action{
		actionType,
		args,
	}, nil
}

func parseActionType(actionType string) ActionType {
	switch actionType {
	case "GET":
		return GET
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	default:
		return NULL
	}
}
