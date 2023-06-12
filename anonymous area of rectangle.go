package main

import "fmt"

func main() {
	area := func(length, breadth int) int {
		return length * breadth
	}
	fmt.Println("The area of the rectangle is: ", area(20, 23))
}
