package main

import "fmt"

func main() {
	new_slice := []int{2, 4, 6, 7, 8, 1, 3, 5, 9}
	for _, value := range new_slice[1:] {
		fmt.Println(value)
	}
}
