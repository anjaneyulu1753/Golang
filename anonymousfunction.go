package main

import "fmt"

func main() {
	var greet = func() {
		fmt.Println("hello how are you?")
	}
	greet()
	sum := func(n1 int, n2 int) {
		total := n1 + n2
		fmt.Println("The sum of given two numbers is: ", total)
	}
	sum(16, 18)

}
