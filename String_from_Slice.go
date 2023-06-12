package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Team", "How", "are", "You", "All"}
	//fmt.Println(words)
	message := strings.Join(words, "-")
	fmt.Println(message)
}
