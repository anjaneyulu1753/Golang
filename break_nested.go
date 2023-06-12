package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {

			if i == 2 {
				break
			}
			fmt.Println(i, j)
		}
		if i == 3 {
			break
		}
		//fmt.Println(i, j)
	}
}
