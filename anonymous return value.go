package main

import "fmt"

func main() {
	sum := func(n1 int, n2 int) int {
		return n1 + n2
	}
	result := sum(34, 23)
	fmt.Println("The sum of given two numbers is: ", result)
}
