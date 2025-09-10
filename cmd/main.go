package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(text string) {
	text = strings.TrimSpace(text)
	fmt.Printf("Texto digitado: %v\n\n", text)
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Digite algo:\n")
		inputText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Have a error to get the input: ", err)
			return
		}

		Parse(inputText)
	}
}
