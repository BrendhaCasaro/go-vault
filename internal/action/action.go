package action

import (
	"bytes"
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
	comand, _, _ := bytes.Cut(content, []byte("\r\n"))

	splitedComand := bytes.Split(comand, []byte(" "))

	argsComand := make([]string, len(splitedComand[1:]))

	for i, a := range splitedComand[1:] {
		argsComand[i] = string(a)
	}

	switch string(splitedComand[0]) {
	case "GET":
		return &Action{
			GET,
			argsComand,
		}, nil
	case "PUT":
		return &Action{
			PUT,
			argsComand,
		}, nil
	case "DELETE":
		return &Action{
			DELETE,
			argsComand,
		}, nil
	}

	return nil, nil
}

func typeFromString(actionType string) {
}
