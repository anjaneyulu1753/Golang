package main

import "fmt"

func main() {
	even_slice := []int{2, 4, 6, 8}
	odd_slice := []int{1, 3, 5, 7}
	fmt.Println(cap(even_slice))
	even_slice = append(even_slice, odd_slice...)
	fmt.Println(even_slice)
	copy(even_slice, odd_slice)
	fmt.Println(even_slice)

	for value := range even_slice {
		fmt.Println(even_slice[value])
	}
}
