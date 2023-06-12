package main

import "fmt"

func add_numbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {
	result := add_numbers(23, 45)
	fmt.Println("the sum of two numbers is: ", result)
}
