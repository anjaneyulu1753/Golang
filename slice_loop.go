package main

import "fmt"

func main() {
	old_slice := []int{}
	new_slice := []int{4, 6, 8, 0, 3, 6, 7, 4, 3, 5, 6, 78}
	length := len(new_slice)
	for i := 0; i < length; i++ {
		old_slice = append(old_slice, new_slice[i])
	}
	fmt.Println(old_slice)
}
