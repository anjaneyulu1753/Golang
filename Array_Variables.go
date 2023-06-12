package main

import "fmt"

func main() {
	var array_of_integer = [...]int{2, 3, 5, 2, 87}
	var array_string = [...]string{"hello", "anji"}
	fmt.Println(array_of_integer)
	for index, item := range array_of_integer {
		fmt.Println(index, ":", item)
	}
	fmt.Println(array_string)

}
