package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)
	slicenum := array[1:5]
	fmt.Println(slicenum)

}
