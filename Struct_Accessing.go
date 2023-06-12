package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func main() {

	rect := Rectangle{24, 14}
	fmt.Println(rect)
	area_of_rectangle := rect.length * rect.breadth
	fmt.Println("The area of given rectangle is: ", area_of_rectangle)
}
