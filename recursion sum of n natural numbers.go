package main

import "fmt"

func sum(number int) int {
	if number > 0 {
		return number + sum(number-1)
	} else {
		return 0
	}
}
func main() {
	total := sum(100)
	fmt.Println("The sum of n natural numbers is: ", total)

}
