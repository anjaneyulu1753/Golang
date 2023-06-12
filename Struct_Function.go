package main

import "fmt"

type Rectangle func(int, int) int

// Create a structure
type rectanglepara struct {
	length  int
	breadth int
	color   string

	rect Rectangle
}

func main() {
	result := rectanglepara{
		length:  24,
		breadth: 14,
		color:   "Red",
		rect:    func(length int, breadth int) int { return length * breadth },
	}
	fmt.Println("color of rectangel is: ", result.color)
	fmt.Println("Area of rectangle is: ", result.rect(result.length, result.breadth))
}
