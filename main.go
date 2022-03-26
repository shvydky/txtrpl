package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Using: txtrpl <search pattern> <replace with>")
		return
	}
	var (
		pattern = os.Args[1]
		replace = os.Args[2]
	)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	buffer := make([]rune, 0, 256)
	for err == nil {
		if char != '\n' {
			buffer = append(buffer, char)
		} else {
			text := string(buffer)
			text = strings.ReplaceAll(text, pattern, replace)
			fmt.Println(text)
			buffer = buffer[:0]
		}
		char, _, err = reader.ReadRune()
	}
}
