package main

import "fmt"

func main() {
	new_slice := []int{4, 6, 8, 9, 5, 3, 7}
	old_slice := []int{3, 6, 4, 2, 8, 7, 6, 3, 1, 9}

	// adding new elements to new_slice

	a := append(new_slice, 1, 2)
	b := append(new_slice, old_slice[5])
	copy(old_slice, new_slice)
	fmt.Println(new_slice)
	fmt.Println(a)
	fmt.Println(b)

}
