package main

import "fmt"

func findSquare(num int) int {
	return num * num
}
func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}
	result := findSquare(sum(9, 6))
	fmt.Println("the result is: ", result)
}
