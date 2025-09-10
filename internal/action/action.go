package action

import (
	"io"
)

type Actions int

const (
	Get Actions = iota
	Put
	Delete
)

type Action struct {
	Type Actions
	Args []string
}

func ActionFromReader(reader io.Reader) (*Action, error) {
	parse, err := parseAction(reader)
	if err != nil {
		return nil, err
	}

	return parse, nil
}

func parseAction(content string) (*Action, error) {
	return nil, nil
}
