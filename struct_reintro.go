package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}
	type rectangle struct {
		length  int
		breadth int
	}

	rect := rectangle{24, 16}
	fmt.Println("length of given rectangle is : ", rect.length)
	fmt.Println("breadth of given rectangle is : ", rect.breadth)
	area_of_rectangle := rect.length * rect.breadth
	fmt.Println("the area of given rectangle is: ", area_of_rectangle)

	person1 := person{"anji", 25}
	person2 := person{
		name: "ram",
		age:  25,
	}
	fmt.Println(person1)
	fmt.Print(person2)
}
