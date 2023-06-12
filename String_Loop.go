package main

import "fmt"

func main() {
	name := "anjaneyulu"
	for _, character := range name {
		fmt.Printf("%c\n", character)
	}
	sent := "This article is about \"String\" in Go Programming."
	fmt.Println(sent)

}
