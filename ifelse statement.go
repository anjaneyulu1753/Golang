package main

import "fmt"

func main() {
	number := 100
	number1 := 100
	if number == number1 {
		fmt.Println("GIVEN TWO NUMBERS ARE EQUAL")
	} else if number > number1 {
		fmt.Println("The first number is greaterthan second number")
	} else {
		fmt.Println("The first number is lessthan second number")
	}
}
