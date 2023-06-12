package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}
	//Create an in instance of struct
	//var person1 person

	person1 := Person{"ram", 25}
	person2 := Person{
		name: "anji",
		age:  25,
	}
	fmt.Println(person1)
	fmt.Println(person2)

}
