package main

import "fmt"

func main() {
	//day := 3
	switch day := 3; day {
	case 1:
		fmt.Println("sunday")
		fallthrough
	case 2:
		fmt.Println("monday")
		fallthrough
	case 3:
		fmt.Println("Tuesday")
		fallthrough
	case 4:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")

	}
}
