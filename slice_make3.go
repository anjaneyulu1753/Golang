package main

import "fmt"

func main() {
	new_slice := make([]int, 5, 7)
	old_slice := []int{2, 3, 8}
	new_slice[0] = 34
	new_slice[1] = 23
	new_slice[2] = 76
	new_slice[3] = 87
	new_slice[4] = 67
	//new_slice[5] = 98
	new_slice = append(new_slice, old_slice...)
	fmt.Println(new_slice)

}
