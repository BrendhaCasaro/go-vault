package action

import (
	"bytes"
	"errors"
	"io"
)

type Actions int

const (
	GET Actions = iota
	PUT
	DELETE
)

type Action struct {
	Type Actions
	Args []string
}

func ActionFromReader(reader io.Reader) (*Action, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	parse, err := parseAction(buf)
	if err != nil {
		return nil, err
	}

	return parse, nil
}

func parseAction(content []byte) (*Action, error) {
	if len(content) == 0 {
		return nil, errors.New("Empty comand")
	}

	index := bytes.Index(content, []byte("\r\n"))
	if index == -1 {
		return nil, errors.New("Missing the CRFL'\r\n'")
	}
	content = content[:index]

	comandFields := bytes.Split(content, []byte(" "))

	actionType := parseActionType(string(comandFields[0]))
	if actionType == -1 {
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

func parseActionType(actionType string) Actions {
	switch actionType {
	case "GET":
		return GET
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	default:
		return -1
	}
}
