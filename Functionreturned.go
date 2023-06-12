package main

import "fmt"

func greet() func() {
	return func() {
		fmt.Println("Hi Anji")
	}
}
func main() {
	g1 := greet()

	g1()
}
