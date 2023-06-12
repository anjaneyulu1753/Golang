package main

import "fmt"

func main() {
	num1 := 65
	num2 := 65

	if num1 >= num2 {
		if num1 == num2 {
			fmt.Println("The first number is equal to second number")
		} else {
			fmt.Println("The first number is greaterthan second number")
		}
	} else {
		fmt.Println("The first number is lessthan second number")
	}
}
