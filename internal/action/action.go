package action

import (
	"bytes"
	"fmt"
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
	endComand := bytes.Index(content, []byte("\r\n"))

	splitedComand := bytes.Split(content, []byte(" "))
	_, comand, found := bytes.Cut(content, content[endComand:endComand+3])

	fmt.Printf("%v\n", comand)
	fmt.Printf("%v\n", found)

	fmt.Printf("%v\n", endComand)
	fmt.Printf("%q", splitedComand)

	// switch string(splitedComand[0]) {
	// case "GET":
	// 	return &Action{
	// 		GET,
	// 	}
	// }
	return nil, nil
}
