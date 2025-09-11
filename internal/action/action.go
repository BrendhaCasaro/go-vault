package action

import (
	"bytes"
	"fmt"
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
	index := bytes.Index(content, []byte("\r\n"))

	fmt.Print(index)

	return nil, nil
}
