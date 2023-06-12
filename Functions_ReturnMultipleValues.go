package main

import "fmt"

func addition_subtraction(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}

func main() {
	add, difference := addition_subtraction(34, 21)
	fmt.Println("The sum of two numbers is: ", add)
	fmt.Println("the difference of two numbers is: ", difference)
}
