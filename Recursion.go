package main

import "fmt"

func countdown(number int) {

	if number > 0 {
		fmt.Println(number)
		countdown(number - 1)
	} else {
		fmt.Println("countdown stops")
	}

}
func main() {
	countdown(10)
}
