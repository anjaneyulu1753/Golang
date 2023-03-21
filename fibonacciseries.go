package main

import "fmt"

func main() {
	a := 0
	b := 1
	c := b
	var num int
	fmt.Print("Enter any positive number: ")
	fmt.Scanln(&num)
	fmt.Printf("Series is: %d %d ", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf("%d ", b)
	}
}
