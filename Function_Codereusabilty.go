package main

import "fmt"

func square_num(num int) {
	square := num * num
	fmt.Printf("The square of the %d is: %d", num, square)
}

func main() {
	square_num(3)
}
