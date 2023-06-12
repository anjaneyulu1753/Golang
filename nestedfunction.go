package main

import "fmt"

func greet(name string) {
	displayName := func() {
		fmt.Println("Hi ", name)
	}

	displayName()
}
func main() {
	greet("Anji")
}
