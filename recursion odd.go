package main

import "fmt"

func odd(number int) {
	if number <= 100 {
		fmt.Println(number)
		odd(number + 2)
	}
}
func main() {
	odd(2)
}
