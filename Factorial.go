package main

import "fmt"

func main() {

	var num int
	fmt.Print("Please enter any positivi number: ")
	fmt.Scan(&num)
	var fact int = 1
	for i := 1; i >= num; i++ {
		temp := fact * i
		fact = fact + temp
	}
	fmt.Printf("The facrorial value is: %d", fact)
}
