package main

import "fmt"

func factorial(number int) int {
	if number > 0 {
		return number * factorial(number-1)
	} else {
		return 1
	}

}
func main() {
	total := factorial(5)
	fmt.Println("the factorial value of 5 is: ", total)
}
