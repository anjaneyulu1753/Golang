package main

import "fmt"

func main() {
	array := [2][4]int{{2, 4}, {4, 5, 8, 9}}

	fmt.Println(array)
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(array[i][j])
		}
	}
}
