package main

import "fmt"

func main() {
	var num int
	var fact = 1
	fmt.Print("Enter the any positive number: ")
	fmt.Scanln(&num)
	if num < 0 {
		fmt.Println("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= num; i++ {
			fact *= i

		}
		fmt.Print("Factorial of given number is :", fact)
	}
}
