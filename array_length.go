package main

import "fmt"

func main() {
	values := [...]int{2, 4, 7, 4, 3, 6, 4, 7, 9, 3, 8, 9}
	length := len(values)
	index := (values[2 : length-1])
	fmt.Println(index)
	fmt.Println("the length of the given array: ", length)
	for i := 0; i < length; i++ {
		fmt.Println(values[i])
	}
}
