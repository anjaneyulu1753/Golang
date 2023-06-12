package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("The name and age are : ")
	fmt.Scanln(&name, &age)
	fmt.Println("given name is: ", name)
	fmt.Println("hello, ")
	fmt.Println(age)
	fmt.Print("WORLD")
}
