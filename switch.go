package main

import "fmt"

func main() {
	num_Of_Days := 28

	switch {
	case 29 >= num_Of_Days:
		fmt.Println("The given month is februaary")
	default:
		fmt.Println(" the given month is not a fubruary")

	}
}
