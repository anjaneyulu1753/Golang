package main

import "fmt"

func main() {
	integers := [6]int{}
	values := [5]int{0: 3, 3: 9}
	integers[0] = 23
	integers[2] = 74
	integers[1] = 38
	integers[5] = 93
	fmt.Println(integers)
	fmt.Println(values)
}
